package main

import (
	"log"

	"github.com/jjimgo/go_gin_mysql/config"
	"github.com/jjimgo/go_gin_mysql/server"
)

func main () {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("connect load config : ", err)
	}

	server, err := server.NewServer(config)

	if err != nil {
		log.Fatal("cannot start server %w", err)
	}

	err = server.Start(config.ServerAddress)

	if err =  server.Start(config.ServerAddress); err != nil {
		log.Fatal("error to start server : ", err)
	}
}