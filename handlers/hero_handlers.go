package handlers

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"api-gin/model"
)

const (
	USER_NOT_FOUND = "usuário não localizado"
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

func (h HeroHandler) UpdateHero(id int, updateHero model.Hero) error {
	if hero, err := h.GetByIdHero(id); err == nil {
		hero.Name = updateHero.Name
		hero.SuperPower = updateHero.SuperPower
		hero.Group = updateHero.Group
		h.DB.Save(&hero)
		return nil
	}
	return errors.New(USER_NOT_FOUND)
}

func (h HeroHandler) DeleteHero(id int) (string, error) {
	if hero, err := h.GetByIdHero(id); err == nil {
		h.DB.Delete(&hero)
		return hero.Name, nil
	}
	return "", errors.New(USER_NOT_FOUND)
}
