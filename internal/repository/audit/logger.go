package audit

import (
	"encoding/json"
	"fmt"

	"github.com/paulwwyvern/urlshortener/internal/model"
	"go.uber.org/zap"
)

type LogLogger struct {
	logger *zap.Logger
}

func NewAuditLogLogger(logger *zap.Logger) *LogLogger {
	auditLog := &LogLogger{logger: logger}

	return auditLog
}

func (a *LogLogger) Update(event *model.AuditEvent) error {
	str, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("audit log logger: failed to marshal audit event: %w ", err)
	}
	a.logger.Info("Get audit event", zap.String("event", string(str)))

	return nil
}

func (a *LogLogger) Close() error {
	return nil
}
