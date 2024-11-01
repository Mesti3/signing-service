package main

import (
	"log"

	"signing-service-challenge/api"
	"signing-service-challenge/persistence"
)

const (
	ListenAddress = ":8080"
)

func main() {
	store := persistence.NewInMemorys()
	server := api.NewServer(ListenAddress, store)

	if err := server.Run(); err != nil {
		log.Fatal("Could not start server on ", ListenAddress)
	}
}
