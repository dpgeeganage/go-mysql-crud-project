package main

import (
	"go-mysql-crud-project/config"
	"go-mysql-crud-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDB()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}

// http://localhost:8080/users