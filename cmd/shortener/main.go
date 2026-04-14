package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/config"
	"github.com/paulwwyvern/urlshortener/internal/handler/chihttp"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/inmemory"
	"github.com/paulwwyvern/urlshortener/internal/service"
	"github.com/paulwwyvern/urlshortener/pkg/strgenerator"
	"log"
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

	conf, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server address: %s", conf.ServerAddress)
	log.Printf("Base url: %s", conf.BaseUrl)

	repo := inmemory.NewStorage()

	generator := strgenerator.NewGenerator(
		strgenerator.Digits+strgenerator.UppercaseLatin+strgenerator.LowercaseLatin,
		shortUrlLen,
		shortUrlGenSeed,
	)

	svc := service.NewShortener(conf.BaseUrl, repo, generator)

	h := chihttp.NewHandler(svc, handlerMaxBodyLength)

	r := chi.NewRouter()

	h.RegisterRoutes(r)

	server := &http.Server{
		Addr:         conf.ServerAddress,
		Handler:      r,
		ReadTimeout:  serverReadTimeout,
		WriteTimeout: serverWriteTimeout,
		IdleTimeout:  serverIdleTimeout,
	}

	log.Fatal(server.ListenAndServe())
}
