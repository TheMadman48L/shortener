package storage

import "github.com/recoilme/slowpoke"

type slowpokeStore struct {
	filePath string
}

func NewSlowpokeStore(fp string) *slowpokeStore {
	return &slowpokeStore{filePath: fp}
}

func (s *slowpokeStore) Get(hash string) (string, error) {
	k := []byte(hash)
	url, err := slowpoke.Get(s.filePath, k)
	return string(url), err
}

func (s *slowpokeStore) Set(hash, url string) error {
	k := []byte(hash)
	v := []byte(url)

	err := slowpoke.Set(s.filePath, k, v)
	if err != nil {
		return err
	}
	return nil
}
