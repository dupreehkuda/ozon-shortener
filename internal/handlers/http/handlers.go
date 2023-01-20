package handlers

import (
	"go.uber.org/zap"

	i "github.com/dupreehkuda/ozon-shortener/internal/interfaces"
)

type handlers struct {
	actions i.Service
	logger  *zap.Logger
}

// New creates new instance of handlers
func New(processor i.Service, logger *zap.Logger) *handlers {
	return &handlers{
		actions: processor,
		logger:  logger,
	}
}
