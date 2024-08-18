package main

import (
	"github.com/gin-gonic/gin"

	"api-gin/controller"
	"api-gin/db"
	"api-gin/handlers"
)

func main() {
	server := gin.Default()
	dbConnection := db.ConnectDB()
	heroHandler := handlers.NewHeroHandler(dbConnection)
	heroController := controller.NewHeroController(heroHandler)

	server.GET("/heros", heroController.GetHeros)
	server.POST("/hero", heroController.CreateHeros)
	server.GET("/hero/:id", heroController.GetByIdHero)
	server.Run(":8080")
}
