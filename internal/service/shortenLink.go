package service

import (
	"go.uber.org/zap"
)

// ShortenLink handles logic for creating shortened link operation
func (s service) ShortenLink(link string) (string, error) {
	token, err := TokenGenerator()
	if err != nil {
		s.logger.Error("Error occurred generating token", zap.Error(err))
		return "", err
	}

	token, err = s.storage.AddNewLink(token, link)
	if err != nil {
		return "", err
	}

	return token, nil
}
