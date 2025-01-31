package repositories

import (
	"errors"
	"fmt"
	"log"
	"microservices-travel-backend/internal/user-service/domain/models"

	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQLUserRepository struct {
	db *gorm.DB
}

func NewPostgreSQLUserRepository() (*PostgreSQLUserRepository, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	sslMode := os.Getenv("DATABASE_SSLMODE")

	// If DATABASE_NAME is not provided, connect without it (to the default 'postgres' database)
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

	return &PostgreSQLUserRepository{db: db}, nil
}

// Create creates a new user in the database using GORM
func (repo *PostgreSQLUserRepository) Create(user models.User) (*models.User, error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	if err := repo.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID retrieves a user by their ID using GORM
func (repo *PostgreSQLUserRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	if err := repo.db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by their email using GORM
func (repo *PostgreSQLUserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetAll retrieves all users using GORM
func (repo *PostgreSQLUserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Update updates a user's information using GORM
func (repo *PostgreSQLUserRepository) Update(id string, user models.User) (*models.User, error) {
	var existingUser models.User
	if err := repo.db.First(&existingUser, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Updating user fields
	if err := repo.db.Model(&existingUser).Updates(user).Error; err != nil {
		return nil, err
	}

	// After successful update, reload the user
	if err := repo.db.First(&existingUser, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}

// Delete removes a user by their ID using GORM
func (repo *PostgreSQLUserRepository) Delete(id string) error {
	if err := repo.db.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
