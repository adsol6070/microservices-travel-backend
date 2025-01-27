package hotel_provider

import (
	"fmt"
	"time"
)

type AmadeusAdapter struct {
	apiKey string
}

// NewAmadeusAdapter creates a new AmadeusAdapter
func NewAmadeusAdapter(apiKey string) *AmadeusAdapter {
	return &AmadeusAdapter{apiKey: apiKey}
}

// GetHotelDetails fetches hotel details from Amadeus and returns raw data as map
func (a *AmadeusAdapter) GetHotelDetails(hotelID string) (map[string]interface{}, error) {
	// Mock API response
	fmt.Printf("Fetching hotel details for %s from Amadeus...\n", hotelID)
	time.Sleep(1 * time.Second) // Simulating network delay

	// Raw response in map format
	rawResponse := map[string]interface{}{
		"hotel_id":   hotelID,
		"hotel_name": "Amadeus Hotel",
		"location": map[string]string{
			"city":    "Rome",
			"country": "Italy",
		},
		"rating":     4.7,
		"facilities": []string{"Free WiFi", "Restaurant", "Parking"},
		"price": map[string]interface{}{
			"min_price": 159.99,
			"max_price": 329.99,
			"currency":  "EUR",
		},
		"availability": map[string]int{
			"rooms_available": 15,
		},
		"policies": map[string]string{
			"cancellation":      "Free cancellation up to 48 hours before check-in",
			"check_in_time":     "3:00 PM",
			"check_out_time":    "12:00 PM",
			"smoking_policy":    "Smoking rooms available",
			"child_policy":      "Children allowed with extra charge",
			"extra_beds_policy": "Extra beds not allowed",
		},
		"payment_methods": []string{"Credit Card", "Bank Transfer"},
		"provider_metadata": map[string]interface{}{
			"provider_name":   "Amadeus",
			"provider_id":     hotelID,
			"provider_rating": 4.6,
			"last_updated":    "2025-01-23T10:00:00Z",
		},
	}

	return rawResponse, nil
}

// SearchHotels searches for hotels in a given location from Amadeus
func (a *AmadeusAdapter) SearchHotels(location, checkInDate, checkOutDate string) ([]map[string]interface{}, error) {
	// Mock API response
	fmt.Printf("Searching for hotels in %s from Amadeus...\n", location)
	time.Sleep(2 * time.Second)

	rawResponse := []map[string]interface{}{
		{
			"hotel_id":   "2",
			"hotel_name": "Amadeus Hotel 1",
			"location": map[string]string{
				"city":    "Rome",
				"country": "Italy",
			},
			"price": map[string]interface{}{
				"min_price": 159.99,
				"max_price": 259.99,
				"currency":  "EUR",
			},
			"availability": map[string]int{
				"rooms_available": 25,
			},
		},
	}

	return rawResponse, nil
}
