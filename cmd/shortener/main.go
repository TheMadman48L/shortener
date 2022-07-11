package main

import (
	"net/http"

	"github.com/TheMadman48L/shortener/internal/app"
)

func main() {
	http.HandleFunc("/", app.ShortenerHandler)
	http.ListenAndServe(":8080", nil)
}
