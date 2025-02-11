package handler

import (
	"Task-Management/internal/repository"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func setupApp() (*Handler, *fiber.App) {
	handler := NewHandler()
	app := fiber.New()

	app.Post("/tasks", handler.CreateTask)
	app.Delete("/tasks/:id", handler.DeleteTask)
	app.Get("/tasks/:id", handler.GetTaskByID)
	app.Get("/tasks", handler.GetTasks)
	app.Put("/tasks/:id", handler.UpdateTask)

	return &handler, app
}

func TestCreateTask(t *testing.T) {

}

func TestDeleteTask(t *testing.T) {
	handler, app := setupApp()

	// Add a task
	task := repository.Task{
		ID:          uuid.New().String(),
		Title:       "Task to Delete",
		Description: "Description",
		Status:      "todo",
		CreatedAt:   time.Now(),
	}
	handler.Tasks[task.ID] = task

	t.Run("Delete existing task", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/tasks/"+task.ID, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)

		// Verify task is removed from handler state
		handler.Mutex.RLock()
		defer handler.Mutex.RUnlock()
		_, exists := handler.Tasks[task.ID]
		assert.False(t, exists)
	})

	t.Run("Delete non-existent task", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/tasks/invalid-id", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}

func TestGetTaskByID(t *testing.T) {
	handler, app := setupApp()

	// Add a task
	task := repository.Task{
		ID:          uuid.New().String(),
		Title:       "Get Task",
		Description: "Task description",
		Status:      "todo",
		CreatedAt:   time.Now(),
	}
	handler.Tasks[task.ID] = task

	t.Run("Get existing task", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/tasks/"+task.ID, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var fetchedTask repository.Task
		err := json.NewDecoder(resp.Body).Decode(&fetchedTask)
		assert.NoError(t, err)
		assert.Equal(t, task.ID, fetchedTask.ID)
	})

	t.Run("Get non-existent task", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/tasks/invalid-id", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}

func TestGetTasks(t *testing.T) {

}

func TestUpdateTask(t *testing.T) {

}
