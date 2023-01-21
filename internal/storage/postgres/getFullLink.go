package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	er "github.com/dupreehkuda/ozon-shortener/internal/customErrors"
)

// GetFullLink gets original link by shortened token
func (s storage) GetFullLink(id string) (string, error) {
	conn, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return "", err
	}
	defer conn.Release()

	var link string

	err = conn.QueryRow(context.Background(), "SELECT original FROM links WHERE id = $1;", id).Scan(&link)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", er.ErrNoSuchURL
		}

		s.logger.Error("Error occurred executing query", zap.Error(err))
		return "", err
	}

	return link, nil
}
