package repo

import (
	"errors"

	"github.com/TheMadman48L/shortener/internal/repo/slow"
)

type Options struct {
	Env string
}

var ErrEmptyOptions = errors.New("empty options")
var ErrInvalidOptions = errors.New("invalid options")

type Repo interface {
	Get(hash string) (string, error)
	Set(hash, url string)
}

func New(opts *Options) (Repo, error) {
	if opts == nil {
		return nil, ErrEmptyOptions
	}

	switch opts.Env {
	case "dev":
		return slow.New(), nil
	default:
		return nil, ErrInvalidOptions
	}
}
