package usecase

import (
	"microservices-travel-backend/internal/hotel-booking/app/dto/request"
	"microservices-travel-backend/internal/hotel-booking/app/dto/response"
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

func (u *HotelUsecase) SearchHotels(req request.HotelSearchRequest) ([]models.EnrichedHotelOffer, error) {
	hotelsWithOffer, err := u.service.SearchHotels(req)
	if err != nil {
		return nil, err
	}
	return hotelsWithOffer, nil
}

func (u *HotelUsecase) HotelDetails(req request.HotelDetailsRequest) (response.HotelDetails, error) {
	hotelDetails, err := u.service.HotelDetails(req)
	if err != nil {
		return response.HotelDetails{}, err
	}
	return hotelDetails, nil
}

func (u *HotelUsecase) FetchHotelOffers(hotelIDs []string, adults int) ([]amadeusHotelModels.HotelOffer, error) {
	offers, err := u.service.FetchHotelOffers(hotelIDs, adults)
	if err != nil {
		return nil, err
	}
	return offers, nil
}

func (u *HotelUsecase) CreateHotelBooking(requestBody amadeusHotelModels.HotelBookingReq) (*amadeusHotelModels.HotelOrderResponseData, error) {
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
