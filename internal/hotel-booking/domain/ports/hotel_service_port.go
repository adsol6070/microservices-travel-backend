package ports

import "microservices-travel-backend/internal/hotel-booking/domain/models"

type HotelService interface {
	GetAllHotels() ([]models.Hotel, error)
	CreateHotel(hotel *models.Hotel) (*models.Hotel, error)
	GetHotelByID(id string) (*models.Hotel, error)
	UpdateHotel(id string, hotel *models.Hotel) (*models.Hotel, error)
	DeleteHotel(id string) error
}
