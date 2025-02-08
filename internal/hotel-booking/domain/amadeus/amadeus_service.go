package amadeus

import (
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/models"
)

type AmadeusService struct {
	client hotels.AmadeusClient
}

func NewAmadeusService(client hotels.AmadeusClient) *AmadeusService {
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

func (a *AmadeusService) SearchHotels(cityCode string) ([]models.Hotel, error) {
	hotels, err := a.client.SearchHotels(cityCode)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
