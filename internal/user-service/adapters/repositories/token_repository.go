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

type PostgreSQLTokenRepository struct {
	db *gorm.DB
}

func NewPostgreSQLTokenRepository() (*PostgreSQLTokenRepository, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	sslMode := os.Getenv("DATABASE_SSLMODE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		databaseUsername, databasePassword, databaseURL, databasePort, databaseName, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database")
	db.AutoMigrate(&models.Token{}) // AutoMigrate Token Model

	return &PostgreSQLTokenRepository{db: db}, nil
}

func (repo *PostgreSQLTokenRepository) StoreToken(token models.Token) (*models.Token, error) {
	if token.ID == "" {
		token.ID = uuid.New().String()
	}
	if err := repo.db.Create(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (repo *PostgreSQLTokenRepository) GetToken(userID string) (*models.Token, error) {
	var token models.Token
	if err := repo.db.First(&token, "user_id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("token not found")
		}
		return nil, err
	}
	return &token, nil
}

func (repo *PostgreSQLTokenRepository) DeleteToken(userID string) error {
	if err := repo.db.Delete(&models.Token{}, "user_id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}
