package handlers

import "github.com/gofiber/fiber/v2"

func GetTest(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
