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
)

func main() {

	conf := config.ParseConfig()

	repo := inmemory.NewStorage()

	generator := strgenerator.NewGenerator(
		strgenerator.Digits+strgenerator.UppercaseLatin+strgenerator.LowercaseLatin,
		10,
		42,
	)

	svc := service.NewShortener(conf.UrlShortenerAddress, repo, generator)

	h := chihttp.NewHandler(svc)

	r := chi.NewRouter()

	h.RegisterRoutes(r)

	server := &http.Server{
		Addr:    conf.ServerAddress,
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
