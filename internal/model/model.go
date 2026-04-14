package model

type GenerateUrlJsonRequest struct {
	Url string `json:"url"`
}

type GenerateUrlJsonResponse struct {
	Result string `json:"result"`
}
