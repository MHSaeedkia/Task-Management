package logger

import (
	"Task-Management/internal/repository"
	"log"
)

func LogTasks(taskChan <-chan repository.Task) {
	for task := range taskChan {
		log.Printf("Task updated: ID=%s, Title=%s, Status=%s", task.ID, task.Title, task.Status)
	}
}
