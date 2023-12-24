package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConnect establishes a connection to a PostgreSQL database
func DBConnect(host string, port string, username string, password string, dbname string, models ...interface{}) (*gorm.DB, error) {
	// create the DSN string for connecting to the database
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, username, dbname, password)

	// open a connection to the database using the DSN string
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// check if any models were passed in, and if so, auto-migrate them
	if len(models) > 0 {
		err = db.AutoMigrate(models...)
		if err != nil {
			return nil, fmt.Errorf("failed to perform auto migration: %w", err)
		}
	}

	return db, nil
}
