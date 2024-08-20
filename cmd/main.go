package main

import (
	"github.com/gin-gonic/gin"

	"api-gin/controller"
	"api-gin/db"
	"api-gin/handlers"
)

const (
	HERO_ENDPOINT = "/hero/:id"
)

func main() {
	server := gin.Default()
	dbConnection := db.ConnectDB()
	heroHandler := handlers.NewHeroHandler(dbConnection)
	heroController := controller.NewHeroController(heroHandler)

	server.GET("/heros", heroController.GetHeros)
	server.POST("/hero", heroController.CreateHeros)
	server.GET(HERO_ENDPOINT, heroController.GetByIdHero)
	server.DELETE(HERO_ENDPOINT, heroController.DeleteHero)
	server.PUT(HERO_ENDPOINT, heroController.UpdateHero)
	server.Run(":8080")
}
