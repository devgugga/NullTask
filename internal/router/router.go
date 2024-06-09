package router

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SetupRoutes sets up the routes for the application.
// It takes an Echo instance, a GORM database connection, and a Charmbracelet log instance as parameters.
// The function creates a group for user routes and sets up the routes using the SetupUserRoutes function.
func SetupRoutes(e *echo.Echo, db *gorm.DB, log *log.Logger) {
	// Create a group for user routes
	userRoutes := e.Group("/users")

	// Set up the routes for user group using the SetupUserRoutes function
	SetupUserRoutes(userRoutes, db, log)
}
