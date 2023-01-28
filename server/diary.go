package server


import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	sqlc "github.com/jjimgo/go_gin_mysql/db/sqlc"
)

type createDiaryRequest struct {
	Content	string 	`json:"content" binding:"required"`
	UserEmail	string 	`json:"user_email" binding:"required,email"`
}

func (server *Server) createDiary(ctx *gin.Context) {
	var req createDiaryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	arg := sqlc.CreateDiaryParams {
		Content : req.Content,
		UserEmail : req.UserEmail,
	}

	_, err := server.query.CreateDiary(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, arg)
}

type idUriRequest struct {
	ID 	int64 	`uri:"id" binding:"required,min=1"`
}

func (server *Server) getDiary(ctx *gin.Context) {
	var req idUriRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	diary, err := server.query.GetDiary(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, diary)
}

type getDiarysRequest struct {
	Email 	string 	`uri:"email" binding:"required,email"`
}

func (server *Server) getDiarys(ctx *gin.Context) {
	var req getDiarysRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	diary, err := server.query.GetDiarys(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, diary)
}

type updateDiaryRequest struct {
	Content	string 	`json:"content" binding:"required"`
	ID		int64 	`json:"id" binding:"required,min=1"`
}

type updateDiaryResponse struct {
	beforeContent 	string
	afterContent	string
	updatedId		int64
}

func (server *Server) updateDiary(ctx *gin.Context) {
	var req updateDiaryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	arg := sqlc.UpdateDiaryParams {
		Content : req.Content,
		ID : req.ID,
	}

	err := server.query.UpdateDiary(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, arg)
}

func (server *Server) deleteDiary(ctx *gin.Context) {
	var req idUriRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return;
	}

	err := server.query.DeleteDiary(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Header("Content-Type", "application/json")
	
	ctx.JSON(http.StatusOK, strconv.Itoa(int(req.ID))  + "is Deleted")
}