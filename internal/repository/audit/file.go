package audit

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/paulwwyvern/urlshortener/internal/model"
)

type LogFile struct {
	f *os.File

	encoder *json.Encoder
}

func NewAuditLogFile(file string) (*LogFile, error) {
	auditLog := &LogFile{}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	auditLog.f = f
	auditLog.encoder = json.NewEncoder(f)

	return auditLog, err
}

func (a *LogFile) Update(event *model.AuditEvent) error {
	err := a.encoder.Encode(event)
	if err != nil {
		return fmt.Errorf("Audit log file error: %w ", err)
	}
	return nil
}

func (a *LogFile) Close() error {
	return a.f.Close()
}
