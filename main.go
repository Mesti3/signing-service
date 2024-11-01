package main

import (
	"log"
	"net/http"

	"signing-service-challenge/api"
	"signing-service-challenge/persistence"
)

const (
	ListenAddress = ":8080"
	apis2         = "/api/v0/devices"
	apis1         = "/api/v0/devices/list"
)

func main() {
	store := persistence.NewInMemorys()
	server := api.NewServer(ListenAddress, store)

	deviceHandler := api.NewDeviceHandler(store)

	http.HandleFunc(apis2, deviceHandler.CreateDevice)
	http.HandleFunc(apis1, deviceHandler.CreateDevice)

	if err := server.Run(); err != nil {
		log.Fatal("Could not start server on ", ListenAddress)
	}
}
