package handler

import (
	"Task-Management/internal/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (handler *Handler) CreateTask(c *fiber.Ctx) error {
	var task repository.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": ErrInvalidInput.Error(), "code": fiber.ErrBadRequest})
	}

	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()

	if err := handler.Validate.Struct(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "code": fiber.ErrBadRequest})
	}

	handler.Mutex.Lock()
	handler.Tasks[task.ID] = task
	handler.Mutex.Unlock()

	handler.TaskChan <- task
	return c.Status(fiber.StatusCreated).JSON(task)
}
