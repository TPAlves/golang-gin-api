package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api-gin/model"
)

func LoadEnviroment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Falha ao carregar o doenv")
	}
}

func ConnectDB() *gorm.DB {
	LoadEnviroment()

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Panic("Não foi possível conectar no banco de dados: ", err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Hero{})
	return db
}
