package chihttp

import (
	"github.com/go-chi/chi/v5"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
				code:     400,
				response: ``,
			},
			wantErr: errs.ErrInternalError,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			svc := NewMockShortenerService(ctrl)

			svc.EXPECT().GenerateURL(gomock.Any(), tt.body).Return(tt.want.response, tt.wantErr)

			h := NewHandler(logger, svc, 1024)
			mux := chi.NewRouter()
			h.RegisterRoutes(mux)

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
				code:     400,
				location: ``,
			},
			wantErr: errs.ErrShortUrlNotFound,
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
			h.RegisterRoutes(mux)

			r := httptest.NewRequest(http.MethodGet, tt.url, nil)

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)
			assert.Equal(t, tt.want.location, w.Header().Get("Location"))

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
		name    string
		url     string
		body    string
		request string
		want    want
		wantErr error
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
			wantErr: nil,
		},
		{
			name:    "Test #2 Internal error",
			url:     "/api/shorten",
			body:    `{"url":"http://example.com"}`,
			request: "http://example.com",
			want: want{
				code: 400,
			},
			wantErr: errs.ErrInternalError,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			svc := NewMockShortenerService(ctrl)

			svc.EXPECT().GenerateURL(gomock.Any(), tt.request).Return(tt.want.response, tt.wantErr)

			h := NewHandler(logger, svc, 1024)
			mux := chi.NewRouter()
			h.RegisterRoutes(mux)

			r := httptest.NewRequest(http.MethodPost, tt.url, strings.NewReader(tt.body))

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)
			assert.Equal(t, tt.want.body, w.Body.String())
			assert.Equal(t, tt.want.contentType, w.Header().Get("Content-Type"))

		})
	}
}

func TestHandler_Router(t *testing.T) {
	type want struct {
		code     int
		body     string
		response string
		location string
	}

	tests := []struct {
		name    string
		method  string
		url     string
		body    string
		request string
		want    want
	}{
		{
			name:    "Test #1 Post /",
			method:  http.MethodPost,
			url:     "/",
			body:    "http://example.com",
			request: "http://example.com",

			want: want{
				code:     201,
				response: `http://localhost:8080/`,
				body:     `http://localhost:8080/`,
			},
		}, {
			name:   "Test #2 Get /short",
			method: http.MethodGet,
			url:    "/short",
			want: want{
				code:     307,
				body:     "<a href=\"http://example.com\">Temporary Redirect</a>.\n\n",
				location: "http://example.com",
			},
		}, {
			name:    "Test #3 Get /api/shorten",
			method:  http.MethodPost,
			url:     "/api/shorten",
			body:    `{"url":"http://example.com"}`,
			request: "http://example.com",
			want: want{
				code:     201,
				body:     "{\"result\":\"http://localhost:8080/\"}\n",
				response: `http://localhost:8080/`,
			},
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			svc := NewMockShortenerService(ctrl)

			if tt.method == http.MethodGet {
				svc.EXPECT().GetURL(gomock.Any(), tt.url[1:]).Return(tt.want.location, nil)
			}
			if tt.method == http.MethodPost {
				svc.EXPECT().GenerateURL(gomock.Any(), tt.request).Return(tt.want.response, nil)
			}

			h := NewHandler(logger, svc, 1024)
			mux := chi.NewRouter()
			h.RegisterRoutes(mux)

			r := httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)
			assert.Equal(t, tt.want.body, w.Body.String())
			assert.Equal(t, tt.want.location, w.Header().Get("Location"))
		})
	}
}
