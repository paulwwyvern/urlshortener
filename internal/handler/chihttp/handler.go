package chihttp

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"go.uber.org/zap"
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

func NewHandler(logger *zap.Logger, service ShortenerService, maxBodyLength int64) *Handler {
	logger.Info("Initializing chi handlers")
	return &Handler{
		maxBodyLength: maxBodyLength,
		service:       service,
	}
}

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.Get("/{url}", h.GetURL)
	r.Post("/", h.GenerateURL)
	r.Post("/api/shorten", h.GenerateUrlJson)
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

func (h *Handler) GenerateUrlJson(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(io.LimitReader(r.Body, h.maxBodyLength))

	req := model.GenerateUrlJsonRequest{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := h.service.GenerateURL(req.Url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := model.GenerateUrlJsonResponse{
		Result: url,
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(res)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
