package middleware

import "errors"

var (
	ErrTooRequest = errors.New("Too many requests, please try again later")
	ErrUnautorate = errors.New("Unauthorized")
)
