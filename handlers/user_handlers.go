package handlers

import (
	"log"

	"gorm.io/gorm"

	"api-gin/model"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) UserHandler {
	return UserHandler{db}
}

func (u UserHandler) CreateUser(user model.User) error {
	if result := u.DB.Create(&user); result.Error != nil {
		log.Println("Não foi possível cadastrar o usuário: ", result.Error)
		return result.Error
	}
	return nil
}
