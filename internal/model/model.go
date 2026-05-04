package model

type GenerateURLJsonRequest struct {
	URL string `json:"url"`
}

type GenerateURLJsonResponse struct {
	Result string `json:"result"`
}

type URL struct {
	ID          string `json:"id,omitempty"`
	ShortUrl    string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
}
