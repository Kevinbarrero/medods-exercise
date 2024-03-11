package main

import (
	"log"
	"medods-exercise/api"
	"medods-exercise/db"
	"medods-exercise/utils"
)

func main() {
	db.Setup()
	runServer()
}

func runServer() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	server := api.NewServer(config)
	err = server.Start(config.HTTPADDRESS)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
