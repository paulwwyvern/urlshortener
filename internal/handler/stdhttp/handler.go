package stdhttp

import (
	"fmt"
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

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /", h.GetURL)
	mux.HandleFunc("POST /", h.GenerateURL)
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
	shortURL := r.URL.Path[1:]

	url, err := h.service.GetURL(shortURL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url)

	w.WriteHeader(http.StatusTemporaryRedirect)
}
