package audit

import (
	"net/http"
	"time"

	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpurl"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser"
	"go.uber.org/zap"
)

type AuditPublisher interface {
	Update(event *model.AuditEvent) error
}

func WithAudit(logger *zap.Logger, pub AuditPublisher, action string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(httpurl.AddURL(r.Context()))

			h.ServeHTTP(w, r)

			timestamp := time.Now().Unix()
			userID := httpuser.GetUserID(r.Context())
			url, ok := httpurl.GetURL(r.Context())

			if !ok {
				return
			}

			event := &model.AuditEvent{
				TS:     timestamp,
				URL:    url,
				Action: action,
				UserID: userID,
			}

			err := pub.Update(event)
			if err != nil {
				logger.Info("Error updating audit event", zap.Error(err))
			}
		})
	}
}
