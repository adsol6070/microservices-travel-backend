package infrastructure

import "microservices-travel-backend/internal/hotel-booking/domain"

type HotelProvider interface {
	SearchHotels(cityCode string) ([]domain.Hotel, error)
}