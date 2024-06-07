package router

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, log *log.Logger) {
	userRoutes := e.Group("/users")
	SetupUserRoutes(userRoutes, db, log)
}
