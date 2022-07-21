package shortener

import (
	"math/rand"
)

type Storage interface {
	Get(hash string) (string, error)
	Set(hash, url string) error
}

type shortener struct {
	store Storage
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewShortener(store Storage) *shortener {
	return &shortener{store: store}
}

func (s *shortener) ShortingURL(url string) string {
	hash := s.getHash()
	for s.checkKey(hash) {
		hash = s.getHash()
	}
	s.store.Set(hash, url)
	return hash
}

func (s *shortener) GetFullURL(hash string) (string, error) {
	url, err := s.store.Get(hash)
	if err != nil {
		return url, err
	}
	return url, nil
}

func (s *shortener) getHash() string {
	b := make([]byte, 7)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (s *shortener) checkKey(hash string) bool {
	_, err := s.store.Get(hash)
	return err == nil
}
