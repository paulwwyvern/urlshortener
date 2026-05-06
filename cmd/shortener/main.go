package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/config"
	"github.com/paulwwyvern/urlshortener/internal/handler/chihttp"
	mwcompress "github.com/paulwwyvern/urlshortener/internal/handler/middleware/compress"
	mwlogger "github.com/paulwwyvern/urlshortener/internal/handler/middleware/logger"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/file"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/inmemory"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/postgres"
	"github.com/paulwwyvern/urlshortener/internal/service"
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

	serverReadTimeout  = 5 * time.Second
	serverWriteTimeout = 5 * time.Second
	serverIdleTimeout  = 30 * time.Second

	handlerMaxBodyLength = 1024 * 1024

	shutdownTimeout = 5 * time.Second
)

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
	confPath := config.ParseConfigPath()

	conf, err := config.ParseConfig(confPath)
	if err != nil {
		if !errors.Is(err, config.ErrConfigFileNotFound) {
			logger.Fatal("failed to parse config", zap.Error(err))
			os.Exit(1)
		}
		logger.Info("no config file found")
	}
	logger.Info("Service config",
		zap.String("config_path", confPath),
		zap.String("server_address", conf.ServerAddress),
		zap.String("base_url", conf.BaseUrl),
		zap.String("file_storage_path", conf.FileStoragePath),
		zap.String("database_dsn", conf.DatabaseDsn),
	)

	// init repo
	var repo service.UrlRepository
	if conf.DatabaseDsn != "" {
		repo, err = postgres.NewStorage(logger, conf.DatabaseDsn)
	} else if conf.FileStoragePath != "" {
		repo, err = file.NewStorage(logger, conf.FileStoragePath)
	} else {
		repo, err = inmemory.NewStorage(logger)
	}
	if err != nil {
		logger.Fatal("failed to init storage", zap.Error(err))
	}
	defer repo.Close()

	// init generator
	generator := strgenerator.NewGenerator(
		strgenerator.Digits+strgenerator.UppercaseLatin+strgenerator.LowercaseLatin,
		shortUrlLen,
		shortUrlGenSeed,
	)

	// init service
	svc := service.NewShortener(logger, conf.BaseUrl, repo, generator)

	// init handler
	h := chihttp.NewHandler(logger, svc, handlerMaxBodyLength)

	r := chi.NewRouter()
	r.Use(mwlogger.WithLogger(logger))
	r.Use(mwcompress.WithCompress())

	h.RegisterRoutes(r)

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
