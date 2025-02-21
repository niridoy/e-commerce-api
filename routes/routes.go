// routes/routes.go
package routes

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes calls all route registration functions
func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	RegisterUserRoutes(e, db) // Import user routes
}
