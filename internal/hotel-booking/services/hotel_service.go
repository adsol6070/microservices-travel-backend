package services

import (
	"errors"
	"log"
	"microservices-travel-backend/internal/hotel-booking/domain/mapper"
	"microservices-travel-backend/internal/hotel-booking/domain/models"
	"microservices-travel-backend/internal/hotel-booking/domain/ports"
)

type HotelService struct {
	db          ports.HotelDB         // Local database interface
	providers   []ports.HotelProvider // External providers interface
	hotelMapper *mapper.HotelMapper   // Dependency injected mapper
}

// NewHotelService initializes and returns a new HotelService instance.
func NewHotelService(db ports.HotelDB, providers []ports.HotelProvider, hotelMapper *mapper.HotelMapper) *HotelService {
	return &HotelService{
		db:          db,
		providers:   providers,
		hotelMapper: hotelMapper,
	}
}

// FetchAndCacheHotels fetches multiple hotels from external providers, maps their responses to local format, and caches them in the local database.
func (s *HotelService) FetchHotels() ([]models.Hotel, error) {
	var allFetchedHotels []models.Hotel
	hotelCache := make(map[string]models.Hotel) // To avoid duplicates across providers

	// Iterate through the list of external providers.
	for _, provider := range s.providers {
		log.Printf("Fetching hotel data from provider: %T\n", provider)

		// Fetch hotel details from the provider.
		providerHotels, err := provider.GetHotels()
		if err != nil {
			log.Printf("Provider %T failed to fetch hotels: %v\n", provider, err)
			continue
		}

		// Map the provider-specific hotels to the local hotel format.
		for _, externalHotel := range providerHotels {
			mappedHotel := s.hotelMapper.MapToLocalHotelFormat(externalHotel)

			// Avoid duplicates using the hotel ID.
			if _, exists := hotelCache[mappedHotel.ID]; !exists {
				hotelCache[mappedHotel.ID] = mappedHotel
				allFetchedHotels = append(allFetchedHotels, mappedHotel)
			}
		}
	}

	if len(allFetchedHotels) == 0 {
		return nil, errors.New("no hotels found from any provider")
	}

	// Cache all unique hotels into the local database.
	for _, hotel := range allFetchedHotels {
		err := s.db.SaveHotel(&hotel)
		if err != nil {
			log.Printf("Failed to save hotel %s to local DB: %v\n", hotel.ID, err)
		}
	}

	log.Printf("Successfully cached %d hotels in local database.\n", len(allFetchedHotels))
	return allFetchedHotels, nil
}
