package usecase

import (
	"microservices-travel-backend/internal/hotel-booking/domain"
)

type HotelUsecase struct {
	provider infrastructure.HotelProvider
}

func NewHotelUsecase(provider infrastructure.HotelProvider) *HotelUsecase {
	return &HotelUsecase{provider: provider}
}

func (u *HotelUsecase) SearchHotels(cityCode string) ([]domain.Hotel, error) {
	return u.provider.SearchHotels(cityCode)
}
