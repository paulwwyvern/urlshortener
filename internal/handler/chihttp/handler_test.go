package chihttp

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//go:generate mockgen -source=handler.go -destination=mock_handler.go -package=chihttp

func TestHandler_GenerateURL(t *testing.T) {
	type want struct {
		code     int
		response string
	}

	tests := []struct {
		name    string
		url     string
		body    string
		want    want
		wantErr error
	}{
		{
			name: "Test #1 Success",
			url:  "/",
			body: "http://example.com",
			want: want{
				code:     201,
				response: `http://localhost:8080`,
			},
			wantErr: nil,
		},
		{
			name: "Test #2 Internal error",
			url:  "/",
			body: "http://example.com",
			want: want{
				code:     500,
				response: ``,
			},
			wantErr: errors.New("Internal error"),
		},
		{
			name: "Test #3 Conflict",
			url:  "/",
			body: "http://example.com",
			want: want{
				code:     409,
				response: `http://localhost:8080`,
			},
			wantErr: errs.ErrOriginalUrlAlreadyExists,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			svc := NewMockShortenerService(ctrl)

			svc.EXPECT().GenerateURL(gomock.Any(), gomock.Any(), tt.body).Return(tt.want.response, tt.wantErr)

			h := NewHandler(logger, svc, 1024)
			mux := chi.NewRouter()
			mux.Post("/", h.GenerateURL)

			r := httptest.NewRequest(http.MethodPost, tt.url, strings.NewReader(tt.body))

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)
			assert.Equal(t, tt.want.response, w.Body.String())

		})
	}
}

func TestHandler_GetURL(t *testing.T) {
	type want struct {
		code     int
		location string
	}

	tests := []struct {
		name    string
		method  string
		url     string
		want    want
		wantErr error
	}{
		{
			name: "Test #1 Success",
			url:  "/Gs7K09wks",
			want: want{
				code:     307,
				location: "http://example.com",
			},
			wantErr: nil,
		}, {
			name: "Test #2 Not found",
			url:  "/Scuf38812",
			want: want{
				code:     404,
				location: ``,
			},
			wantErr: errs.ErrShortUrlNotFound,
		}, {
			name: "Test #3 Gone",
			url:  "/Scuf38812",
			want: want{
				code:     410,
				location: ``,
			},
			wantErr: errs.ErrShortUrlGone,
		}, {
			name: "Test #4 Internal error",
			url:  "/Scuf38812",
			want: want{
				code:     500,
				location: ``,
			},
			wantErr: errors.New("Internal error"),
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			svc := NewMockShortenerService(ctrl)

			svc.EXPECT().GetURL(gomock.Any(), tt.url[1:]).Return(tt.want.location, tt.wantErr)

			h := NewHandler(logger, svc, 1024)
			mux := chi.NewRouter()
			mux.Get("/{url}", h.GetURL)

			r := httptest.NewRequest(http.MethodGet, tt.url, nil)

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)
			assert.Equal(t, tt.want.location, w.Header().Get("Location"))

		})
	}
}

func TestHandler_GetUserURLs(t *testing.T) {
	type want struct {
		code int
		body string
	}

	tests := []struct {
		name    string
		method  string
		url     string
		want    want
		wantErr error
	}{
		{
			name: "Test #1 Success",
			url:  "/api/user/urls",
			want: want{
				code: 200,
				body: `[{"short_url": "http://localhost:8080/gIb8VTucox","original_url": "https://example.com"},
						{"short_url": "http://localhost:8080/et6TAyR6xm","original_url": "https://yandex.com"}]`,
			},
			wantErr: nil,
		}, {
			name: "Test #2 No Content",
			url:  "/api/user/urls",
			want: want{
				code: 204,
				body: "",
			},
			wantErr: nil,
		}, {
			name: "Test #3 Internal error",
			url:  "/api/user/urls",
			want: want{
				code: 500,
				body: "",
			},
			wantErr: errors.New("Internal error"),
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			svc := NewMockShortenerService(ctrl)

			svcResponse := []model.GetUserURLResponse{}
			json.Unmarshal([]byte(tt.want.body), &svcResponse)

			svc.EXPECT().GetUserURLs(gomock.Any(), gomock.Any()).Return(svcResponse, tt.wantErr)

			h := NewHandler(logger, svc, 1024)
			mux := chi.NewRouter()
			mux.Get("/api/user/urls", h.GetUserURLs)

			r := httptest.NewRequest(http.MethodGet, tt.url, nil)

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)
			if tt.want.body != "" {
				assert.JSONEq(t, tt.want.body, w.Body.String())
			} else {
				assert.Empty(t, w.Body.String())
			}
		})
	}
}

