package gateway

import (
	"Task-Management/internal/gateway/handler"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run(serviceIp, servicePort string) {
	// handler1 := handler.NewHandler()

	app := fiber.New()

	// Start logging goroutine

	// handler.Wg.Add(1)
	// go func(handler handler.Handler) {
	// 	defer wg.Done()
	// 	logTasks(taskChan)
	// }(handler)

	// Middleware
	app.Use(handler.LoggingMiddleware)
	app.Use("/tasks", handler.AuthMiddleware)

	// Routes
	app.Get("/tasks", handler.GetTasks)
	app.Get("/tasks/:id", handler.GetTaskByID)
	app.Post("/tasks", handler.CreateTask)
	app.Put("/tasks/:id", handler.UpdateTask)
	app.Delete("/tasks/:id", handler.DeleteTask)

	// Graceful shutdown
	defer func() {
		// close(taskChan)
		// handler.Wg.Wait()
	}()

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", serviceIp, servicePort)))
}
