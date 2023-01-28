package main

import (
	"log"
	"os"
	"io"
	"time"

	"github.com/jjimgo/go_gin_mysql/config"
	"github.com/jjimgo/go_gin_mysql/server"
	"github.com/gin-gonic/gin"
)

func main () {

	gin.DisableConsoleColor()

	t := time.Now()
	startTime := t.Format("2006-01-02 15:04:05")
	logFile := "logs/server_log -" + startTime

	f, err := os.Create(logFile)

	gin.DefaultWriter = io.MultiWriter(f)

	if err != nil {
		log.Fatal(err)
	}

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