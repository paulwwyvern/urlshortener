package chihttp

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type ShortenerService interface {
	GenerateURL(ctx context.Context, url string) (string, error)
	GetURL(ctx context.Context, shortURL string) (string, error)
	Ping(ctx context.Context) error
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
	r.Get("/ping", h.Ping)
	r.Get("/{url}", h.GetURL)
	r.Post("/", h.GenerateURL)
	r.Post("/api/shorten", h.GenerateURLJson)
}

func (h *Handler) GenerateURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(io.LimitReader(r.Body, h.maxBodyLength))
	if err != nil {
		if errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, context.Canceled) {
			w.WriteHeader(http.StatusBadRequest)
		} else if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusRequestTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	shortURL, err := h.service.GenerateURL(ctx, string(body))

	if err != nil {
		if errors.Is(err, errs.ErrInternalError) {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(shortURL))
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	shortURL := chi.URLParam(r, "url")

	url, err := h.service.GetURL(ctx, shortURL)
	if err != nil {
		if errors.Is(err, errs.ErrShortUrlNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else if errors.Is(err, errs.ErrInternalError) {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) GenerateURLJson(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	body, err := io.ReadAll(io.LimitReader(r.Body, h.maxBodyLength))
	if err != nil {
		if errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, context.Canceled) {
			w.WriteHeader(http.StatusBadRequest)
		} else if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusRequestTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	req := model.GenerateURLJsonRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := h.service.GenerateURL(context.Background(), req.URL)
	if err != nil {
		if errors.Is(err, errs.ErrInternalError) {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	res := model.GenerateURLJsonResponse{
		Result: url,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.service.Ping(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
