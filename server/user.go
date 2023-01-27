package server


import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	sqlc "github.com/jjimgo/go_gin_mysql/db/sqlc"
)

type createUserRequest struct {
	Email	string 	`json:"email" binding:"required,email"`
	Gender	string 	`json:"gender" binding:"required,oneof=Man Woman"`
	Age		int64	`json:"age" binding:"required,min=1"`
	Country	string 	`json:"country" binding:"required"`
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

	_, err := server.query.CreateUser(ctx, arg)

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

	ctx.JSON(http.StatusOK, arg)
}

type getDeleteUserRequest struct {
	Email 	string 	`uri:"email" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getDeleteUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	user, err := server.query.GetUser(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			// DB에 없다면
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) getAllUsers(ctx *gin.Context) {
	user, err := server.query.GetAllUsers(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			// DB에 없다면
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req getDeleteUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	err := server.query.DeleteUser(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			// DB에 없다면
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	ctx.JSON(http.StatusOK, "Deleted Email is : " + req.Email)
}