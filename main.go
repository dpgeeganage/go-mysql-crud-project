package main

import (
	"go-mysql-crud-project/config"
	"go-mysql-crud-project/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// Connect to the database
	config.ConnectDB()
	defer config.DB.Close()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":" + os.Getenv("SERVER_PORT"))
}

// http://localhost:8080/users