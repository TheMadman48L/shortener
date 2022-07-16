package slow

import (
	"github.com/recoilme/slowpoke"
)

type slowstorage struct {
}

var file = "db/data.db"

func New() *slowstorage {
	return &slowstorage{}
}

func (s slowstorage) Get(hash string) (string, error) {
	k := []byte(hash)
	url, err := slowpoke.Get(file, k)
	return string(url), err
}

func (s slowstorage) Set(hash, url string) {
	k := []byte(hash)
	v := []byte(url)

	slowpoke.Set(file, k, v)
}
