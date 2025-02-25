package postgres

import (
	"context"
	"fmt"
	"log"
	"microservices-travel-backend/internal/user-service/domain/user"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresHotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository() (*PostgresHotelRepository, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	sslMode := os.Getenv("DATABASE_SSLMODE")

	var dsn string
	if databaseName == "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/?sslmode=%s",
			databaseUsername, databasePassword, databaseURL, databasePort, sslMode)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			databaseUsername, databasePassword, databaseURL, databasePort, databaseName, sslMode)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed connect to database: %v", err)
	}

	log.Println("Successfully connected to the database")

	return &PostgresHotelRepository{db: db}, nil
}

func (r *PostgresHotelRepository) Create(ctx context.Context, reciept *user.User) (*user.User, error) {
	if err := r.db.WithContext(ctx).Create(reciept).Error; err != nil {
		return nil, err
	}
	return reciept, nil
}