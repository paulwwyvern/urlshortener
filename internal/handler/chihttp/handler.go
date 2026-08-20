package chihttp

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/model/dto"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httperr"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpurl"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser"
	"go.uber.org/zap"
)

type ShortenerService interface {
	GetURL(ctx context.Context, shortURL string) (string, error)
	GetUserURLs(ctx context.Context, userID int32) ([]dto.GetUserURLResponse, error)
	GenerateURL(ctx context.Context, userID int32, url string) (string, error)
	GenerateURLBatch(ctx context.Context, userID int32, urls []dto.GenerateURLBatchRequest) ([]dto.GenerateURLBatchResponse, error)
	DeleteURLBatch(ctx context.Context, userID int32, shortURLs []string) error
	GetStats(ctx context.Context) (dto.GetStatsResponse, error)
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

// GenerateURL получает в теле запроса урл в формате text/plain,
// для которого необходимо сгенерить новый короткий урл и возвращает его в формате text/plain
//
// Возвращаемые коды:
//
// 201 - если новый короткий урл успешно создан
//
// 409 - если для оригинального урла уже существует короткий урл
//
// 400 - невалидный запрос
//
// 500 - внутренняя ошибка сервера
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

	url := string(body)

	shortURL, err := h.service.GenerateURL(ctx, userID, url)
	w.Header().Set("Content-Type", "text/plain")
	if err != nil {
		if errors.Is(err, errs.ErrOriginalURLAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	httpurl.SetURL(r, url)
	w.Write([]byte(shortURL))

	return nil
}

// GetURL получает в uri короткий урл и редиректит на урл, который стоит за этим коротким урлом
//
// Возвращаемые коды:
//
// 307 - оригинальный урл по заданному короткому урлу найден и редирект выполнен
//
// 404 - если для данного короткого урла не найден оригинальный
//
// 409 - если для данного короткого урла найден оригинальный, но он в процессе удаления(считай удалён)
//
// 500 - внутренняя ошибка сервера
func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	httperr.Adapt(h.getURL)(w, r)
}

func (h *Handler) getURL(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	shortURL := chi.URLParam(r, "url")

	url, err := h.service.GetURL(ctx, shortURL)
	if err != nil {
		if errors.Is(err, errs.ErrShortURLNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else if errors.Is(err, errs.ErrShortURLGone) {
			w.WriteHeader(http.StatusGone)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
		return nil
	}

	httpurl.SetURL(r, url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
}

// GetUserURLs для данного пользователя возвращает все сгенеренные им урлы в формате json
//
// Возвращаемые коды:
//
// 200 - OK
//
// 201 - если пользователь ещё ни один урл не сгенерил
//
// 500 - внутренняя ошибка сервера
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

// GenerateURLJson получает в теле запроса урл в формате json,
// для которого необходимо сгенерить новый короткий урл и возвращает его в формате json
//
// Возвращаемые коды:
//
// 201 - если новый короткий урл успешно создан
//
// 409 - если для оригинального урла уже существует короткий урл
//
// 400 - невалидный запрос
//
// 500 - внутренняя ошибка сервера
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

	req := dto.GenerateURLJsonRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	userID := httpuser.GetUserID(r)

	url, err := h.service.GenerateURL(ctx, userID, req.URL)

	if err != nil {
		if errors.Is(err, errs.ErrOriginalURLAlreadyExists) {
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

	res := dto.GenerateURLJsonResponse{
		Result: url,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	httpurl.SetURL(r, url)

	return nil
}

// GenerateURLJsonBatch получает в теле запроса урлы в формате json,
// для которых необходимо сгенерить новые короткий урлы и возвращают их в формате json
//
// Возвращаемые коды:
//
// 201 - если новые короткие урлы успешно созданы(даже если для некоторых уже существовали короткие)
//
// 400 - невалидный запрос
//
// 500 - внутренняя ошибка сервера
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

	req := []dto.GenerateURLBatchRequest{}
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

// DeleteURLJsonBatch получает в теле запроса короткие урлы в формате json,
// которые необходимо удалить
//
// Возвращаемые коды:
//
// 202 - Запрос на удаление принят(сами урлы при попытке запроса будут пока выдавать 409, пока они окончательно не удалятся)
//
// 400 - невалидный запрос
//
// 500 - внутренняя ошибка сервера
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

// GetStats возвращает статистику(количество зареганных урлов и юзеров)
//
// Возвращаемые коды:
//
// 200 - Всё ок
//
// 500 - внутренняя ошибка сервера
func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	httperr.Adapt(h.getStats)(w, r)
}

func (h *Handler) getStats(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	stats, err := h.service.GetStats(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		return err
	}

	return nil
}

// Ping пингует сервер
//
// Возвращаемые коды:
//
// 200 - Всё ок
//
// 500 - внутренняя ошибка сервера
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
