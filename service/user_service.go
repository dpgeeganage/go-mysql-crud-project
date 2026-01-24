package service

import (
	"database/sql"
	"go-mysql-crud-project/model"
	"go-mysql-crud-project/repository"
)

// UserService provides user-related services
type UserService struct {
	repo *repository.UserRepository
}

// CreateUser creates a new user
func CreateUser(user model.User) error {
	err := repository.CreateUser(user)
	if err != nil {
		if err.Error() == "Email already exists" {
			return ErrEmailExists
		}
		return err
	}
	return nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int) (model.User, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}

// GetAllUsers retrieves all users
func GetAllUsers() ([]model.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates an existing user
func UpdateUser(id int, user model.User) error {
	err := repository.UpdateUser(id, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id int) error {
	err := repository.DeleteUser(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}

