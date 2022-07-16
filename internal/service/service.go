package service

import (
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Storage interface {
	Get(hash string) (string, error)
	Set(hash, url string)
}

type service struct {
	storage Storage
}

func New(storage Storage) *service {
	return &service{storage: storage}
}

func (s *service) Shorting(url string) string {
	hash := s.getHash()
	for s.checkKey(hash) {
		hash = s.getHash()
	}
	s.storage.Set(hash, url)
	return hash
}

func (s *service) GetFullURL(hash string) (string, error) {
	url, err := s.storage.Get(hash)
	if err != nil {
		return url, err
	}
	return url, nil
}

func (s *service) getHash() string {
	b := make([]byte, 7)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (s *service) checkKey(hash string) bool {
	_, err := s.storage.Get(hash)
	return err == nil
}
