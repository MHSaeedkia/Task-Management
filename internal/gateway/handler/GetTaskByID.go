package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Fetch a single task by ID
func (handler *Handler) GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": ErrInvalidID.Error(), "code": fiber.ErrBadRequest})
	}

	handler.Mutex.RLock()
	defer handler.Mutex.RUnlock()

	task, exists := handler.Tasks[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": ErrTaskNotFound.Error(), "code": fiber.ErrNotFound})
	}
	return c.JSON(task)
}
