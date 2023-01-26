package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jjimgo/go_gin_mysql/config"
	"github.com/jjimgo/go_gin_mysql/controllers"
	"github.com/jjimgo/go_gin_mysql/db"
)

type Server struct {
	config config.Config
	router *gin.Engine
}

func NewServer(config config.Config) (*Server, error) {
	server := &Server{config : config}
	server.setUpRouter()
	db.MigrateDataBase()

	return server, nil
}

func (server *Server) setUpRouter() {
	router:= gin.Default()

	setTestRoute(router) // sample Test Router
	setUserRoute(router) // sample User Router
	setDiaryRoute(router) // sample Tx Router

	server.router = router
}

func setDiaryRoute(router *gin.Engine) {
	diary := router.Group("/diary")

	diary.GET("/getDiary")
	diary.POST("/createDiary")
	diary.DELETE("/deleteDiary")
}

func setUserRoute(router *gin.Engine) {
	userRoutes := router.Group("/user")

	userRoutes.GET("/getUser")
	userRoutes.GET("/getUsers")
	userRoutes.POST("/createUser")
	userRoutes.DELETE("/deleteUser")
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
