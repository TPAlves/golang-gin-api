package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"api-gin/controller"
	"api-gin/db"
	_ "api-gin/docs"
	"api-gin/handlers"
)

const (
	HERO_ENDPOINT = "/hero/:id"
)

//	@title			Hero API
//	@version		1.0
//	@description	API Criada com o framework gin

// @host		localhost:8080
// @BasePath	/
func main() {
	server := gin.Default()
	dbConnection := db.ConnectDB()
	heroHandler := handlers.NewHeroHandler(dbConnection)
	heroController := controller.NewHeroController(heroHandler)

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/heros", heroController.GetHeros)
	server.POST("/hero", heroController.CreateHeros)
	server.GET(HERO_ENDPOINT, heroController.GetByIdHero)
	server.DELETE(HERO_ENDPOINT, heroController.DeleteHero)
	server.PUT(HERO_ENDPOINT, heroController.UpdateHero)
	server.Run(":8080")
}
