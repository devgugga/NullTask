package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConfig stores the configuration parameters for connecting to a PostgreSQL database.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// ConnectDB establishes a connection to a PostgreSQL database using the provided configuration.
// It returns the GORM database instance and an error if the connection fails.
//
// Parameters:
// conf: A pointer to a DBConfig struct containing the database connection configuration parameters.
//
// Returns:
// db: A pointer to a gorm.DB instance representing the established database connection.
// err: An error object that will be nil if the connection is successful, otherwise it will contain the error details.
//
// Note: The function creates a connection string using the provided configuration parameters and opens the connection using gorm.Open.
// It also executes a SQL command to create the "uuid-ossp" extension if it does not already exist in the database.
func ConnectDB(conf *DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)

	// Open the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	return db, nil
}
