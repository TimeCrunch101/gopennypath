package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/timecrunch101/gopennypath/internal/database"
	"github.com/timecrunch101/gopennypath/internal/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbService, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	app := fiber.New()

	routes.InitGetRoutes(app, dbService)

	app.Listen(":8080")
}
