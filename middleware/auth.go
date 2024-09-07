package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gin/auth"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "requisição não possui token"})
			ctx.Abort()
			return
		}
		err := auth.VerifyToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
