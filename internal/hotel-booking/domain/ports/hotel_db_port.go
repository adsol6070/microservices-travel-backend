package ports

import "microservices-travel-backend/internal/hotel-booking/domain/models"

type HotelDB interface {
	SaveHotel(hotel *models.Hotel) error
}
