package memory

import (
	"sync"

	"go.uber.org/zap"
)

// storage provides service's storing layer
type storage struct {
	mtx         sync.Mutex
	logger      *zap.Logger
	shortToFull map[string]string
	fullToShort map[string]string
}

// New creates a new instance of database layer using memory
func New(logger *zap.Logger) *storage {
	logger.Info("Launched with memory")

	return &storage{
		mtx:         sync.Mutex{},
		logger:      logger,
		shortToFull: make(map[string]string, 32),
		fullToShort: make(map[string]string, 32),
	}
}
