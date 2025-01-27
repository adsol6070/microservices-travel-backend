package ports

import "microservices-travel-backend/internal/hotel-booking/domain/models"

type HotelService interface {
	FetchHotels() ([]models.Hotel, error)
}
