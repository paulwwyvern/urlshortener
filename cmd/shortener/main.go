package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/config"
	"github.com/paulwwyvern/urlshortener/internal/handler/chihttp"
	mwauth "github.com/paulwwyvern/urlshortener/internal/handler/middleware/auth"
	mwcompress "github.com/paulwwyvern/urlshortener/internal/handler/middleware/compress"
	mwlogger "github.com/paulwwyvern/urlshortener/internal/handler/middleware/logger"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/file"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/inmemory"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/postgres"
	"github.com/paulwwyvern/urlshortener/internal/repository/userstorage"
	"github.com/paulwwyvern/urlshortener/internal/service/shortener"
	"github.com/paulwwyvern/urlshortener/internal/service/shortener/workers"
	"github.com/paulwwyvern/urlshortener/internal/service/user"
	"github.com/paulwwyvern/urlshortener/pkg/strgenerator"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	shortUrlLen     = 10
	shortUrlGenSeed = 42

	batchSize = 10

	serverReadTimeout  = 5 * time.Second
	serverWriteTimeout = 5 * time.Second
	serverIdleTimeout  = 30 * time.Second

	handlerMaxBodyLength = 1024 * 1024

	shutdownTimeout = 5 * time.Second

	migrationSource = "./migrations"

	authSignKey = "super sign key"

	purgeWorkersCount = 2
	purgeBatchSize    = 10
	purgeInterval     = 10 * time.Second
)

type UrlRepository interface {
	GetURL(ctx context.Context, shortUrl string) (string, error)
	GetShortURL(ctx context.Context, url string) (string, error)
	GetUserURL(ctx context.Context, userId int32) ([]model.GetUserURLResponse, error)
	SaveURL(ctx context.Context, userId int32, shortUrl string, url string) error
	SaveURLBatch(ctx context.Context, userId int32, urls []model.URL) error
	SoftDeleteURLBatch(ctx context.Context, userId int32, shortUrls []string) error
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
			os.Exit(1)
		}
		logger.Info("no config file found")
	}
	logger.Info("Service config",
		zap.String("config_path", conf.ConfigPath),
		zap.String("server_address", conf.ServerAddress),
		zap.String("base_url", conf.BaseUrl),
		zap.String("file_storage_path", conf.FileStoragePath),
		zap.String("database_dsn", conf.DatabaseDsn),
	)

	// init repo
	var repo UrlRepository
	if conf.DatabaseDsn != "" {
		repo, err = postgres.NewStorage(logger, conf.DatabaseDsn, true, migrationSource)
	} else if conf.FileStoragePath != "" {
		repo, err = file.NewStorage(logger, conf.FileStoragePath)
	} else {
		repo, err = inmemory.NewStorage(logger)
	}
	if err != nil {
		logger.Fatal("failed to init storage", zap.Error(err))
	}
	defer repo.Close()

	// init user repo
	userRepo := userstorage.NewStorage()

	// init generator
	generator := strgenerator.NewGenerator(
		strgenerator.Digits+strgenerator.UppercaseLatin+strgenerator.LowercaseLatin,
		shortUrlLen,
		shortUrlGenSeed,
	)

	logger.Info("Init random generator")

	// init service
	shortenerServiceConfig := shortener.ShortenerServiceConfig{
		BaseUrl:           conf.BaseUrl,
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

	// init handler
	h := chihttp.NewHandler(logger, shortenerService, handlerMaxBodyLength)

	r := chi.NewRouter()

	// routes

	r.Use(mwlogger.WithLogger(logger))
	r.Use(mwcompress.WithCompress())

	r.Get("/{url}", h.GetURL)
	r.Get("/ping", h.Ping)
	r.Group(func(r chi.Router) {
		r.Use(mwauth.WithAuth(authSignKey, userService))

		r.Get("/api/user/urls", h.GetUserURLs)
		r.Delete("/api/user/urls", h.DeleteURLJsonBatch)
	})
	r.Group(func(r chi.Router) {
		r.Use(mwauth.WithAuth(authSignKey, userService))

		r.Post("/", h.GenerateURL)
		r.Post("/api/shorten", h.GenerateURLJson)
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
		os.Exit(1)
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
