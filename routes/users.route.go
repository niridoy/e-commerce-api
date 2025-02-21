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
	userRoutes := e.Group("/api/users")
	userRoutes.GET("", h.GetUsers)
	userRoutes.POST("", h.CreateUser)
	userRoutes.GET("/:id", h.GetUser)
	userRoutes.PUT("/:id", h.UpdateUser)
	userRoutes.PATCH("/:id", h.UpdateUser)
	userRoutes.DELETE("/:id", h.DeleteUser)
}
