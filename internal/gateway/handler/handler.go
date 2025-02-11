package handler

import (
	"Task-Management/internal/repository"
	"sync"

	"github.com/go-playground/validator"
)

type Handler struct {
	Tasks    map[string]repository.Task
	Mutex    sync.RWMutex
	Validate *validator.Validate
	TaskChan chan repository.Task
	Wg       sync.WaitGroup
}

func NewHandler() Handler {
	return Handler{
		Tasks:    make(map[string]repository.Task),
		Mutex:    sync.RWMutex{},
		Validate: validator.New(),
		TaskChan: make(chan repository.Task),
		Wg:       sync.WaitGroup{},
	}
}
