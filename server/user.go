package server


import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	sqlc "github.com/jjimgo/go_gin_mysql/db/sqlc"
)

// `email` varchar(255) UNIQUE PRIMARY KEY,
// `gender` varchar(255) NOT NULL,
// `age` int NOT NULL,
// `country` int NOT NULL,
// `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP

type createUserRequest struct {
	Email	string 	`json:"email" binding:"required,email"`
	Gender	string 	`json:"gender" binding:"required,oneof=Man Woman"`
	Age		int64	`json:"age" binding:"required,min=1"`
	Country	string 	`json:"country" binding:"required"`
}

type createUserREsponse struct {

}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createUserRequest
	// check key is existed
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	arg := sqlc.CreateUserParams{
		Email : req.Email,
		Gender : req.Gender,
		Age : req.Age,
		Country : req.Country,
	}

	result, err := server.query.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) getUser(ctx *gin.Context) {

}


func (server *Server) getUsers(ctx *gin.Context) {

}


func (server *Server) deleteUser(ctx *gin.Context) {

}