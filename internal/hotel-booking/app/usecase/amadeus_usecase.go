package usecase

import (
	"microservices-travel-backend/internal/hotel-booking/domain"
)

type HotelUsecase struct {}

func NewHotelUsecase( /* provider infrastructure.HotelProvider */ ) *HotelUsecase {
	return &HotelUsecase{}
}

func (u *HotelUsecase) SearchHotels(cityCode string) ([]domain.Hotel, error) {
	// return u.provider.SearchHotels(cityCode)
}

func (u *HotelUsecase) FetchHotelOffers(hotelIDs []string, adults int) ([]models.HotelOffer, error) {

}
