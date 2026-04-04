package chihttp

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type ShortenerService interface {
	GenerateURL(url string) (string, error)
	GetURL(shortURL string) (string, error)
}

type Handler struct {
	service ShortenerService
}

func NewHandler(service ShortenerService) *Handler {
	return &Handler{

		service: service,
	}
}

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.Get("/{url}", h.GetURL)
	r.Post("/", h.GenerateURL)
}

func (h *Handler) GenerateURL(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	shortURL, err := h.service.GenerateURL(string(body))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "%s", shortURL)
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "url")

	url, err := h.service.GetURL(shortURL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url)

	w.WriteHeader(http.StatusTemporaryRedirect)
}
