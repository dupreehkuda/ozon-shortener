package memory

import (
	"context"

	er "github.com/dupreehkuda/ozon-shortener/internal/customErrors"
)

// AddNewLink checks if link is present and if not adds it to the storage
func (s storage) AddNewLink(_ context.Context, id, link string) (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.shortToFull[id]; ok {
		return "", er.ErrExistingToken
	}

	if token, ok := s.fullToShort[link]; ok {
		return token, nil
	}

	s.shortToFull[id] = link
	s.fullToShort[link] = id

	return id, nil
}
