// config/config.go
package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables from the .env file.
func LoadConfig() {
	err := godotenv.Load() // Load the .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
