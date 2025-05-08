package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/timecrunch101/gopennypath/internal/handlers"
)

func InitGetRoutes(app *fiber.App) {
	getRouter := app.Group("/api/v1/get")

	getRouter.Get("/", handlers.GetTest)

}
