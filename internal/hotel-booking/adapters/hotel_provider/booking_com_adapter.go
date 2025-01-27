package hotel_provider

import (
	"fmt"
	"time"
)

type BookingComAdapter struct {
	apiKey string
}

// NewBookingComAdapter creates a new BookingComAdapter
func NewBookingComAdapter(apiKey string) *BookingComAdapter {
	return &BookingComAdapter{apiKey: apiKey}
}

// GetHotelDetails fetches hotel details from Booking.com and returns raw data as map
func (b *BookingComAdapter) GetHotelDetails(hotelID string) (map[string]interface{}, error) {
	// Mock API response
	fmt.Printf("Fetching hotel details for %s from Booking.com...\n", hotelID)
	time.Sleep(1 * time.Second) // Simulating network delay

	rawResponse := map[string]interface{}{
		"hotel_id":   hotelID,
		"hotel_name": "Booking.com Hotel",
		"location": map[string]string{
			"city":    "New York",
			"country": "USA",
		},
		"rating":     4.3,
		"facilities": []string{"Free Breakfast", "Pet-Friendly", "Gym"},
		"price": map[string]interface{}{
			"min_price": 249.99,
			"max_price": 499.99,
			"currency":  "USD",
		},
		"availability": map[string]int{
			"rooms_available": 8,
		},
		"policies": map[string]string{
			"cancellation":      "No cancellation allowed for discounted rates",
			"check_in_time":     "1:00 PM",
			"check_out_time":    "10:00 AM",
			"smoking_policy":    "Non-smoking hotel",
			"child_policy":      "Children allowed with no extra charge",
			"extra_beds_policy": "Extra beds available for a fee",
		},
		"payment_methods": []string{"Credit Card", "PayPal", "Apple Pay"},
		"provider_metadata": map[string]interface{}{
			"provider_name":   "Booking.com",
			"provider_id":     hotelID,
			"provider_rating": 4.4,
			"last_updated":    "2025-01-23T10:00:00Z",
		},
	}

	return rawResponse, nil
}

// SearchHotels searches for hotels in a given location from Booking.com
func (b *BookingComAdapter) SearchHotels(location, checkInDate, checkOutDate string) ([]map[string]interface{}, error) {
	// Mock API response
	fmt.Printf("Searching for hotels in %s from Booking.com...\n", location)
	time.Sleep(2 * time.Second)

	rawResponse := []map[string]interface{}{
		{
			"hotel_id":   "3",
			"hotel_name": "Booking.com Hotel 1",
			"location": map[string]string{
				"city":    "New York",
				"country": "USA",
			},
			"price": map[string]interface{}{
				"min_price": 249.99,
				"max_price": 349.99,
				"currency":  "USD",
			},
			"availability": map[string]int{
				"rooms_available": 12,
			},
		},
	}

	return rawResponse, nil
}
