package routes

import (
	"go-mysql-crud-project/controller"
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API routes
func SetupRoutes(router *gin.Engine) {
	router.POST("/users", controller.CreateUser)
	router.GET("/users/:id", controller.GetUserByID)
	router.GET("/users", controller.GetAllUsers)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
}