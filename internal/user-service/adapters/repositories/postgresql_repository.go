package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"microservices-travel-backend/internal/user-service/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	DB *gorm.DB
}

// NewPostgresUserRepository initializes a new PostgresUserRepository with a DB connection.
func NewPostgresUserRepository() (*PostgresUserRepository, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databasePort := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/?sslmode=require",
		databaseUsername, databasePassword, databaseURL, databasePort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database")

	return &PostgresUserRepository{DB: db}, nil
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	if err := r.DB.Create(user).Error; err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	return nil
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	return &user, nil
}

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	if err := r.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

func (r *PostgresUserRepository) DeleteUser(ctx context.Context, id string) error {
	if err := r.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}

func (r *PostgresUserRepository) AuthenticateUser(ctx context.Context, email, password string) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, "email = ? AND password = ?", email, password).Error; err != nil {
		return nil, fmt.Errorf("authentication failed: %v", err)
	}
	return &user, nil
}
