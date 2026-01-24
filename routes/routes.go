package routes

import (
	"go-mysql-crud-project/controller"
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API routes
func SetupRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", controller.CreateUser)
		userRoutes.GET("/:id", controller.GetUserByID)
		userRoutes.GET("/", controller.GetAllUsers)
		userRoutes.PUT("/:id", controller.UpdateUser)
		userRoutes.DELETE("/:id", controller.DeleteUser)
	}
}