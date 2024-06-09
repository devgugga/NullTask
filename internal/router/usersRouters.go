package router

import (
	"github.com/charmbracelet/log"
	"github.com/devgugga/NullTask/internal/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SetupUserRoutes sets up the user related routes for the given echo group.
// It requires a pointer to an echo.Group, a pointer to a gorm.DB instance, and a pointer to a log.Logger instance.
// The function initializes a Handler instance with the provided DB and Log instances.
// Then, it registers various HTTP methods and their corresponding handlers for user-related operations.
//
// POST /create: Creates a new user.
// GET /get/email/:email: Retrieves a user by their email.
// GET /get/id/:id: Retrieves a user by their ID.
// GET /get/all/: Retrieves all users.
// PUT /update/email/:email: Updates a user by their email.
// PUT /update/id/:id: Updates a user by their ID.
// DELETE /delete/email/:email: Deletes a user by their email.
func SetupUserRoutes(g *echo.Group, db *gorm.DB, log *log.Logger) {
	h := &handlers.Handler{DB: db, Log: log}

	g.POST("/create", h.CreateUser)
	g.GET("/get/email/:email", h.GetUserByEmail)
	g.GET("/get/id/:id", h.GetUserById)
	g.GET("/get/all/", h.GetAllUsers)
	g.PUT("/update/email/:email", h.UpdateUserByEmail)
	g.PUT("/update/id/:id", h.UpdateUserByEmail)
	g.DELETE("/delete/email/:email", h.DeleteUserByEmail)
}
