package app

import (
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Port = "8080"

type Shortener interface {
	Shorting(url string) string
	GetFullURL(hash string) (string, error)
}

type app struct {
	r       *mux.Router
	service Shortener
}

func New(service Shortener) *app {
	r := mux.NewRouter()

	return &app{service: service, r: r}
}

func (a *app) handleShortingURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(body) == 0 || r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url := string(body)
	hash := a.service.Shorting(url)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("http://localhost:" + Port + "/" + hash))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *app) handleGetFullURL(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]
	if len(hash) != 7 || r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := a.service.GetFullURL(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (a *app) Run() error {
	if len(os.Args) > 1 {
		Port = os.Args[1]
	}
	a.r.HandleFunc("/", a.handleShortingURL)
	a.r.HandleFunc("/{hash}", a.handleGetFullURL)
	return http.ListenAndServe(":"+Port, a.r)
}
