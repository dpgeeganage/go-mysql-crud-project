package repository

import (
	"database/sql"
	"go-mysql-crud-project/config"
	"go-mysql-crud-project/model"
)

// UserRepository handles database operations for users
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new repository instance
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser inserts a new user into the database
func CreateUser(user model.User) error {
	_, err := config.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	return err
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(id int) (model.User, error) {
	row := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
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
	return users, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(id int, user model.User) error {
	_, err := config.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	return err
}

// DeleteUser deletes a user by ID from the database
func DeleteUser(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}


