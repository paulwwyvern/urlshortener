package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/handler/chihttp"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/inmemory"
	"github.com/paulwwyvern/urlshortener/internal/service"
	"github.com/paulwwyvern/urlshortener/pkg/strgenerator"
	"log"
	"net/http"
)

func main() {

	repo := inmemory.NewStorage()

	generator := strgenerator.NewGenerator(
		strgenerator.Digits+strgenerator.UppercaseLatin+strgenerator.LowercaseLatin,
		10,
		42,
	)

	svc := service.NewShortener("http://localhost:8080", repo, generator)

	h := chihttp.NewHandler(svc)

	r := chi.NewRouter()

	h.RegisterRoutes(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
