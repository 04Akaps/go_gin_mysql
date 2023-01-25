package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	e "github.com/jjimgo/go_gin_mysql/err"
)

type Test struct {
	Name 	string		`json:"name"`
	Num		int64		`json:"num"`
}

var newTestArea = []Test{
	{Name: "hojin", Num: 3},
	{Name: "hojin-2", Num: 4},
}


func CreateTest(ctx *gin.Context) {
	var newTest Test
	if err := ctx.ShouldBindJSON(&newTest); err != nil {
		ctx.JSON(http.StatusBadRequest, e.ErrorResponse(err))
	}

	newTestArea = append(newTestArea, newTest);
	
	ctx.IndentedJSON(http.StatusOK, newTest)
}

func GetTest(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, newTestArea)
}

func GetTestHello(ctx *gin.Context) { 
	ctx.JSON(http.StatusOK, "hello")
}