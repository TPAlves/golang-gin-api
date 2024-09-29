package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"api-gin/model"
)

type MockHeroController struct {
	mock.Mock
}

var testHeroController *MockHeroController
var testModelHero *model.Hero
var testHeros []*model.Hero

func SetRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestGetAllHeros(t *testing.T) {
	t.Run("busca todos os herois", func(t *testing.T) {
		r := SetRouter()
		r.GET("/heros", testHeroController.GetHeros)
		req, _ := http.NewRequest("GET", "/heros", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		err := json.Unmarshal(w.Body.Bytes(), &testHeros)
		if err != nil {
			return
		}
		assert.NotEmpty(t, testHeros)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestGetByIdHero(t *testing.T) {
	getHeroByID := func(t *testing.T, heroId int) (w *httptest.ResponseRecorder) {
		t.Helper()
		r := SetRouter()
		r.GET("/hero/:id", testHeroController.GetByIdHero)
		req, _ := http.NewRequest("GET", fmt.Sprintf("/hero/%d", heroId), nil)
		w = httptest.NewRecorder()
		r.ServeHTTP(w, req)

		err := json.Unmarshal(w.Body.Bytes(), &testModelHero)
		if err != nil {
			return nil
		}
		return w
	}
	t.Run("busca um heroi existente", func(t *testing.T) {
		heroId := 1
		w := getHeroByID(t, heroId)
		assert.NotEmpty(t, testModelHero)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, heroId, testModelHero.ID)
	})

	t.Run("busca um heroi inexistente", func(t *testing.T) {
		heroId := 11518548
		w := getHeroByID(t, heroId)
		assert.Empty(t, testModelHero)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func (m *MockHeroController) GetHeros(ctx *gin.Context) {
	mockHeroes := []model.Hero{
		{
			ID:         4,
			Name:       "Homem de Ferro",
			SuperPower: "Inteligência e armaduras",
			Group:      "Vingadores",
		},
		{
			ID:         1,
			Name:       "Superman",
			SuperPower: "É o Superman!!",
			Group:      "Liga da Justiça",
		},
	}
	ctx.JSON(http.StatusOK, mockHeroes)
}

func (m *MockHeroController) GetByIdHero(ctx *gin.Context) {
	reqID := ctx.Param("id")
	switch reqID {
	case "1":
		mockHero := model.Hero{
			ID:         1,
			Name:       "Superman",
			SuperPower: "É o Superman!!",
			Group:      "Liga da Justiça",
		}
		ctx.JSON(http.StatusOK, mockHero)
	default:
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
}
