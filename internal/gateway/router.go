package gateway

import (
	"Task-Management/config"
	"Task-Management/internal/gateway/handler"
	"Task-Management/internal/gateway/middleware"
	"Task-Management/internal/logger"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run(config *config.Config) {
	handler := handler.NewHandler()

	app := fiber.New()

	// Start logging goroutine
	handler.Wg.Add(1)
	go func() {
		defer handler.Wg.Done()
		logger.LogTasks(handler.TaskChan)
	}()

	// Middleware
	app.Use(middleware.LoggingMiddleware)
	app.Use("/tasks", middleware.RateLimiter(config.MaxRequest, config.ExpirationTime, config.IgnoreIp))
	app.Use("/tasks", middleware.AuthMiddleware(config.AutKey, config.AutPass))

	// Routes
	app.Get("/tasks", handler.GetTasks)
	app.Get("/tasks/:id", handler.GetTaskByID)
	app.Post("/tasks", handler.CreateTask)
	app.Put("/tasks/:id", handler.UpdateTask)
	app.Delete("/tasks/:id", handler.DeleteTask)

	// Graceful shutdown
	defer func() {
		close(handler.TaskChan)
		handler.Wg.Wait()
	}()

	log.Fatal(app.Listen(fmt.Sprintf("%s:%v", config.ServiceIp, config.ServicePort)))
}
