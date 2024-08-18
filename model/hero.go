package model

type Hero struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"Nome"`
	SuperPower string `json:"Super Poder"`
	Group      string `json:"Equipe"`
}

func (h Hero) TableName() string {
	return "hero"
}
