package main

import (
	"os"

	"github.com/devgugga/NullTask/internal/config"
	"github.com/devgugga/NullTask/internal/models"
	"github.com/joho/godotenv"
)

// main function is the entry point of the application.
func main() {
	// Initializing logger using the config package's Logger function.
	logger := config.Logger()

	// Loading.env file using godotenv's Load function.
	// If.env file is not found, it will return an error.
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading.env file", "error", err)
	}

	// Getting the environment variables HOST, PORT, USER, PASSWORD, DBNAME using os.Getenv function.
	host := os.Getenv("HOST")
	dbPort := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	// Creating a DBConfig struct with the obtained environment variables.
	dbConfig := &config.DBConfig{
		Host:     host,
		Port:     dbPort,
		User:     user,
		Password: password,
		DBName:   dbname,
	}

	// Connecting to the database using the config package's ConnectDB function.
	// If there is an error during the connection, it will return an error.
	db, err := config.ConnectDB(dbConfig)
	if err != nil {
		logger.Fatal("Error connecting to the database", "error", err)
	}

	// Migrating the User and Tasks models to the database using GORM's AutoMigrate function.
	// If there is an error during the migration, it will return an error.
	err = db.AutoMigrate(&models.User{}, &models.Tasks{})
	if err != nil {
		logger.Error("Failed to creating tables in the database", "error", err)
	}

	// Starting the server on port "1323" using the config package's StartServer function.
	// It takes the port number, database connection, and logger as parameters.
	config.StartServer("1323", db, logger)
}
