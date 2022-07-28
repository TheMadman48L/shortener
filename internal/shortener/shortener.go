package shortener

import (
	"math/rand"
)

type Storager interface {
	Get(hash string) (string, error)
	Set(hash, url string) error
}

type shortService struct {
	store   Storager
	symbols string
}

func NewShortener(store Storager, symbols string) *shortService {
	return &shortService{store: store, symbols: symbols}
}

func (s *shortService) ShortingURL(url string) string {
	hash := s.getHash()
	_, err := s.store.Get(hash)
	for err == nil {
		hash = s.getHash()
		_, err = s.store.Get(hash)
	}
	s.store.Set(hash, url)
	return hash
}

func (s *shortService) GetFullURL(hash string) (string, error) {
	url, err := s.store.Get(hash)
	if err != nil {
		return url, err
	}
	return url, nil
}

func (s *shortService) getHash() string {
	b := make([]byte, 7)
	for i := range b {
		b[i] = s.symbols[rand.Intn(len(s.symbols))]
	}
	return string(b)
}
