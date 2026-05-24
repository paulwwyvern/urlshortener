package chihttp

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httperr"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type ShortenerService interface {
	GetURL(ctx context.Context, shortURL string) (string, error)
	GetUserURLs(ctx context.Context, userId int32) ([]model.GetUserURLResponse, error)
	GenerateURL(ctx context.Context, userId int32, url string) (string, error)
	GenerateURLBatch(ctx context.Context, userId int32, urls []model.GenerateURLBatchRequest) ([]model.GenerateURLBatchResponse, error)
	DeleteURLBatch(ctx context.Context, userId int32, shortURLs []string) error
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

func (h *Handler) GenerateURL(w http.ResponseWriter, r *http.Request) {
	httperr.Adapt(h.generateURL)(w, r)
}

func (h *Handler) generateURL(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	body, err := io.ReadAll(io.LimitReader(r.Body, h.maxBodyLength))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusRequestTimeout)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return err
	}

	userID := httpuser.GetUserID(r)

	shortURL, err := h.service.GenerateURL(ctx, userID, string(body))
	w.Header().Set("Content-Type", "text/plain")
	if err != nil {
		if errors.Is(err, errs.ErrOriginalUrlAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	w.Write([]byte(shortURL))

	return nil
}

func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	httperr.Adapt(h.getURL)(w, r)
}

func (h *Handler) getURL(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	shortURL := chi.URLParam(r, "url")

	url, err := h.service.GetURL(ctx, shortURL)
	if err != nil {
		if errors.Is(err, errs.ErrShortUrlNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else if errors.Is(err, errs.ErrShortUrlGone) {
			w.WriteHeader(http.StatusGone)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
		return nil
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
}

func (h *Handler) GetUserURLs(w http.ResponseWriter, r *http.Request) {
	httperr.Adapt(h.getUserURLs)(w, r)
}

func (h *Handler) getUserURLs(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	userID := httpuser.GetUserID(r)

	userURLs, err := h.service.GetUserURLs(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	if len(userURLs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(userURLs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return nil
}

func (h *Handler) GenerateURLJson(w http.ResponseWriter, r *http.Request) {
	httperr.Adapt(h.generateURLJson)(w, r)
}

func (h *Handler) generateURLJson(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	body, err := io.ReadAll(io.LimitReader(r.Body, h.maxBodyLength))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusRequestTimeout)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return err
	}

	req := model.GenerateURLJsonRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	userID := httpuser.GetUserID(r)

	url, err := h.service.GenerateURL(ctx, userID, req.URL)

	if err != nil {
		if errors.Is(err, errs.ErrOriginalUrlAlreadyExists) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

	res := model.GenerateURLJsonResponse{
		Result: url,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return nil
}

func (h *Handler) GenerateURLJsonBatch(w http.ResponseWriter, r *http.Request) {
	httperr.Adapt(h.generateURLJsonBatch)(w, r)
}

func (h *Handler) generateURLJsonBatch(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	body, err := io.ReadAll(io.LimitReader(r.Body, h.maxBodyLength))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusRequestTimeout)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return err
	}

	req := []model.GenerateURLBatchRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	userID := httpuser.GetUserID(r)

	res, err := h.service.GenerateURLBatch(ctx, userID, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return nil
}

func (h *Handler) DeleteURLJsonBatch(w http.ResponseWriter, r *http.Request) {
	httperr.Adapt(h.deleteURLJsonBatch)(w, r)
}

func (h *Handler) deleteURLJsonBatch(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	body, err := io.ReadAll(io.LimitReader(r.Body, h.maxBodyLength))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusRequestTimeout)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return err
	}

	req := []string{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	userID := httpuser.GetUserID(r)
	err = h.service.DeleteURLBatch(ctx, userID, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusAccepted)
	return nil
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	httperr.Adapt(h.ping)(w, r)
}

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	err := h.service.Ping(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
