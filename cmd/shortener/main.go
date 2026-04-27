package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/config"
	"github.com/paulwwyvern/urlshortener/internal/handler/chihttp"
	mwcompress "github.com/paulwwyvern/urlshortener/internal/handler/middleware/compress"
	mwlogger "github.com/paulwwyvern/urlshortener/internal/handler/middleware/logger"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/file"
	"github.com/paulwwyvern/urlshortener/internal/service"
	"github.com/paulwwyvern/urlshortener/pkg/strgenerator"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	shortUrlLen     = 10
	shortUrlGenSeed = 42

	serverReadTimeout  = 5 * time.Second
	serverWriteTimeout = 5 * time.Second
	serverIdleTimeout  = 30 * time.Second

	handlerMaxBodyLength = 1024 * 1024
)

func main() {

	// init logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// parse config
	conf, err := config.ParseConfig()
	if err != nil {
		logger.Fatal("failed to parse config", zap.Error(err))
	}

	logger.Info("Service config",
		zap.String("server_address", conf.ServerAddress),
		zap.String("base_url: %s", conf.BaseUrl),
	)

	// init repo
	repo, err := file.NewStorage(conf.FileStoragePath, logger)
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
		ReadTimeout:  serverReadTimeout,
		WriteTimeout: serverWriteTimeout,
		IdleTimeout:  serverIdleTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
