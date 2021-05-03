package utils

import (
	"log"
	"os"
	"user/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	godotenv_err := godotenv.Load()
	if godotenv_err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := gorm.Open(os.Getenv("databaseType"), os.Getenv("databaseName"))
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Base{})
	db.AutoMigrate(&models.User{})
	DB = db
}
