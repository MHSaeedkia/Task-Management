package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter(maxReq int, expirTime time.Duration, ignorIp string) fiber.Handler {
	return limiter.New(limiter.Config{
		// Ignore rate limiting for requests from IP "127.0.0.1"
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == ignorIp
		},

		// Maximum of 10 requests per 30 seconds per IP
		Max:        maxReq,
		Expiration: expirTime * time.Second,

		// Use IP address as the key for rate limiting
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},

		// Custom response when limit is reached
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": ErrTooRequest.Error(), "code": fiber.StatusTooManyRequests,
			})
		},
	})
}
