package main

import (
	"github.com/3110Y/reminder/pkg/Domain/models"
	"github.com/3110Y/reminder/pkg/infrastructure/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := database.NewProviderGorm()
	connection, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	err = connection.AutoMigrate(
		models.Event{},
	)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
