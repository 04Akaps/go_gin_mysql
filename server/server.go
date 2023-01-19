package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jjimgo/go_gin_mysql/config"
	"github.com/jjimgo/go_gin_mysql/controllers"
)

type Server struct {
	config config.Config
	router *gin.Engine
}


func NewServer(config config.Config) (*Server, error) {
	server := &Server{config : config}
	server.setUpRouter()

	return server, nil
}

func (server *Server) setUpRouter() {
	router:= gin.Default()
	server.router = router
}

func setTestRoute(router *gin.Engine) {
	testRoutes := router.Group("/test")

	testRoutes.GET("/getTestHello", func(ctx *gin.Context) {
		controllers.GetTestHello(ctx)
	})


	testRoutes.GET("/getTest", func(ctx *gin.Context) {
		controllers.GetTest(ctx)
	})

	testRoutes.POST("/makeTest", func(ctx *gin.Context) {
		controllers.CreateTest(ctx)
	})


}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
