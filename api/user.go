package api

import (
	"context"
	"medods-exercise/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createUserRequest struct {
	UserName string `json:"username" binding:"required"`
}

type userResponse struct {
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	newUUID, err := uuid.NewRandom()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := db.CreateUserParams{
		UUID:     newUUID,
		UserName: req.UserName,
	}
	_, err = server.store.Collection("users").InsertOne(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	rsp := userResponse{
		Username: req.UserName,
		UUID:     newUUID.String(),
	}
	ctx.JSON(http.StatusOK, rsp)

}
