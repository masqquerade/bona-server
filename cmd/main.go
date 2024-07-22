package main

import (
	"bonaserver/pkg/server"
	"bonaserver/pkg/store"
	"log"
	"net/http"

	"github.com/gofor-little/env"
)

func main() {
	if err := env.Load("./.env"); err != nil {
		log.Fatal(err)
	}

	config, err := store.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	store, err := store.NewStore(config)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	server, err := server.NewServer(env.Get("TG__SECRET_TOCKEN", ""), store)

	if err != nil {
		log.Fatal(err)
	}

	handler := server.InitServer()

	log.Fatal(http.ListenAndServe(":8080", handler))
}
