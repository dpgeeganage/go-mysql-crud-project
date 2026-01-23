package service

import (
	"go-mysql-crud-project/model"
	"go-mysql-crud-project/repository"
)

// UserService provides user-related services
type UserService struct {
	repo *repository.UserRepository
}

// CreateUser creates a new user
func CreateUser(user model.User) error {
	return repository.CreateUser(user)
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int) (model.User, error) {
	return repository.GetUserByID(id)
}

// GetAllUsers retrieves all users
func GetAllUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}

// UpdateUser updates an existing user
func UpdateUser(id int, user model.User) error {
	return repository.UpdateUser(id, user)
}

// DeleteUser deletes a user by ID
func DeleteUser(id int) error {
	return repository.DeleteUser(id)
}

