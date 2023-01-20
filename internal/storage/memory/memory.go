package memory

import "sync"

type storage struct {
	mtx       sync.RWMutex
	shortened map[string]string
}

func New() *storage {
	return &storage{shortened: map[string]string{}}
}
