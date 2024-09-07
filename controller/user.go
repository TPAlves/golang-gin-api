package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gin/handlers"
	"api-gin/model"
)

type UserCOntroller struct {
	userHandler handlers.UserHandler
}

func NewUserController(userHandler handlers.UserHandler) UserCOntroller {
	return UserCOntroller{
		userHandler,
	}
}

func (u *UserCOntroller) RegisterUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	err := u.userHandler.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"ID": user.ID, "email": user.Email, "username": user.Username})

}
