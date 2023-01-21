package models

// ShortenRequest provides URL for shortening
type ShortenRequest struct {
	URL string `json:"url"`
}

// ShortenResponse provides shortened Link
type ShortenResponse struct {
	Link string `json:"Link"`
}
