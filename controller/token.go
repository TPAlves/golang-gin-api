package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gin/auth"
	"api-gin/handlers"
	"api-gin/model"
)

type TokenController struct {
	tokenHandler handlers.TokenHandler
}

func NewTokenController(tokenHandler handlers.TokenHandler) TokenController {
	return TokenController{tokenHandler}
}

// @Summary			Gerar Token.
// @Description		Gerar Token.
// @Tags 			User
// @Accept 			json
// @Produce			json
// @Param 			token body model.TokenRequest true "token"
// @Success			200 {object} string
// @Router			/api/token/ [post]
func (t *TokenController) GenerateToken(ctx *gin.Context) {
	var request model.TokenRequest
	var user model.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	err := t.tokenHandler.CheckEmailAndPassword(&request, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)

	if credentialError != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "senha inválida"})
		ctx.Abort()
		return
	}

	tokenString, err := auth.CreateTOken(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
