package amadeus

import (
	"encoding/json"
	"errors"
	"log"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/models"
	"sync"
)

type AmadeusService struct {
	client *hotels.AmadeusClient
}

type SearchHotelsRequest struct {
	CityCode     string `json:"cityCode"`
	CheckInDate  string `json:"checkInDate"`
	CheckOutDate string `json:"checkOutDate"`
	Rooms        int    `json:"rooms"`
	Persons      int    `json:"persons"`
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

func (a *AmadeusService) SearchHotels(req SearchHotelsRequest) ([]models.HotelOffer, error) {
	log.Printf("INFO: Starting SearchHotels - CityCode: %s, CheckIn: %s, CheckOut: %s, Rooms: %d, Persons: %d",
		req.CityCode, req.CheckInDate, req.CheckOutDate, req.Rooms, req.Persons)

	// Fetch hotels for the given city code
	hotels, err := a.client.HotelSearch(req.CityCode)
	if err != nil {
		log.Println("ERROR: Failed to fetch hotels -", err)
		return nil, err
	}

	log.Printf("INFO: Found %d hotels for city code: %s", len(hotels), req.CityCode)

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

	offersJSON, err := json.MarshalIndent(hotelOffers, "", "  ")
	if err != nil {
		log.Printf("ERROR: Failed to marshal hotel offers: %v", err)
	} else {
		log.Printf("INFO: Retrieved %d hotel offers: %s", len(hotelOffers), string(offersJSON))
	}

	if len(hotelOffers) == 0 {
		log.Println("WARN: No hotel offers available")
		return nil, errors.New("no hotel offers available")
	}

	// Filter hotel offers based on search criteria
	filteredHotels := filterHotelsByRequest(hotelOffers, req)
	log.Printf("INFO: %d hotels matched the search criteria", len(filteredHotels))

	if len(filteredHotels) == 0 {
		log.Println("WARN: No hotels matched the search criteria")
		return nil, errors.New("no matching hotels found")
	}

	log.Println("INFO: SearchHotels execution completed successfully")

	return filteredHotels, nil
}

func extractHotelIDs(hotels []models.HotelData) []string {
	var hotelIDs []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, hotel := range hotels {
		wg.Add(1)
		go func(h models.HotelData) {
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

func filterHotelsByRequest(hotelOffers []models.HotelOffer, req SearchHotelsRequest) []models.HotelOffer {
	var filteredHotels []models.HotelOffer

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
