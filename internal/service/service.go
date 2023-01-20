package service

import (
	"go.uber.org/zap"

	i "github.com/dupreehkuda/ozon-shortener/internal/interfaces"
)

// service provides business-logic
type service struct {
	storage i.Storage
	logger  *zap.Logger
}

// New creates new instance of actions
func New(storage i.Storage, logger *zap.Logger) *service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}
