package repository

import "time"

type Task struct {
	ID          string    `json:"id" validate:"required,uuid4"`
	Title       string    `json:"title" validate:"required,min=3"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status" validate:"oneof=todo in_progress done"`
	CreatedAt   time.Time `json:"created_at"`
}
