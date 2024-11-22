package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"todo_app_3/config/drivers"
	"todo_app_3/handlers"
)

func Run() {
	fmt.Println("Starting project!")
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

	db, err := connector.Connect()
	if err != nil {
		log.Fatalf("Error connecting to db: %s", err.Error())
	}

	handler := handlers.NewHandler(db)

	err = handler.Run()
	if err != nil {
		log.Fatalf("Error starting router: %s", err.Error())
	}
}

func main() {
	Run()
}
