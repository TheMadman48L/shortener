package main

import (
	"log"

	"github.com/TheMadman48L/shortener/internal/server"
	"github.com/TheMadman48L/shortener/internal/shortener"
	"github.com/TheMadman48L/shortener/internal/storage"
)

func main() {
	store := storage.NewSlowpokeStore()
	shorten := shortener.NewShortener(store)
	serv := server.NewServer(shorten)

	log.Fatalln(serv.Run())
}
