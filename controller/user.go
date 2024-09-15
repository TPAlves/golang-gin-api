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

// @Summary			Cadastrar usuário.
// @Description		Cadastrar usuário.
// @Tags 			User
// @Accept 			json
// @Produce			json
// @Param 			user body model.User true "user"
// @Success			201 {object} string
// @Router			/api/secured/user/register [post]
// @securityDefinitions.apiKey authorization
// @in header
// @name Authorization
// @Security JWT
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
	ctx.JSON(http.StatusCreated, gin.H{"email": user.Email, "username": user.Username})
}
