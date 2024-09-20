package utils

import "errors"

var (
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidToken          = errors.New("invalid token")
	ErrTokenExpired          = errors.New("token expired")
	ErrAccountNotFound       = errors.New("account not found")
	ErrAccountCreationFailed = errors.New("account creation failed")
	// Add more custom errors as needed
)
