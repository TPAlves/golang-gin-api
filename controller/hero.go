package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"api-gin/handlers"
	"api-gin/model"
)

type HeroController struct {
	heroHandler handlers.HeroHandler
}

func NewHeroController(heroHandler handlers.HeroHandler) HeroController {
	return HeroController{
		heroHandler,
	}
}

func (h *HeroController) GetHeros(ctx *gin.Context) {
	heros, err := h.heroHandler.GetHeros()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, heros)
}

func (h *HeroController) CreateHeros(ctx *gin.Context) {
	var hero model.Hero
	err := ctx.BindJSON(&hero)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.heroHandler.CreateHero(hero)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, "Herói cadastrado com sucesso!")
}

func (h *HeroController) GetByIdHero(ctx *gin.Context) {
	param := ctx.Param("id")
	if param == "" {
		res := model.Response{
			Message: "Parâmetro inválido",
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		res := model.Response{
			Message: "ID inválido",
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	hero, err := h.heroHandler.GetByIdHero(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		res := model.Response{
			Message: "Herói não localizado",
		}
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, hero)
}
