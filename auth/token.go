package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateTOken(userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userName,
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Panicf("Falha ao criar o token: %s", err)
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

	if err != nil {
		return fmt.Errorf("falha ao validar o token: %s", err)
	}

	if !token.Valid {
		return fmt.Errorf("token inválido")
	}

	return nil
}
