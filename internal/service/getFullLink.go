package service

import (
	"context"

	er "github.com/dupreehkuda/ozon-shortener/internal/customErrors"
)

// GetFullLink handles logic for redirecting operation
func (s service) GetFullLink(ctx context.Context, id string) (string, error) {
	link, err := s.storage.GetFullLink(ctx, id)
	if err != nil {
		if err == er.ErrNoSuchURL {
			return "", err
		}

		s.logger.Error("Error occurred in call to storage")
		return "", err
	}

	return link, nil
}
