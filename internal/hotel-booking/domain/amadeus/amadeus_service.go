package amadeus

import (
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/models"
)

type AmadeusService struct {
	client *hotels.AmadeusClient
}

func NewAmadeusService(client *hotels.AmadeusClient) *AmadeusService {
	return &AmadeusService{
		client: client,
	}
}

func (a *AmadeusService) FetchHotelOffers(hotelIDs []string, adults int) ([]models.HotelOffer, error) {
	offers, err := a.client.FetchHotelOffers(hotelIDs, adults)
	if err != nil {
		return nil, err
	}
	return offers, nil
}

func (a *AmadeusService) SearchHotels(cityCode string) ([]models.HotelData, error) {
	hotels, err := a.client.HotelSearch(cityCode)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}

func (a *AmadeusService) CreateHotelBooking(requestBody models.HotelBookingRequest) (*models.HotelOrderResponse, error) {
	booking, err := a.client.CreateHotelBooking(requestBody)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (a *AmadeusService) FetchHotelRatings(hotelIDs []string) (*models.HotelSentimentResponse, error) {
	ratings, err := a.client.FetchHotelRatings(hotelIDs)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (a *AmadeusService) HotelNameAutoComplete(keyword string, subtype string) (*models.HotelNameResponse, error) {
	hotels, err := a.client.HotelNameAutoComplete(keyword, subtype)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
