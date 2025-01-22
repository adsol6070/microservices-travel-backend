package services

import (
	"microservices-travel-backend/internal/booking-service/domain/models"
	"microservices-travel-backend/internal/booking-service/domain/ports"
)

type BookingService struct {
	db ports.BookingDB
}

// NewBookingService initializes a new BookingService
func NewBookingService(db ports.BookingDB) *BookingService {
	return &BookingService{db: db}
}

// GetAllBookings retrieves all bookings from the repository
func (b *BookingService) GetAllBookings() ([]models.Booking, error) {
	return b.db.GetAllBookings()
}

// GetBookingByID retrieves a booking by its ID
func (b *BookingService) GetBookingByID(id string) (*models.Booking, error) {
	booking, err := b.db.GetBookingByID(id)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

// CreateBooking creates a new booking
func (b *BookingService) CreateBooking(booking *models.Booking) error {
	if err := b.db.CreateBooking(booking); err != nil {
		return err
	}
	return nil
}

// UpdateBookingStatus updates the status of a booking
func (b *BookingService) UpdateBookingStatus(id string, status string) error {
	if err := b.db.UpdateBookingStatus(id, status); err != nil {
		return err
	}
	return nil
}

// GetBookingsByUserID retrieves bookings for a specific user
func (b *BookingService) GetBookingsByUserID(userID string) ([]models.Booking, error) {
	bookings, err := b.db.GetBookingsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

// DeleteBooking deletes a booking by its ID
func (b *BookingService) DeleteBooking(id string) error {
	if err := b.db.DeleteBooking(id); err != nil {
		return err
	}
	return nil
}

// UpdateBooking updates an existing booking
func (b *BookingService) UpdateBooking(id string, booking *models.Booking) (*models.Booking, error) {
	updatedBooking, err := b.db.UpdateBooking(id, booking)
	if err != nil {
		return nil, err
	}
	return updatedBooking, nil
}
