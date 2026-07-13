package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paulwwyvern/urlshortener/internal/config"
	"github.com/paulwwyvern/urlshortener/internal/handler/chihttp"
	mwaudit "github.com/paulwwyvern/urlshortener/internal/handler/middleware/audit"
	mwauth "github.com/paulwwyvern/urlshortener/internal/handler/middleware/auth"
	mwcompress "github.com/paulwwyvern/urlshortener/internal/handler/middleware/compress"
	mwlogger "github.com/paulwwyvern/urlshortener/internal/handler/middleware/logger"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/dto"
	auditlog "github.com/paulwwyvern/urlshortener/internal/repository/audit"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/file"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/inmemory"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/postgres"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/throughcache"
	"github.com/paulwwyvern/urlshortener/internal/repository/userstorage"
	auditpub "github.com/paulwwyvern/urlshortener/internal/service/audit"
	"github.com/paulwwyvern/urlshortener/internal/service/shortener"
	"github.com/paulwwyvern/urlshortener/internal/service/shortener/workers"
	"github.com/paulwwyvern/urlshortener/internal/service/user"
	"github.com/paulwwyvern/urlshortener/pkg/lrucache"
	"github.com/paulwwyvern/urlshortener/pkg/strgenerator"
	"go.uber.org/zap"
)

const (
	shortURLLen     = 10
	shortURLGenSeed = 42

	batchSize = 10

	serverReadTimeout  = 5 * time.Second
	serverWriteTimeout = 31 * time.Second
	serverIdleTimeout  = 31 * time.Second

	handlerMaxBodyLength = 1024 * 1024

	shutdownTimeout = 5 * time.Second

	migrationSource = "./migrations"

	authSignKey = "super sign key"

	purgeWorkersCount = 2
	purgeBatchSize    = 10
	purgeInterval     = 10 * time.Second

	cacheCapacity = 10
)

type URLRepository interface {
	GetURL(ctx context.Context, shortURL string) (string, error)
	GetShortURL(ctx context.Context, url string) (string, error)
	GetUserURL(ctx context.Context, userID int32) ([]dto.GetUserURLResponse, error)
	SaveURL(ctx context.Context, userID int32, shortURL string, url string) error
	SaveURLBatch(ctx context.Context, userID int32, urls []model.URL) error
	SoftDeleteURLBatch(ctx context.Context, userID int32, shortURLs []string) error
	PurgeURLBatch(ctx context.Context, urls []string) error
	Ping(context.Context) error
	Close() error
}

func main() {

	// init context
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// init logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	// parse config
	conf, err := config.ParseConfig()
	if err != nil {
		if !errors.Is(err, config.ErrConfigFileNotFound) {
			logger.Fatal("failed to parse config", zap.Error(err))
			panic("failed to parse config")
		}
		logger.Info("no config file found")
	}
	logger.Info("Service config",
		zap.String("config_path", conf.ConfigPath),
		zap.String("server_address", conf.ServerAddress),
		zap.String("base_url", conf.BaseURL),
		zap.String("file_storage_path", conf.FileStoragePath),
		zap.String("database_dsn", conf.DatabaseDsn),
	)

	// init repo
	var repo URLRepository
	if conf.DatabaseDsn != "" {
		cache := lrucache.NewLRUCache[string, string](cacheCapacity)
		storage, err := postgres.NewStorage(logger, conf.DatabaseDsn, true, migrationSource)
		if err != nil {
			logger.Fatal("failed to init postgres storage", zap.Error(err))
		}

		repo = throughcache.NewCache(cache, storage)
	} else if conf.FileStoragePath != "" {
		repo, err = file.NewStorage(logger, conf.FileStoragePath)
		if err != nil {
			logger.Fatal("failed to init file storage", zap.Error(err))
		}
	} else {
		repo, err = inmemory.NewStorage(logger)
		if err != nil {
			logger.Fatal("failed to init inmemory storage", zap.Error(err))
		}
	}

	defer repo.Close()

	// init user repo
	userRepo := userstorage.NewStorage()

	// init generator
	generator := strgenerator.NewGenerator(
		strgenerator.Digits+strgenerator.UppercaseLatin+strgenerator.LowercaseLatin,
		shortURLLen,
		shortURLGenSeed,
	)

	logger.Info("Init random generator")

	// init service
	shortenerServiceConfig := shortener.ShortenerServiceConfig{
		BaseURL:           conf.BaseURL,
		BatchSize:         batchSize,
		URLRepository:     repo,
		URLGenerator:      generator,
		PurgeWorkersCount: purgeWorkersCount,
		PurgeWorkersConfig: workers.PurgeWorkerConfig{
			BatchSize:     purgeBatchSize,
			PurgeInterval: purgeInterval,
			URLRepository: repo,
		},
	}

	shortenerService := shortener.NewShortener(logger, shortenerServiceConfig)
	defer shortenerService.Close()

	userService := user.NewService(logger, userRepo)

	// audit

	logger.Info("Init audit service")

	auditPub := auditpub.NewPublisher()

	{
		logger.Info("Init audit log logger service")
		auditSub := auditlog.NewAuditLogLogger(logger)
		auditPub.Register(auditSub)
		defer auditSub.Close()
	}

	if conf.AuditFile != "" {
		logger.Info("Init audit log file service")
		auditSub, err := auditlog.NewAuditLogFile(conf.AuditFile)
		if err != nil {
			logger.Fatal("Init audit file service", zap.Error(err))
		} else {
			auditPub.Register(auditSub)
			defer auditSub.Close()
		}
	}

	if conf.AuditURL != "" {
		logger.Info("Init audit log url service")

		auditSub := auditlog.NewAuditLogURL(conf.AuditURL)

		auditPub.Register(auditSub)
		defer auditSub.Close()
	}

	// init handler
	h := chihttp.NewHandler(logger, shortenerService, handlerMaxBodyLength)

	r := chi.NewRouter()

	// routes

	r.Use(mwlogger.WithLogger(logger))
	r.Use(mwcompress.WithCompress())

	r.Get("/ping", h.Ping)
	r.Mount("/debug", middleware.Profiler())

	r.Group(func(r chi.Router) {
		r.Use(mwaudit.WithAudit(logger, auditPub, "follow"))
		r.Get("/{url}", h.GetURL)
	})
	r.Group(func(r chi.Router) {
		r.Use(mwauth.WithAuth(authSignKey, userService))

		r.Get("/api/user/urls", h.GetUserURLs)
		r.Delete("/api/user/urls", h.DeleteURLJsonBatch)
	})
	r.Group(func(r chi.Router) {
		r.Use(mwauth.WithAuth(authSignKey, userService))

		r.Group(func(r chi.Router) {
			r.Use(mwaudit.WithAudit(logger, auditPub, "shorten"))

			r.Post("/", h.GenerateURL)
			r.Post("/api/shorten", h.GenerateURLJson)
		})
		r.Post("/api/shorten/batch", h.GenerateURLJsonBatch)
	})

	// init server
	server := &http.Server{
		Addr:         conf.ServerAddress,
		Handler:      r,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  serverReadTimeout,
		WriteTimeout: serverWriteTimeout,
		IdleTimeout:  serverIdleTimeout,
	}

	// run server
	servErr := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			servErr <- err
		}
	}()

	select {
	case err := <-servErr:
		logger.Fatal("failed to start server", zap.Error(err))
		panic("failed to start server")
	case <-ctx.Done():
		stop()
		logger.Info("shutdown signal received")
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("failed to graceful shutdown server", zap.Error(err))
	}
	logger.Info("shutdown complete")

}
