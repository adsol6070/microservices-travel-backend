package request

type HotelDetailsRequest struct {
	HotelID string `json:"hotelId"`
	GooglePlaceID string `json:"googlePlaceId"`
}