package main

import (
	"log"

	"github.com/joho/godotenv"
	api "github.com/lemeow125/goapi/internal/api"
	migrations "github.com/lemeow125/goapi/internal/migrations"
)
func main(){
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Migrations
	db := migrations.Setup()

	// Start API
	api.Run(db)
}