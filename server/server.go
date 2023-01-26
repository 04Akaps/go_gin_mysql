package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jjimgo/go_gin_mysql/config"
	"github.com/jjimgo/go_gin_mysql/db"
	sqlc "github.com/jjimgo/go_gin_mysql/db/sqlc"
)

type Server struct {
	config config.Config
	router *gin.Engine
	query *sqlc.Queries
}

func NewServer(config config.Config) (*Server, error) {
	server := &Server{config : config}
	server.setUpRouter()
	query := db.MigrateDataBase()

	server.query = query

	return server, nil
}

func (server *Server) setUpRouter() {
	router:= gin.Default()

	server.setTestRoute(router) // sample Test Router
	server.setUserRoute(router) // sample User Router
	server.setDiaryRoute(router) // sample Tx Router

	server.router = router
}

func (server *Server)  setDiaryRoute(router *gin.Engine) {
	diary := router.Group("/diary")

	diary.GET("/getDiary")
	diary.POST("/createDiary")
	diary.DELETE("/deleteDiary")
}

func (server *Server) setUserRoute(router *gin.Engine) {
	userRoutes := router.Group("/user")

	userRoutes.GET("/getUser", server.getUser)
	userRoutes.GET("/getUsers", server.getUsers)
	userRoutes.POST("/createUser", server.createAccount)
	userRoutes.DELETE("/deleteUser", server.deleteUser)
}

func(server *Server) setTestRoute(router *gin.Engine) {
	testRoutes := router.Group("/test")

	testRoutes.GET("/getTestHello", server.getTestHello)
	testRoutes.GET("/getTest", server.getTest)
	testRoutes.POST("/makeTest", server.createTest)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
