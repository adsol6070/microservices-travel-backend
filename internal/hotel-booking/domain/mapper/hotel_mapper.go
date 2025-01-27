package mapper

import (
	"microservices-travel-backend/internal/hotel-booking/domain/models"
)

// HotelMapper is responsible for mapping external provider data to the local hotel format.
type HotelMapper struct{}

// NewHotelMapper creates a new instance of HotelMapper.
func NewHotelMapper() *HotelMapper {
	return &HotelMapper{}
}

// MapToLocalHotelFormat maps a provider-specific hotel structure to the local hotel structure.
func (m *HotelMapper) MapToLocalHotelFormat(externalHotel map[string]interface{}) models.Hotel {
	// Utility functions for safely getting values from the map
	getString := func(key string) string {
		if val, ok := externalHotel[key].(string); ok {
			return val
		}
		return ""
	}
	getFloat64 := func(key string) float64 {
		if val, ok := externalHotel[key].(float64); ok {
			return val
		}
		return 0.0
	}
	getInt := func(key string) int {
		if val, ok := externalHotel[key].(int); ok {
			return val
		}
		return 0
	}
	getStringSlice := func(key string) []string {
		if val, ok := externalHotel[key].([]string); ok {
			return val
		}
		return []string{}
	}

	// Map nested structures
	location := models.Location{
		City:       getString("city"),
		Country:    getString("country"),
		Latitude:   getFloat64("latitude"),
		Longitude:  getFloat64("longitude"),
		Address:    getString("address"),
		PostalCode: getString("postal_code"),
	}

	priceRange := models.PriceRange{
		MinPrice: getFloat64("min_price"),
		MaxPrice: getFloat64("max_price"),
		Currency: getString("currency"),
	}

	policies := models.Policies{
		Cancellation:    getString("cancellation_policy"),
		CheckInTime:     getString("check_in_time"),
		CheckOutTime:    getString("check_out_time"),
		SmokingPolicy:   getString("smoking_policy"),
		ChildPolicy:     getString("child_policy"),
		ExtraBedsPolicy: getString("extra_beds_policy"),
	}

	availability := models.Availability{
		AvailableRooms: getInt("available_rooms"),
		TotalRooms:     getInt("total_rooms"),
	}

	// Utility functions for safely getting values from a nested map
	getStringFromMap := func(m map[string]interface{}, key string) string {
		if val, ok := m[key].(string); ok {
			return val
		}
		return ""
	}
	getFloat64FromMap := func(m map[string]interface{}, key string) float64 {
		if val, ok := m[key].(float64); ok {
			return val
		}
		return 0.0
	}
	getIntFromMap := func(m map[string]interface{}, key string) int {
		if val, ok := m[key].(int); ok {
			return val
		}
		return 0
	}
	getBoolFromMap := func(m map[string]interface{}, key string) bool {
		if val, ok := m[key].(bool); ok {
			return val
		}
		return false
	}
	getStringSliceFromMap := func(m map[string]interface{}, key string) []string {
		if val, ok := m[key].([]string); ok {
			return val
		}
		return []string{}
	}

	// Handle rooms mapping
	var rooms []models.Room
	if externalRooms, ok := externalHotel["rooms"].([]map[string]interface{}); ok {
		for _, extRoom := range externalRooms {
			room := models.Room{
				ID:           getStringFromMap(extRoom, "id"),
				Type:         getStringFromMap(extRoom, "type"),
				Capacity:     getIntFromMap(extRoom, "capacity"),
				Price:        getFloat64FromMap(extRoom, "price"),
				BedType:      getStringFromMap(extRoom, "bed_type"),
				Availability: getBoolFromMap(extRoom, "availability"),
				Images:       getStringSliceFromMap(extRoom, "images"),
				Facilities:   getStringSliceFromMap(extRoom, "facilities"),
			}
			rooms = append(rooms, room)
		}
	}

	// Map metadata and return the local hotel model
	providerMetadata := models.ProviderMetadata{
		ProviderName:   getString("provider_name"),
		ProviderID:     getString("provider_id"),
		ProviderRating: getFloat64("provider_rating"),
		LastUpdated:    getString("last_updated"),
	}

	return models.Hotel{
		ID:               getString("hotel_id"),
		Name:             getString("name"),
		Brand:            getString("brand"),
		Location:         location,
		Rating:           getFloat64("rating"),
		Facilities:       getStringSlice("facilities"),
		RoomTypes:        rooms,
		Images:           getStringSlice("images"),
		PriceRange:       priceRange,
		Policies:         policies,
		Availability:     availability,
		PaymentMethods:   getStringSlice("payment_methods"),
		ProviderMetadata: providerMetadata,
	}
}
