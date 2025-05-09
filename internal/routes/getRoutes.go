package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/timecrunch101/gopennypath/internal/database"
	"github.com/timecrunch101/gopennypath/internal/handlers"
)

func InitGetRoutes(app *fiber.App, db *database.DatabaseService) {

	GetHandler := handlers.NewUserHandler(db)

	getRouter := app.Group("/api/v1/get")

	getRouter.Get("/", GetHandler.GetUsers)

}
