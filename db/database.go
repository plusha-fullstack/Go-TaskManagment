package db

import (
	"fmt"
	"task-tracker/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the PostgreSQL database
func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=user password=pass dbname=gorm port=5432 sslmode=disable"

	// Retry logic for database connection
	var db *gorm.DB
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// AutoMigrate models
	db.AutoMigrate(&models.User{}, &models.Task{})

	return db, nil
}
