// routes/user.routes.go
package routes

import (
	"database/sql"
	"echo-go-api/handler"
	"echo-go-api/repository"
	"echo-go-api/usecase"

	"github.com/labstack/echo/v4"
)

// RegisterUserRoutes handles user-related routes
func RegisterUserRoutes(e *echo.Echo, db *sql.DB) {
	// Initialize repository, usecase, and handler
	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(repo)
	h := handler.NewUserHandler(uc)

	// Group user routes
	userRoutes := e.Group("/users")
	userRoutes.GET("", h.GetAllUsers)             // Get all users
	userRoutes.POST("", h.CreateUser)             // Create a new user
	userRoutes.GET("/:id", h.GetUser)             // Get user by ID
	userRoutes.PUT("/:id", h.UpdateUser)          // Update user by ID
	userRoutes.PATCH("/:id", h.PartialUpdateUser) // Partially update user
	userRoutes.DELETE("/:id", h.DeleteUser)       // Delete user by ID
}
