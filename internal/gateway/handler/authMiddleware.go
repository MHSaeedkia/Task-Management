package handler

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
	if c.Method() != fiber.MethodGet {
		apiKey := c.Get("X-API-Key")
		if apiKey != "12345" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}
	}
	return c.Next()
}
