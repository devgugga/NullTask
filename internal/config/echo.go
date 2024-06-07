package config

import (
	"github.com/charmbracelet/log"
	"github.com/devgugga/NullTask/internal/router"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func StartServer(port string, db *gorm.DB, log *log.Logger) {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &CustomValidator{validator: validator.New()}

	router.SetupRoutes(e, db, log)

	err := e.Start(":" + port)
	if err != nil {
		log.Fatal("Failed to start the api server", "error", err)
	}
}
