package handlers

import (
	"gorm.io/gorm"

	"api-gin/model"
)

type TokenHandler struct {
	DB *gorm.DB
}

func NewTokenHandler(db *gorm.DB) TokenHandler {
	return TokenHandler{db}
}

func (t *TokenHandler) CheckEmailAndPassword(request *model.TokenRequest, user *model.User) error {
	if err := t.DB.Where("email = ?", request.Email).First(&user); err.Error != nil {
		return err.Error
	}
	return nil
}
