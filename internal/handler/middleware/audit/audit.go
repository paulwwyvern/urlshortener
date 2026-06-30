package audit

import (
	"net/http"
	"time"

	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httperr"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpurl"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser"
)

type AuditPublisher interface {
	Update(event *model.AuditEvent) error
}

func WithAudit(pub AuditPublisher, action string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return httperr.Adapt(func(w http.ResponseWriter, r *http.Request) error {
			h.ServeHTTP(w, r)

			timestamp := time.Now().Unix()
			userId := httpuser.GetUserID(r)
			url := httpurl.GetURL(r)

			event := &model.AuditEvent{
				TS:     timestamp,
				URL:    url,
				Action: action,
				UserID: userId,
			}

			return pub.Update(event)
		})
	}
}
