package memory

import "sync"

type storage struct {
	mtx         sync.Mutex
	shortToFull map[string]string
	fullToShort map[string]string
}

func New() *storage {
	return &storage{
		mtx:         sync.Mutex{},
		shortToFull: make(map[string]string, 32),
		fullToShort: make(map[string]string, 32),
	}
}
