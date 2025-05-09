package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/timecrunch101/gopennypath/internal/database"
)

type GetHandler struct {
	DB *database.DatabaseService
}

func NewUserHandler(db *database.DatabaseService) *GetHandler {
	return &GetHandler{DB: db}
}

func (h *GetHandler) GetUsers(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
