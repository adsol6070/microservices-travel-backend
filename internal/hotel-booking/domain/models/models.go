package models

import "microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/amadeusHotelModels"

type PlaceReview struct {
	AuthorName string `json:"authorName"`
	Rating     int    `json:"rating"`
	Text       string `json:"text"`
}

type EnrichedHotelOffer struct {
	HotelID           string                     `json:"hotelId"`
	Name              string                     `json:"name"`
	Latitude          *float64                   `json:"latitude"`
	Longitude         *float64                   `json:"longitude"`
	Rating            float64                    `json:"rating,omitempty"`
	Reviews           []PlaceReview              `json:"reviews,omitempty"`
	PhotoURL          string                     `json:"photoUrl"`
	Offers            []amadeusHotelModels.Offer `json:"offers"`
	GooglePlaceID     string                     `json:"googlePlaceId"`
	GoogleDisplayName string                     `json:"googleDisplayName"`
}
