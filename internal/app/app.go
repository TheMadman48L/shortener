package app

import (
	"io"
	"math/rand"
	"net/http"
	"strings"

	"github.com/recoilme/slowpoke"
)

var file = "db/data.db"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ShortenerHandler(w http.ResponseWriter, r *http.Request) {
	if path := r.URL.Path; r.Method == http.MethodPost && path == "/" {
		CreateLinkHandler(w, r)
	} else if r.Method == http.MethodGet && len(strings.Trim(path, "/")) == 7 {
		GetLinkHandler(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func CreateLinkHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	hash := shorting()
	for CheckKey([]byte(hash)) {
		hash = shorting()
	}
	SaveToDB(hash, string(body))
	_, err = w.Write([]byte("http://localhost:8080/" + hash))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetLinkHandler(w http.ResponseWriter, r *http.Request) {
	hash := strings.Trim(r.URL.Path, "/")
	url, err := GetFromDB(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func CheckKey(key []byte) bool {
	_, err := slowpoke.Get(file, key)
	return err == nil
}

func SaveToDB(hash, url string) {
	k := []byte(hash)
	v := []byte(url)

	if !CheckKey(k) {
		slowpoke.Set(file, k, v)
	}
}

func GetFromDB(hash string) (string, error) {
	k := []byte(hash)
	url, err := slowpoke.Get(file, k)
	return string(url), err
}

func shorting() string {
	b := make([]byte, 7)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
