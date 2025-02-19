package amadeus

import (
	"errors"
	"log"
	"microservices-travel-backend/internal/hotel-booking/app/dto/request"
	"microservices-travel-backend/internal/hotel-booking/app/dto/response"
	"microservices-travel-backend/internal/hotel-booking/domain/models"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/amadeusHotelModels"
	"microservices-travel-backend/internal/shared/api_provider/google/places"
	"microservices-travel-backend/internal/shared/api_provider/google/places/googlePlaceModels"
	"sync"
)

type AmadeusService struct {
	client *hotels.AmadeusClient
	places *places.PlacesClient
}

func NewAmadeusService(client *hotels.AmadeusClient, places *places.PlacesClient) *AmadeusService {
	return &AmadeusService{
		client: client,
		places: places,
	}
}

func (a *AmadeusService) FetchHotelOffers(hotelIDs []string, adults int) ([]amadeusHotelModels.HotelOffer, error) {
	offers, err := a.client.FetchHotelOffers(hotelIDs, adults)
	if err != nil {
		return nil, err
	}
	return offers, nil
}

func (a *AmadeusService) SearchHotels(req request.HotelSearchRequest) ([]models.EnrichedHotelOffer, error) {
	hotels, err := a.client.HotelSearch(req.CityCode)
	if err != nil {
		log.Println("ERROR: Failed to fetch hotels -", err)
		return nil, err
	}

	if len(hotels) == 0 {
		log.Println("WARN: No hotels found for the given city code")
		return nil, errors.New("no hotels found for the given city code")
	}

	// Extract hotel IDs
	hotelIDs := extractHotelIDs(hotels)
	log.Printf("INFO: Extracted %d hotel IDs", len(hotelIDs))

	if len(hotelIDs) == 0 {
		log.Println("WARN: No valid hotels found")
		return nil, errors.New("no valid hotels found")
	}

	// Fetch hotel offers using extracted hotel IDs
	hotelOffers, err := a.client.FetchHotelOffers(hotelIDs, req.Persons)
	if err != nil {
		log.Println("ERROR: Failed to fetch hotel offers -", err)
		return nil, err
	}

	if len(hotelOffers) == 0 {
		log.Println("WARN: No hotel offers available")
		return nil, errors.New("no hotel offers available")
	}

	// Filter hotel offers based on search criteria
	filteredHotels := filterHotelsByRequest(hotelOffers, req)

	if len(filteredHotels) == 0 {
		log.Println("WARN: No hotels matched the search criteria")
		return nil, errors.New("no matching hotels found")
	}

	var enrichedHotels []models.EnrichedHotelOffer

	for _, hotel := range filteredHotels {

		textQuery := googlePlaceModels.TextQueryRequest{
			TextQuery: hotel.Hotel.Name,
			LocationBias: googlePlaceModels.LocationBias{
				Circle: googlePlaceModels.Circle{
					Center: googlePlaceModels.Coordinates{
						Latitude:  hotel.Hotel.Latitude,
						Longitude: hotel.Hotel.Longitude,
					},
					Radius: 5000,
				},
			},
		}

		placeDetails, err := a.places.SearchPlaces(
			textQuery,
			"places.id,places.displayName,places.rating,places.photos",
		)

		if err != nil || len(placeDetails.Places) == 0 {
			log.Printf("WARN: No Google Places data found for %s", hotel.Hotel.Name)
			continue
		}

		// Extract first place details
		place := placeDetails.Places[0]
		var photoURL string

		// Fetch photo if available
		if len(place.Photos) > 0 {
			photoResp, err := a.places.GetPlacePhoto(place.Photos[0].Name, 400, 400)
			if err == nil {
				photoURL = photoResp.PhotoURI
			}
		}

		enrichedHotel := models.EnrichedHotelOffer{
			HotelID:           hotel.Hotel.HotelID,
			Name:              hotel.Hotel.Name,
			Latitude:          hotel.Hotel.Latitude,
			Longitude:         hotel.Hotel.Longitude,
			PhotoURL:          photoURL,
			Offers:            hotel.Offers,
			GooglePlaceID:     place.ID,
			Rating:            place.Rating,
			GoogleDisplayName: place.DisplayName.Text,
		}

		enrichedHotels = append(enrichedHotels, enrichedHotel)
	}

	if len(enrichedHotels) == 0 {
		log.Println("WARN: No enriched hotels found")
		return nil, errors.New("no enriched hotels found")
	}

	return enrichedHotels, nil
}

func extractHotelIDs(hotels []amadeusHotelModels.HotelData) []string {
	var hotelIDs []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, hotel := range hotels {
		wg.Add(1)
		go func(h amadeusHotelModels.HotelData) {
			defer wg.Done()
			if h.HotelID != "" {
				mu.Lock()
				hotelIDs = append(hotelIDs, h.HotelID)
				mu.Unlock()
			}
		}(hotel)
	}

	wg.Wait()
	return hotelIDs
}

func filterHotelsByRequest(hotelOffers []amadeusHotelModels.HotelOffer, req request.HotelSearchRequest) []amadeusHotelModels.HotelOffer {
	var filteredHotels []amadeusHotelModels.HotelOffer

	for _, hotelOffer := range hotelOffers {
		if !hotelOffer.Available {
			continue
		}

		for _, offer := range hotelOffer.Offers {
			if offer.CheckInDate == req.CheckInDate && offer.CheckOutDate == req.CheckOutDate {
				if req.Rooms > 0 && req.Rooms == len(hotelOffer.Offers) {
					filteredHotels = append(filteredHotels, hotelOffer)
					break
				}
			}
		}
	}

	return filteredHotels
}

func (a *AmadeusService) HotelDetails(req request.HotelDetailsRequest) ([]response.HotelDetails, error) {

}

func (a *AmadeusService) CreateHotelBooking(requestBody amadeusHotelModels.HotelBookingRequest) (*amadeusHotelModels.HotelOrderResponse, error) {
	booking, err := a.client.CreateHotelBooking(requestBody)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (a *AmadeusService) FetchHotelRatings(hotelIDs []string) (*amadeusHotelModels.HotelSentimentResponse, error) {
	ratings, err := a.client.FetchHotelRatings(hotelIDs)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (a *AmadeusService) HotelNameAutoComplete(keyword string, subtype string) (*amadeusHotelModels.HotelNameResponse, error) {
	hotels, err := a.client.HotelNameAutoComplete(keyword, subtype)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
