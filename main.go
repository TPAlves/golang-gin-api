package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"api-gin/controller"
	"api-gin/db"
	_ "api-gin/docs"
	"api-gin/handlers"
	"api-gin/middleware"
)

const (
	HERO_ENDPOINT = "/hero/:id"
)

// @title	Hero API
// @version 1.0
// @description API Criada com o framework gin
// @host localhost:8080
// @BasePath /
func main() {
	server := initRouter()
	server.Run(":8080")
}

func initRouter() *gin.Engine {
	server := gin.Default()
	api := server.Group("/api")
	dbConnection := db.ConnectDB()
	heroHandler := handlers.NewHeroHandler(dbConnection)
	tokenHandler := handlers.NewTokenHandler(dbConnection)
	userHandler := handlers.NewUserHandler(dbConnection)
	heroController := controller.NewHeroController(heroHandler)
	tokenController := controller.NewTokenController(tokenHandler)
	userController := controller.NewUserController(userHandler)
	{
		api.POST("/token", tokenController.GenerateToken)
		api.POST("user/register", userController.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
			secured.GET("/heros", heroController.GetHeros)
			secured.POST("/hero", heroController.CreateHeros)
			secured.GET(HERO_ENDPOINT, heroController.GetByIdHero)
			secured.DELETE(HERO_ENDPOINT, heroController.DeleteHero)
			secured.PUT(HERO_ENDPOINT, heroController.UpdateHero)
		}
	}
	return server
}
