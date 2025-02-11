package handler

import (
	"Task-Management/internal/repository"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) GetTasks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		return c.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{"error": ErrTimeOut.Error(), "code": fiber.StatusRequestTimeout})
	case <-time.After(1 * time.Second):
		handler.Mutex.RLock()
		defer handler.Mutex.RUnlock()

		var result []repository.Task
		for _, task := range handler.Tasks {
			result = append(result, task)

		}
		return c.JSON(result)
	}
}
