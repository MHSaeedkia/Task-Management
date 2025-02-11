package handler

import "errors"

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidID    = errors.New("invalid ID format")
	ErrInvalidInput = errors.New("invalid input")
	ErrTimeOut      = errors.New("request timed out")
)
