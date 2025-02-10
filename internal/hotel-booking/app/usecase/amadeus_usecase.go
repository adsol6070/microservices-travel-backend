package usecase

import (
	"microservices-travel-backend/internal/hotel-booking/domain/amadeus"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/models"
)

type HotelUsecase struct {
	service *amadeus.AmadeusService
}

func NewHotelUsecase(service *amadeus.AmadeusService) *HotelUsecase {
	return &HotelUsecase{
		service: service,
	}
}

func (u *HotelUsecase) SearchHotels(cityCode string) ([]models.HotelData, error) {
	hotels, err := u.service.SearchHotels(cityCode)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}

func (u *HotelUsecase) FetchHotelOffers(hotelIDs []string, adults int) ([]models.HotelOffer, error) {
	offers, err := u.service.FetchHotelOffers(hotelIDs, adults)
	if err != nil {
		return nil, err
	}
	return offers, nil
}