func TestHandler_GenerateUrlJson(t *testing.T) {
	type want struct {
		code        int
		body        string
		response    string
		contentType string
	}

	tests := []struct {
		name            string
		url             string
		body            string
		request         string
		want            want
		wantErr         error
		wantServiceCall bool
	}{
		{
			name:    "Test #1 Success",
			url:     "/api/shorten",
			body:    `{"url":"http://example.com"}`,
			request: "http://example.com",
			want: want{
				code:        201,
				body:        "{\"result\":\"http://localhost:8080/\"}\n",
				response:    `http://localhost:8080/`,
				contentType: `application/json`,
			},
			wantErr:         nil,
			wantServiceCall: true,
		},
		{
			name:    "Test #2 Internal error",
			url:     "/api/shorten",
			body:    `{"url":"http://example.com"}`,
			request: "http://example.com",
			want: want{
				code: 500,
			},
			wantErr:         errors.New("Internal error"),
			wantServiceCall: true,
		},
		{
			name:    "Test #3 Bad request",
			url:     "/api/shorten",
			body:    `{"url":"http://exa`,
			request: "http://example.com",
			want: want{
				code: 400,
			},
			wantErr:         nil,
			wantServiceCall: false,
		},
		{
			name:    "Test #4 Conflict",
			url:     "/api/shorten",
			body:    `{"url":"http://example.com"}`,
			request: "http://example.com",
			want: want{
				code:        409,
				body:        "{\"result\":\"http://localhost:8080/\"}\n",
				response:    `http://localhost:8080/`,
				contentType: `application/json`,
			},
			wantErr:         errs.ErrOriginalUrlAlreadyExists,
			wantServiceCall: true,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			svc := NewMockShortenerService(ctrl)
			if tt.wantServiceCall {
				svc.EXPECT().GenerateURL(gomock.Any(), gomock.Any(), tt.request).Return(tt.want.response, tt.wantErr)
			}
			h := NewHandler(logger, svc, 1024)
			mux := chi.NewRouter()
			mux.Post("/api/shorten", h.GenerateURLJson)

			r := httptest.NewRequest(http.MethodPost, tt.url, strings.NewReader(tt.body))

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)
			assert.Equal(t, tt.want.body, w.Body.String())
			assert.Equal(t, tt.want.contentType, w.Header().Get("Content-Type"))

		})
	}
}

func TestHandler_GenerateUrlBatch(t *testing.T) {
	type want struct {
		code        int
		body        string
		contentType string
	}

	tests := []struct {
		name            string
		url             string
		body            string
		want            want
		wantErr         error
		wantServiceCall bool
	}{
		{
			name: "Test #1 Success",
			url:  "/api/shorten/batch",
			body: `[{"correlation_id": "a", "original_url": "https://yandex.ru"},
 			        {"correlation_id": "b", "original_url": "https://google.com"},
				    {"correlation_id": "c", "original_url": "https://example.com"}]`,
			want: want{
				code: 201,
				body: `[{"correlation_id": "a", "short_url": "http://localhost:8080/gIb8VTucox"},
                        {"correlation_id": "b", "short_url": "http://localhost:8080/et6TAyR6xm"},
                        {"correlation_id": "c", "short_url": "http://localhost:8080/pvPgKzcNdC"}]`,
				contentType: `application/json`,
			},
			wantErr:         nil,
			wantServiceCall: true,
		},
		{
			name: "Test #2 Internal Error",
			url:  "/api/shorten/batch",
			body: `[{"correlation_id": "a", "original_url": "https://yandex.ru"},
 			        {"correlation_id": "b", "original_url": "https://google.com"},
				    {"correlation_id": "c", "original_url": "https://example.com"}]`,
			want: want{
				code: 500,
			},
			wantErr:         errors.New("Internal error"),
			wantServiceCall: true,
		},
		{
			name: "Test #3 Bad request",
			url:  "/api/shorten/batch",
			body: `[{"correlation_id": "a", "original_url": "https://yandex.ru"},
 			        {"correlation_id": "b", "original_ur`,
			want: want{
				code: 400,
			},
			wantErr:         errors.New("Internal error"),
			wantServiceCall: false,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			svc := NewMockShortenerService(ctrl)

			if tt.wantServiceCall {
				svcRequest := []model.GenerateURLBatchRequest{}
				json.Unmarshal([]byte(tt.body), &svcRequest)

				svcResponse := []model.GenerateURLBatchResponse{}
				json.Unmarshal([]byte(tt.want.body), &svcResponse)

				svc.EXPECT().GenerateURLBatch(gomock.Any(), gomock.Any(), svcRequest).Return(svcResponse, tt.wantErr)
			}
			h := NewHandler(logger, svc, 1024)
			mux := chi.NewRouter()
			mux.Post("/api/shorten/batch", h.GenerateURLJsonBatch)

			r := httptest.NewRequest(http.MethodPost, tt.url, strings.NewReader(tt.body))

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)
			if tt.want.body != "" {
				assert.JSONEq(t, tt.want.body, w.Body.String())
			} else {
				assert.Empty(t, w.Body.String())
			}
			assert.Equal(t, tt.want.contentType, w.Header().Get("Content-Type"))

		})
	}
}

func TestHandler_DeleteURLJsonBatch(t *testing.T) {
	type want struct {
		code int
	}

	tests := []struct {
		name    string
		method  string
		url     string
		body    string
		want    want
		wantErr error
	}{
		{
			name: "Test #1 Success",
			url:  "/api/user/urls",
			body: `["aaa","bbb","ccc","ddd"]`,
			want: want{
				code: 202,
			},
			wantErr: nil,
		}, {
			name: "Test #3 Internal error",
			url:  "/api/user/urls",
			body: `["aaa","bbb","ccc","ddd"]`,

			want: want{
				code: 500,
			},
			wantErr: errors.New("Internal error"),
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			svc := NewMockShortenerService(ctrl)

			svcRequest := []string{}
			json.Unmarshal([]byte(tt.body), &svcRequest)

			svc.EXPECT().DeleteURLBatch(gomock.Any(), gomock.Any(), svcRequest).Return(tt.wantErr)

			h := NewHandler(logger, svc, 1024)
			mux := chi.NewRouter()
			mux.Delete("/api/user/urls", h.DeleteURLJsonBatch)

			r := httptest.NewRequest(http.MethodDelete, tt.url, strings.NewReader(tt.body))

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)

			assert.Empty(t, w.Body.String())

		})
	}
}
