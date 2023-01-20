package handlers

import (
	"go.uber.org/zap"

	i "github.com/dupreehkuda/ozon-shortener/internal/interfaces"
)

// handlers provides http-handlers for service
type handlers struct {
	service i.Service
	logger  *zap.Logger
}

// New creates new instance of handlers
func New(service i.Service, logger *zap.Logger) *handlers {
	return &handlers{
		service: service,
		logger:  logger,
	}
}
