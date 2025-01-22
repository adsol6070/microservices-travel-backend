package ports

import "microservices-travel-backend/internal/booking-service/domain/models"

type BookingService interface {
	GetAllBookings() ([]models.Booking, error)
	GetBookingByID(id string) (*models.Booking, error)
	CreateBooking(booking *models.Booking) error
	UpdateBookingStatus(id string, status string) error
	GetBookingsByUserID(userID string) ([]models.Booking, error)
	DeleteBooking(id string) error
	UpdateBooking(id string, booking *models.Booking) (*models.Booking, error)
}
