package controller

import (
	"errors"
	"fmt"
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

// @Summary			Consultar dados dos heróis.
// @Description		Consultar dados dos heróis..
// @Tags 			Hero
// @Accept 			json
// @Produce			json
// @Success			200 {object} model.Hero
// @Router			/api/secured/heros [get]
// @securityDefinitions.apiKey authorization
// @in header
// @name Authorization
// @Security JWT
func (h *HeroController) GetHeros(ctx *gin.Context) {
	heros, err := h.heroHandler.GetHeros()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, heros)
}

// @Summary			Registrar dados do herói.
// @Description		Registrar dados do herói.
// @Tags 			Hero
// @Accept 			json
// @Produce			json
// @Param 			hero body model.Hero true "hero"
// @Success			201 {object} string
// @Router			/api/secured/hero [post]
// @securityDefinitions.apiKey authorization
// @in header
// @name Authorization
// @Security JWT
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

// @Summary			Buscar dados do herói pelo ID.
// @Description		Buscar dados do herói pelo ID.
// @Tags 			Hero
// @Accept 			json
// @Produce			json
// @Param 			id path  int true "ID do herói"
// @Success			200 {object} model.Hero
// @Router			/api/secured/hero/{id} [get]
// @securityDefinitions.apiKey authorization
// @in header
// @name Authorization
// @Security JWT
func (h *HeroController) GetByIdHero(ctx *gin.Context) {
	id, shouldReturn := CheckId(ctx)
	if shouldReturn {
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

// @Summary			Atualizar dados do herói.
// @Description		Atualizar dados do herói.
// @Tags 			Hero
// @Accept 			json
// @Produce			json
// @Param 			id path  int true "ID do herói"
// @Param 			hero body model.Hero true "hero"
// @Success			200 {object} string
// @Router			/api/secured/hero/{id} [put]
// @securityDefinitions.apiKey authorization
// @in header
// @name Authorization
// @Security JWT
func (h *HeroController) UpdateHero(ctx *gin.Context) {
	var updateHero model.Hero
	id, shouldReturn := CheckId(ctx)

	if shouldReturn {
		return
	}

	if err := ctx.BindJSON(&updateHero); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	err := h.heroHandler.UpdateHero(id, updateHero)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, "Herói atualizado com sucesso!")
}

// @Summary			Deletar dados do herói.
// @Description		Deletar dados do herói.
// @Tags 			Hero
// @Accept 			json
// @Produce			json
// @Param 			id path  int true "ID do herói"
// @Success			200 {object} string
// @Router			/api/secured/hero/{id} [delete]
// @securityDefinitions.apiKey authorization
// @in header
// @name Authorization
// @Security JWT
func (h *HeroController) DeleteHero(ctx *gin.Context) {
	id, shouldReturn := CheckId(ctx)
	if shouldReturn {
		return
	}

	heroName, err := h.heroHandler.DeleteHero(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Herói %s deletado com sucesso!", heroName))
}

func CheckId(ctx *gin.Context) (int, bool) {
	param := ctx.Param("id")
	if param == "" {
		res := model.Response{
			Message: "Parâmetro inválido",
		}
		ctx.JSON(http.StatusBadRequest, res)
		return 0, true
	}
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		res := model.Response{
			Message: "ID inválido",
		}
		ctx.JSON(http.StatusBadRequest, res)
		return 0, true
	}
	return id, false
}
