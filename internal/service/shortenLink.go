package service

import (
	"context"

	"go.uber.org/zap"
)

// ShortenLink handles logic for creating shortened link operation
func (s service) ShortenLink(ctx context.Context, link string) (string, error) {
	token, err := TokenGenerator()
	if err != nil {
		s.logger.Error("Error occurred generating token", zap.Error(err))
		return "", err
	}

	token, err = s.storage.AddNewLink(ctx, token, link)
	if err != nil {
		return "", err
	}

	return token, nil
}
