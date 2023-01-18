package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jjimgo/go_gin_mysql/config"
)

type Server struct {
	config config.Config
	router *gin.Engine
}


func NewServer(config config.Config) (*Server, error) {
	// config, err := config.LoadConfig("../")

	server := &Server{config : config}
	server.setUpRouter()

	return server, nil
}

func (server *Server) setUpRouter() {
	router:= gin.Default()

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}