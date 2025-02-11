package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(autKey, autPass string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() != fiber.MethodGet {
			apiKey := c.Get(autKey)
			if apiKey != autPass {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": ErrUnautorate.Error(), "code": fiber.StatusUnauthorized,
				})
			}
		}
		return c.Next()
	}
}
