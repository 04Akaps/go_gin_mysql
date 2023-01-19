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

	setTestRoute(router)
	server.router = router
}

func setUserRoute(router *gin.Engine) {

}

func setTxRoute(router *gin.Engine) {
	
}

func setTestRoute(router *gin.Engine) {
	testRoutes := router.Group("/test")

	testRoutes.GET("/getTestHello", controllers.GetTestHello)
	testRoutes.GET("/getTest",controllers.GetTest)
	testRoutes.POST("/makeTest", controllers.CreateTest)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
