package config

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:root123@tcp(127.0.0.1:3306)/go_crud"

	// Open the database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}

	DB = db
	fmt.Println("Database connected successfully!")
}

