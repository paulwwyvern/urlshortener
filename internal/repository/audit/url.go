package audit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/paulwwyvern/urlshortener/internal/model"
)

type LogUrl struct {
	url string

	client *http.Client
}

func NewAuditLogUrl(url string) *LogUrl {
	auditLog := &LogUrl{url: url}
	auditLog.client = &http.Client{
		Timeout: time.Second * 5,
	}

	return auditLog
}

func (a *LogUrl) Update(event *model.AuditEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("Audit log url: failed to marshal audit event: %w ", err)
	}

	req, err := http.NewRequest(http.MethodPost, a.url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("Audit log url: failed to create new http request: %w ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("Audit log url: failed to send http request: %w ", err)
	}

	defer resp.Body.Close()
	return nil
}

func (a *LogUrl) Close() error {
	return nil
}
