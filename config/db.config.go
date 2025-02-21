// config/db.config.go
package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB initializes the database connection.
func ConnectDB() *sql.DB {
	// Ensure environment variables are loaded
	LoadConfig()

	// Build the database connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Test the database connection
	if err = db.Ping(); err != nil {
		log.Fatal("Database is unreachable:", err)
	}

	log.Println("Connected to the database successfully!")
	return db
}
