package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Shortener interface {
	ShortingURL(url string) string
	GetFullURL(hash string) (string, error)
}

type server struct {
	r       *mux.Router
	shorten Shortener
}

var port = "8080"

var ErrEmptyBody = errors.New("empty body")
var ErrHashLen = errors.New("hash must be 7 letters")

func NewServer(shorten Shortener) *server {
	s := &server{
		r:       mux.NewRouter(),
		shorten: shorten,
	}
	s.r.HandleFunc("/", s.handleShortingURL).Methods("POST")
	s.r.HandleFunc("/{hash}", s.handleGetFullURL).Methods("GET")
	s.r.HandleFunc("/", s.handleBadRequest)
	return s
}

func (s *server) Run() error {
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	return http.ListenAndServe(":"+port, s.r)
}

func (s *server) handleShortingURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(body) == 0 {
		http.Error(w, ErrEmptyBody.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	hash := s.shorten.ShortingURL(string(body))
	shortURL := fmt.Sprintf("http://localhost:%s/%s", port, hash)
	_, err = w.Write([]byte(shortURL))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) handleGetFullURL(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]
	if len(hash) != 7 {
		http.Error(w, ErrHashLen.Error(), http.StatusBadRequest)
		return
	}

	url, err := s.shorten.GetFullURL(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (s *server) handleBadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}
