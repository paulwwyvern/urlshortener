package model

type URL struct {
	ID          string `json:",omitempty"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	IsExist     bool   `json:"-"`
}

type GenerateURLBatchRequest struct {
	ID          string `json:"correlation_id"`
	OriginalURL string `json:"original_url"`
}

type GenerateURLBatchResponse struct {
	ID       string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}

type GenerateURLJsonRequest struct {
	URL string `json:"url"`
}

type GenerateURLJsonResponse struct {
	Result string `json:"result"`
}

type URLFile struct {
	ID          string `json:"id,omitempty"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
