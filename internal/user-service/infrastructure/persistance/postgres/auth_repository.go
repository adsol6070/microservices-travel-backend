package postgres

import (
	"context"
	"fmt"
	"log"
	"microservices-travel-backend/internal/user-service/domain/user"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository() (*PostgresAuthRepository, error) {
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

	return &PostgresAuthRepository{db: db}, nil
}

func (r *PostgresAuthRepository) UpdatePassword(ctx context.Context, userID string, newPasswordHash string) error {
	var user user.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
		return fmt.Errorf("user not found with ID: %s, error: %w", userID, err)
	}

	user.Password = newPasswordHash
	user.UpdatedAt = time.Now()

	if err := r.db.WithContext(ctx).Save(&user).Error; err != nil {
		return fmt.Errorf("failed to update password for user %s: %w", userID, err)
	}

	return nil
}
