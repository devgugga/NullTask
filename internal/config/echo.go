package config

import (
	"github.com/charmbracelet/log"
	"github.com/devgugga/NullTask/internal/router"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

// CustomValidator is a custom implementation of echo.Validator interface.
// It uses validator.v10 package for struct validation.
type CustomValidator struct {
	validator *validator.Validate
}

// Validate method validates the given struct using the validator.v10 package.
// It returns an error if the struct fails validation, otherwise it returns nil.
//
// Parameters:
// - i: The struct to be validated. It can be any struct type.
//
// Return values:
// - error: An error if the struct fails validation, otherwise nil.
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// StartServer initializes and starts the API server.
// It sets up Echo framework, registers middlewares, sets up custom validator,
// and calls SetupRoutes function to register routes.
//
// Parameters:
// - port: The port number on which the server should listen.
// - db: A pointer to the gorm.DB instance for database operations.
// - log: A pointer to the log.Logger instance for logging.
//
// Return values:
// - This function does not return any value.
//
// Errors:
// - If there is an error starting the server, it will be logged and the program will exit.
func StartServer(port string, db *gorm.DB, log *log.Logger) {

	e := echo.New()

	// Use Echo's built-in Logger middleware
	e.Use(middleware.Logger())

	// Use Echo's built-in Recover middleware to recover from panics
	e.Use(middleware.Recover())

	// Set up a custom validator using validator.v10 package
	e.Validator = &CustomValidator{validator: validator.New()}

	// Register routes using the SetupRoutes function from the router package
	router.SetupRoutes(e, db, log)

	// Start the server on the specified port
	err := e.Start(":" + port)
	if err != nil {
		// Log the error and exit the program if there is an error starting the server
		log.Fatal("Failed to start the api server", "error", err)
	}
}
