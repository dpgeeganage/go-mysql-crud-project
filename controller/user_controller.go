package controller

import (
	"go-mysql-crud-project/service"
	"go-mysql-crud-project/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"errors"
)

// CreateUser handles the creation of a new user
func CreateUser(c *gin.Context) {
	var user model.User

	// validate JSON body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call service
	err := service.CreateUser(user)
	if err != nil {
		if errors.Is(err, service.ErrEmailExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GetUserByID retrieves a user by ID
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetAllUsers retrieves all users
func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user model.User
	// validate JSON body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = service.UpdateUser(id, user)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = service.DeleteUser(id)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}