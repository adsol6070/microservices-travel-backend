package usecase

import (
	"microservices-travel-backend/internal/hotel-booking/domain/amadeus"
	"microservices-travel-backend/internal/hotel-booking/domain/models"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/amadeusHotelModels"
)

type HotelUsecase struct {
	service *amadeus.AmadeusService
}

func NewHotelUsecase(service *amadeus.AmadeusService) *HotelUsecase {
	return &HotelUsecase{
		service: service,
	}
}

type SearchHotelsRequest struct {
	CityCode     string `json:"cityCode"`
	CheckInDate  string `json:"checkInDate"`
	CheckOutDate string `json:"checkOutDate"`
	Rooms        int    `json:"rooms"`
	Persons      int    `json:"persons"`
}

func (u *HotelUsecase) SearchHotels(req SearchHotelsRequest) ([]models.EnrichedHotelOffer, error) {
	amadeusReq := amadeus.SearchHotelsRequest{
		CityCode:     req.CityCode,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		Rooms:        req.Rooms,
		Persons:      req.Persons,
	}

	hotelsWithOffer, err := u.service.SearchHotels(amadeusReq)
	if err != nil {
		return nil, err
	}
	return hotelsWithOffer, nil
}

func (u *HotelUsecase) FetchHotelOffers(hotelIDs []string, adults int) ([]amadeusHotelModels.HotelOffer, error) {
	offers, err := u.service.FetchHotelOffers(hotelIDs, adults)
	if err != nil {
		return nil, err
	}
	return offers, nil
}

func (u *HotelUsecase) CreateHotelBooking(requestBody amadeusHotelModels.HotelBookingRequest) (*amadeusHotelModels.HotelOrderResponse, error) {
	booking, err := u.service.CreateHotelBooking(requestBody)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (u *HotelUsecase) FetchHotelRatings(hotelIDs []string) (*amadeusHotelModels.HotelSentimentResponse, error) {
	ratings, err := u.service.FetchHotelRatings(hotelIDs)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (u *HotelUsecase) HotelNameAutoComplete(keyword string, subtype string) (*amadeusHotelModels.HotelNameResponse, error) {
	hotels, err := u.service.HotelNameAutoComplete(keyword, subtype)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
