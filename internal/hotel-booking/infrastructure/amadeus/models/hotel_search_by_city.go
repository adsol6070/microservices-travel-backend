package models

// Hotel represents basic hotel details returned from Amadeus API.
type Hotel struct {
	HotelID     string  `json:"hotelId"`
	ChainCode   string  `json:"chainCode"`
	IATACode    string  `json:"iataCode"`
	DupeID      int     `json:"dupeId"`
	Name        string  `json:"name"`
	GeoCode     GeoCode `json:"geoCode"`
	Address     Address `json:"address"`
	LastUpdate  string  `json:"lastUpdate"`
}

// GeoCode represents latitude and longitude of a hotel.
type GeoCode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Address represents the address details of a hotel.
type Address struct {
	CountryCode string `json:"countryCode"`
}

type HotelResponse struct {
	HotelID     string  `json:"id"`
	ChainCode   string  `json:"chainCode"`
	IATACode    string  `json:"iataCode"`
	CountryCode string  `json:"countryCode"`
	DupeID      int     `json:"dupeId"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}