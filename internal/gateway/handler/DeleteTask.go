package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (handler *Handler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": ErrInvalidID.Error(), "code": fiber.ErrBadRequest})
	}

	handler.Mutex.Lock()
	defer handler.Mutex.Unlock()

	_, exists := handler.Tasks[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": ErrTaskNotFound.Error(), "code": fiber.ErrNotFound})
	}

	delete(handler.Tasks, id)
	return c.Status(fiber.StatusNoContent).Send(nil)
}
