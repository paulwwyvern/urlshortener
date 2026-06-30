package model

type URL struct {
	ID          string `json:",omitempty"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	IsExist     bool   `json:"-"`
}

type URLFile struct {
	ID          string `json:"id,omitempty"`
	UserID      int32  `json:"user_id"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type AuditEvent struct {
	TS     int64  `json:"ts"`
	Action string `json:"action"`
	UserID int32  `json:"user_id"`
	URL    string `json:"url"`
}
