package audit

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestLogFile_UpdateNotExistFile(t *testing.T) {
	event := model.AuditEvent{
		TS:     123456,
		Action: "follow",
		UserID: 654321,
		URL:    "https://www.example.com",
	}

	dir := t.TempDir()

	path := filepath.Join(dir, "audit.log")

	auditLogFile, err := NewAuditLogFile(path)
	assert.NoError(t, err)

	err = auditLogFile.Update(&event)
	assert.NoError(t, err)

	content, err := os.ReadFile(path)
	assert.NoError(t, err)

	actualEvent := model.AuditEvent{}

	err = json.Unmarshal(content, &actualEvent)
	assert.NoError(t, err)

	assert.Equal(t, event, actualEvent)
}

func TestLogFile_UpdateExistFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "audit.log")

	events := []model.AuditEvent{
		{
			TS:     111,
			Action: "follow",
			UserID: 222,
			URL:    "https://www.example.com",
		},
		{
			TS:     333,
			Action: "abc",
			URL:    "https://www.google.com",
		},
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	assert.NoError(t, err)

	encode := json.NewEncoder(f)

	for _, event := range events {
		err := encode.Encode(event)
		assert.NoError(t, err)
	}
	err = f.Close()
	assert.NoError(t, err)

	addEvent := model.AuditEvent{
		TS:     123456,
		Action: "shorten",
		UserID: 654321,
		URL:    "https://www.audit.com",
	}

	events = append(events, addEvent)

	auditLogFile, err := NewAuditLogFile(path)
	assert.NoError(t, err)

	err = auditLogFile.Update(&addEvent)
	assert.NoError(t, err)

	content, err := os.ReadFile(path)
	assert.NoError(t, err)

	decoder := json.NewDecoder(bytes.NewReader(content))

	actualEvents := make([]model.AuditEvent, len(events))

	for i := 0; i < len(events); i++ {
		err = decoder.Decode(&actualEvents[i])
		assert.NoError(t, err)
	}
	dummyEvent := &model.AuditEvent{}
	err = decoder.Decode(dummyEvent)
	assert.ErrorIs(t, io.EOF, err)

	assert.Equal(t, events, actualEvents)
}
