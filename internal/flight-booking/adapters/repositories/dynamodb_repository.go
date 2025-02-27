package repositories

import (
	"fmt"
	"log"
	"microservices-travel-backend/internal/flight-booking/domain/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresBookingRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository() (*PostgresBookingRepository, error) {
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
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database")

	return &PostgresBookingRepository{DB: db}, nil
}

func (r *PostgresBookingRepository) CreateFlight(flight *models.Flight) (*models.Flight, error) {
	fmt.Println("Flight created in DynamoDB")
	return flight, nil
}

func (r *PostgresBookingRepository) GetFlightByID(id string) (*models.Flight, error) {
	return &models.Flight{}, nil
}

func (r *PostgresBookingRepository) UpdateFlight(id string, flight *models.Flight) (*models.Flight, error) {
	return flight, nil
}

func (r *PostgresBookingRepository) DeleteFlight(id string) error {
	return nil
}
