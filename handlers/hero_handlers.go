package handlers

import (
	"log"

	"gorm.io/gorm"

	"api-gin/model"
)

type HeroHandler struct {
	DB *gorm.DB
}

func NewHeroHandler(db *gorm.DB) HeroHandler {
	return HeroHandler{db}
}

func (h HeroHandler) GetHeros() ([]model.Hero, error) {
	var heros []model.Hero
	if result := h.DB.Find(&heros); result.Error != nil {
		log.Println("Não foi possível buscar os heróis: ", result.Error)
		return []model.Hero{}, result.Error
	}
	return heros, nil
}

func (h HeroHandler) GetByIdHero(id int) (*model.Hero, error) {
	var hero model.Hero
	if result := h.DB.First(&hero, id); result.Error != nil {
		log.Println("Não foi possível buscar o herói: ", result.Error)
		return nil, result.Error
	}

	return &hero, nil
}

func (h HeroHandler) CreateHero(hero model.Hero) error {
	if result := h.DB.Create(&hero); result.Error != nil {
		log.Println("Não foi possível cadastrar o herói: ", result.Error)
		return result.Error
	}
	return nil
}

