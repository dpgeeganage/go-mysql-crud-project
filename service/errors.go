package service

import "errors"

var (
	ErrUserNotFound = errors.New("User not found")
	ErrEmailExists = errors.New("Email already exists")
)