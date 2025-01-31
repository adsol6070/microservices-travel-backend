package repositories

import (
	"fmt"
	"log"
	"microservices-travel-backend/internal/hotel-booking/domain/models"
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

func (r *PostgresBookingRepository) GetAllHotels() ([]models.Hotel, error) {
	var hotels []models.Hotel
	if err := r.DB.Find(&hotels).Error; err != nil {
		return nil, fmt.Errorf("error fetching hotels: %v", err)
	}
	return hotels, nil
}

func (r *PostgresBookingRepository) SaveHotel(hotel *models.Hotel) error {
	if err := r.DB.Create(hotel).Error; err != nil {
		return fmt.Errorf("error creating hotel: %v", err)
	}
	return nil
}

func (r *PostgresBookingRepository) GetHotelByID(id string) (*models.Hotel, error) {
	var hotel models.Hotel
	if err := r.DB.First(&hotel, "hotel_id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("hotel not found: %v", err)
	}
	return &hotel, nil
}

func (r *PostgresBookingRepository) UpdateHotel(id string, hotel *models.Hotel) (*models.Hotel, error) {
	if err := r.DB.Model(&models.Hotel{}).Where("hotel_id = ?", id).Updates(hotel).Error; err != nil {
		return nil, fmt.Errorf("error updating hotel: %v", err)
	}
	return hotel, nil
}

func (r *PostgresBookingRepository) DeleteHotel(id string) error {
	if err := r.DB.Delete(&models.Hotel{}, "hotel_id = ?", id).Error; err != nil {
		return fmt.Errorf("error deleting hotel: %v", err)
	}
	return nil
}
