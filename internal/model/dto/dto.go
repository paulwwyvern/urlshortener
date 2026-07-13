package dto

type GenerateURLJsonResponse struct {
	Result string `json:"result"`
}

type GetUserURLResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
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
