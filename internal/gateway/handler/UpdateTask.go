package handler

import (
	"Task-Management/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (handler *Handler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": ErrInvalidID.Error(), "code": fiber.ErrBadRequest})
	}

	var updates repository.Task
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": ErrInvalidInput, "code": fiber.ErrBadRequest})
	}

	handler.Mutex.Lock()
	defer handler.Mutex.Unlock()

	task, exists := handler.Tasks[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": ErrTaskNotFound.Error(), "code": fiber.ErrNotFound})
	}

	if updates.Title != "" {
		task.Title = updates.Title
	}
	if updates.Description != "" {
		task.Description = updates.Description
	}
	if updates.Status != "" {
		task.Status = updates.Status
	}

	handler.Tasks[id] = task
	handler.TaskChan <- task
	return c.JSON(task)
}
