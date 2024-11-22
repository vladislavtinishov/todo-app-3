package main

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"todo_app_3/config/drivers"
	todomodels "todo_app_3/modules/todo/models"
	usermodels "todo_app_3/modules/users/models"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connector := drivers.Mysql{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		DB:   os.Getenv("DB_NAME"),
	}

	DB, err = connector.Connect()
	if err != nil {
		log.Fatalf("Error connecting to db: %s", err.Error())
	}
}

func main() {
	DB.AutoMigrate(usermodels.User{})
	DB.AutoMigrate(todomodels.Todo{})
}
