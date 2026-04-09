package chihttp

import (
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type ShortenerService interface {
	GenerateURL(url string) (string, error)
	GetURL(shortURL string) (string, error)
}

type Handler struct {
	maxBodyLength int64

	service ShortenerService
}

func NewHandler(service ShortenerService, maxBodyLength int64) *Handler {
	return &Handler{
		maxBodyLength: maxBodyLength,
		service:       service,
	}
}

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.Get("/{url}", h.GetURL)
	r.Post("/", h.GenerateURL)
}

func (h *Handler) GenerateURL(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(io.LimitReader(r.Body, h.maxBodyLength))

	shortURL, err := h.service.GenerateURL(string(body))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(shortURL))
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "url")

	url, err := h.service.GetURL(shortURL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
