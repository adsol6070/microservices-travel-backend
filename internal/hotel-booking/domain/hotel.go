package domain

// Hotel represents a generic hotel entity
type Hotel struct {
	HotelID     string  `json:"id"`
	ChainCode   string  `json:"chainCode"`
	IATACode    string  `json:"iataCode"`
	CountryCode string  `json:"countryCode"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
