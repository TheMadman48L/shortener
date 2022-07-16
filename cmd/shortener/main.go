package main

import (
	"log"

	"github.com/TheMadman48L/shortener/internal/app"
	"github.com/TheMadman48L/shortener/internal/repo"
	"github.com/TheMadman48L/shortener/internal/service"
)

func main() {
	rep, err := repo.New(&repo.Options{Env: "dev"})
	if err != nil {
		log.Fatalln(err)
	}

	serv := service.New(rep)
	myapp := app.New(serv)
	log.Fatalln(myapp.Run())
}
