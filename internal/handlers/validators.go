package handlers

import (
	"regexp"

	er "github.com/dupreehkuda/ozon-shortener/internal/customErrors"
)

func ValidateURL(url string) error {
	if url == "" {
		return er.ErrEmptyRequest
	}

	pattern := `^(https?://|www.)?[a-zA-Z0-9-]{1,256}([.][a-zA-Z-]{1,256})?([.][a-zA-Z]{1,30})([/]?[a-zA-Z0-9/?=%&#_.-]+)`

	valid, err := regexp.Match(pattern, []byte(url))
	if err != nil {
		return err
	}

	if !valid {
		return er.ErrInvalidUrl
	}

	return nil
}

func ValidateToken(token string) error {
	if token == "" {
		return er.ErrEmptyRequest
	}

	pattern := `^[a-zA-Z0-9_]{10}$`

	valid, err := regexp.Match(pattern, []byte(token))
	if err != nil {
		return err
	}

	if !valid {
		return er.ErrInvalidToken
	}

	return nil
}
