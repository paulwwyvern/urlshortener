package audit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestLogUrl_Update(t *testing.T) {
	url := "http://example.com/abcdef"

	event := model.AuditEvent{
		TS:     123456,
		Action: "follow",
		UserID: 654321,
		URL:    "https://www.example.com",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/abcdef", r.URL.String())

		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		actualEvent := model.AuditEvent{}
		err := json.NewDecoder(r.Body).Decode(&actualEvent)
		assert.NoError(t, err)

		assert.Equal(t, event, actualEvent)

		w.WriteHeader(http.StatusOK)

	}))
	defer server.Close()

	auditLog := NewAuditLogUrl(url)
	auditLog.client = server.Client()

	err := auditLog.Update(&event)
	assert.NoError(t, err)

}
