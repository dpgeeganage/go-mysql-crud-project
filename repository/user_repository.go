package repository

import (
	"database/sql"
	"errors"
	"go-mysql-crud-project/config"
	"go-mysql-crud-project/model"
)

// UserRepository handles database operations for users
type UserRepository struct {
	DB *sql.DB
}

// IsEmailExists checks if an email already exists in the database
func IsEmailExists(email string) (bool, error) {
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&exists)
	return exists, err
}

// CreateUser inserts a new user into the database
func CreateUser(user model.User) error {
	exists, err := IsEmailExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("Email already exists")
	}

	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(id int) (model.User, error) {
	var user model.User
	err := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]model.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(id int, user model.User) error {
	result, err := config.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		return err
	}
	
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// DeleteUser deletes a user by ID from the database
func DeleteUser(id int) error {
	result, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
