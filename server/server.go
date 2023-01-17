package main

import (
	"fmt"
	"log"

	"github.com/jjimgo/go_gin_mysql/config"
)

func main () {
	config, err := config.LoadConfig("../")

	if err != nil{
		log.Fatal("connect load config", err)
	}

	fmt.Println(config.ServerAddress)
}