package repositories

import (
	"fmt"
	"microservices-travel-backend/internal/booking-service/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresBookingRepository struct {
	DB *gorm.DB
}

func NewPostgresBookingRepository(dsn string) (*PostgresBookingRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Booking{})
	return &PostgresBookingRepository{DB: db}, nil
}

func (r *PostgresBookingRepository) GetAllBookings() ([]models.Booking, error) {
	var bookings []models.Booking
	if err := r.DB.Find(&bookings).Error; err != nil {
		return nil, fmt.Errorf("error fetching bookings: %v", err)
	}
	return bookings, nil
}

func (r *PostgresBookingRepository) GetBookingByID(id string) (*models.Booking, error) {
	var booking models.Booking
	if err := r.DB.First(&booking, "booking_id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("booking not found: %v", err)
	}
	return &booking, nil
}

func (r *PostgresBookingRepository) CreateBooking(booking *models.Booking) error {
	if err := r.DB.Create(booking).Error; err != nil {
		return fmt.Errorf("error creating booking: %v", err)
	}
	return nil
}

func (r *PostgresBookingRepository) UpdateBookingStatus(id string, status string) error {
	if err := r.DB.Model(&models.Booking{}).Where("booking_id = ?", id).Update("status", status).Error; err != nil {
		return fmt.Errorf("error updating booking status: %v", err)
	}
	return nil
}

func (r *PostgresBookingRepository) GetBookingsByUserID(userID string) ([]models.Booking, error) {
	var bookings []models.Booking
	if err := r.DB.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, fmt.Errorf("error fetching bookings for user: %v", err)
	}
	return bookings, nil
}

func (r *PostgresBookingRepository) DeleteBooking(id string) error {
	if err := r.DB.Delete(&models.Booking{}, "booking_id = ?", id).Error; err != nil {
		return fmt.Errorf("error deleting booking: %v", err)
	}
	return nil
}

func (r *PostgresBookingRepository) UpdateBooking(id string, booking *models.Booking) (*models.Booking, error) {
	if err := r.DB.Model(&models.Booking{}).Where("booking_id = ?", id).Updates(booking).Error; err != nil {
		return nil, fmt.Errorf("error updating booking: %v", err)
	}
	return booking, nil
}