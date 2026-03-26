package main

import (
	"github.com/paulwwyvern/urlshortener/internal/handler/stdhttp"
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

	h := stdhttp.NewHandler(svc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
