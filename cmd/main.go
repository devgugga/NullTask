package main

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/devgugga/NullTask/internal/config"
	"github.com/devgugga/NullTask/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	// Editing the log settings to make it easily to debug and understanding
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "NullTask 🚫",
	})

	// Loading .env file
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading .env file", "error", err)
	}

	// Getting the environment variables
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	// Setting the database config using the environment value
	dbConfig := &config.DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
	}

	// Starting the database connection using the dbConfig
	db, err := config.ConnectDB(dbConfig)
	if err != nil {
		logger.Fatal("Error connecting to the database", "error", err)
	}

	// Migrating models to the database
	err = db.AutoMigrate(&models.User{}, &models.Tasks{})
	if err != nil {
		logger.Error("Failed to creating tables in the database", "error", err)
	}

}
