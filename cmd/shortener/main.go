package main

import (
	"log"

	"github.com/TheMadman48L/shortener/internal/server"
	"github.com/TheMadman48L/shortener/internal/shortener"
	"github.com/TheMadman48L/shortener/internal/storage"
)

func main() {
	store := storage.NewSlowpokeStore("db/data.db")
	shorten := shortener.NewShortener(store, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	serv := server.NewServer(shorten)

	log.Fatalln(serv.Run())
}
