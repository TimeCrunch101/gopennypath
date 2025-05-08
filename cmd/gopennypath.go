package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/timecrunch101/gopennypath/internal/routes"
)

func main() {
	app := fiber.New()

	routes.InitGetRoutes(app)

	app.Listen(":8080")
}
