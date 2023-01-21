package memory

import er "github.com/dupreehkuda/ozon-shortener/internal/customErrors"

// GetFullLink gets original link by shortened token
func (s storage) GetFullLink(id string) (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	link, ok := s.shortToFull[id]
	if !ok {
		return "", er.ErrNoSuchURL
	}

	return link, nil
}
