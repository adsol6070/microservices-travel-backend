package hotel_provider

import (
	"fmt"
	"time"
)

type ExpediaAdapter struct {
	apiKey string
}

// NewExpediaAdapter creates a new ExpediaAdapter
func NewExpediaAdapter(apiKey string) *ExpediaAdapter {
	return &ExpediaAdapter{apiKey: apiKey}
}

// GetHotels fetches a list of hotel details from Expedia and returns raw data as a slice of maps
func (e *ExpediaAdapter) GetHotels() ([]map[string]interface{}, error) {
	// Mock API response - in production, you'd call Expedia's API here
	fmt.Println("Fetching list of hotels from Expedia...")
	time.Sleep(2 * time.Second) // Simulating network delay

	// Raw response in slice format for multiple hotels
	rawResponse := []map[string]interface{}{
		{
			"hotel_id":   "1",
			"hotel_name": "Expedia Hotel Paris",
			"location": map[string]string{
				"city":    "Paris",
				"country": "France",
			},
			"rating":     4.5,
			"facilities": []string{"Free WiFi", "Pool", "Gym"},
			"price": map[string]interface{}{
				"min_price": 199.99,
				"max_price": 399.99,
				"currency":  "USD",
			},
			"availability": map[string]int{
				"rooms_available": 10,
			},
			"policies": map[string]string{
				"cancellation":      "Free cancellation up to 24 hours before check-in",
				"check_in_time":     "2:00 PM",
				"check_out_time":    "11:00 AM",
				"smoking_policy":    "Non-smoking rooms available",
				"child_policy":      "Children allowed",
				"extra_beds_policy": "Extra beds allowed",
			},
			"payment_methods": []string{"Credit Card", "PayPal"},
			"provider_metadata": map[string]interface{}{
				"provider_name":   "Expedia",
				"provider_id":     "1",
				"provider_rating": 4.2,
				"last_updated":    "2025-01-23T10:00:00Z",
			},
		},
		{
			"hotel_id":   "2",
			"hotel_name": "Expedia Hotel London",
			"location": map[string]string{
				"city":    "London",
				"country": "United Kingdom",
			},
			"rating":     4.7,
			"facilities": []string{"Free Breakfast", "Parking", "Spa"},
			"price": map[string]interface{}{
				"min_price": 249.99,
				"max_price": 499.99,
				"currency":  "GBP",
			},
			"availability": map[string]int{
				"rooms_available": 5,
			},
			"policies": map[string]string{
				"cancellation":      "Non-refundable",
				"check_in_time":     "3:00 PM",
				"check_out_time":    "12:00 PM",
				"smoking_policy":    "No smoking allowed",
				"child_policy":      "Children under 12 free",
				"extra_beds_policy": "Extra beds not allowed",
			},
			"payment_methods": []string{"Credit Card", "Bank Transfer"},
			"provider_metadata": map[string]interface{}{
				"provider_name":   "Expedia",
				"provider_id":     "2",
				"provider_rating": 4.5,
				"last_updated":    "2025-01-22T12:00:00Z",
			},
		},
		// Add more hotels as needed
	}

	// Return the raw response
	return rawResponse, nil
}
