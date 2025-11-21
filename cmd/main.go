package main

import (
	"log"

	"github.com/joho/godotenv"
	api "github.com/lemeow125/goapi/internal/api"
	migrations "github.com/lemeow125/goapi/internal/migrations"
)
func main(){
	// Load .env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	// Migrations
	migrations.Setup()

	// Start API
	api.Run()
}