package utils

import "github.com/gofiber/fiber/v3"

func ResponseJSON(c fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
