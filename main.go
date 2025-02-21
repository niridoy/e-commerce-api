// main.go
package main

import (
	"echo-go-api/config"
	"echo-go-api/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize database connection
	db := config.ConnectDB()
	defer db.Close() // Ensure the database connection is closed when the app exits

	// Set up the Echo web framework
	e := echo.New()

	// Register all routes
	routes.RegisterRoutes(e, db)

	// Start the server
	log.Println("Server started on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
