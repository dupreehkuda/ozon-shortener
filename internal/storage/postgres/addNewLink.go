package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (s storage) AddNewLink(id, link string) (string, error) {
	conn, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.Error("Error while acquiring connection", zap.Error(err))
		return "", err
	}
	defer conn.Release()

	var token string

	err = conn.QueryRow(context.Background(), "SELECT id FROM links WHERE original = $1;", link).Scan(&token)
	if err != nil && err != pgx.ErrNoRows {
		s.logger.Error("Error occurred executing query", zap.Error(err))
		return "", err
	}

	if token != "" {
		return token, nil
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO links (original, id) VALUES ($1, $2);", link, id)
	if err != nil {
		s.logger.Error("Error occurred executing insert query", zap.Error(err))
		return "", err
	}

	return id, nil
}
