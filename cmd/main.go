package main

import (
	"anime/internal/apisever"
	"anime/internal/config"
	"anime/internal/transport/rest"
	"log"
)

func main() {
	cfg := config.NewConfig("")

	router := rest.New()

	server := apisever.New(cfg, router)

	if err := server.Run(); err != nil {
		log.Fatal("server not start!")
	}
}
