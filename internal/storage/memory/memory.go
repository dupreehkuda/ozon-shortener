package memory

import (
	"sync"

	"go.uber.org/zap"
)

type storage struct {
	mtx         sync.Mutex
	logger      *zap.Logger
	shortToFull map[string]string
	fullToShort map[string]string
}

func New(logger *zap.Logger) *storage {
	logger.Info("Launched with memory")

	return &storage{
		mtx:         sync.Mutex{},
		logger:      logger,
		shortToFull: make(map[string]string, 32),
		fullToShort: make(map[string]string, 32),
	}
}
