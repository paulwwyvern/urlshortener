package logger

import (
	"github.com/paulwwyvern/urlshortener/internal/handler/httperr"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type responseData struct {
	status int
	size   int
}

type loggerResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func newLoggerResponseWriter(w http.ResponseWriter, response *responseData) *loggerResponseWriter {
	return &loggerResponseWriter{
		ResponseWriter: w,
		responseData:   response,
	}
}

func (w *loggerResponseWriter) Write(data []byte) (int, error) {
	size, err := w.ResponseWriter.Write(data)
	w.responseData.size += size
	return size, err
}

func (w *loggerResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.responseData.status = statusCode
}

func WithLogger(logger *zap.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			uri := r.RequestURI
			method := r.Method

			response := &responseData{}
			ww := newLoggerResponseWriter(w, response)
			h.ServeHTTP(ww, r)

			duration := time.Since(start)

			err := httperr.GetError(r)

			logger.Info("Get request",
				zap.String("uri", uri),
				zap.String("method", method),
				zap.Int("response status", response.status),
				zap.Int("response size", response.size),
				zap.Duration("duration", duration),
				zap.String("request content type", r.Header.Get("Content-Type")),
				zap.String("response content type", ww.Header().Get("Content-Type")),
				zap.String("request content encoding", r.Header.Get("Content-Encoding")),
				zap.String("accept encoding", r.Header.Get("Accept-Encoding")),
				zap.String("response content encoding", ww.Header().Get("Content-Encoding")),
				zap.Error(err),
			)

		})
	}
}
