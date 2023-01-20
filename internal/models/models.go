package models

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Link string `json:"Link"`
}
