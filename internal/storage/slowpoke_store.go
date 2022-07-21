package storage

import "github.com/recoilme/slowpoke"

type slowpokeStore struct {
}

var file = "db/data.db"

func NewSlowpokeStore() *slowpokeStore {
	return &slowpokeStore{}
}

func (s slowpokeStore) Get(hash string) (string, error) {
	k := []byte(hash)
	url, err := slowpoke.Get(file, k)
	return string(url), err
}

func (s slowpokeStore) Set(hash, url string) error {
	k := []byte(hash)
	v := []byte(url)

	err := slowpoke.Set(file, k, v)
	if err != nil {
		return err
	}
	return nil
}
