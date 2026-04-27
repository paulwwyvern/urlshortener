package model

type GenerateUrlJsonRequest struct {
	Url string `json:"url"`
}

type GenerateUrlJsonResponse struct {
	Result string `json:"result"`
}

type Url struct {
	Id          string `json:"id,omitempty"`
	ShortUrl    string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
}
