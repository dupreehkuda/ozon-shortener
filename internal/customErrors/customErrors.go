package customErrors

import "errors"

var (
	ErrNoSuchURL    = errors.New("this token do not exist")
	ErrEmptyRequest = errors.New("request body is empty")
	ErrInvalidUrl   = errors.New("url is invalid")
	ErrInvalidToken = errors.New("url token is invalid")
)
