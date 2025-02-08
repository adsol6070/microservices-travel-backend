package models

type HotelListResponse struct {
	Data []HotelData `json:"data"`
	Meta Meta        `json:"meta"`
}

type HotelData struct {
	ChainCode  string  `json:"chainCode"`
	IATACode   string  `json:"iataCode"`
	DupeID     int     `json:"dupeId"`
	Name       string  `json:"name"`
	HotelID    string  `json:"hotelId"`
	GeoCode    GeoCode `json:"geoCode"`
	Address    Address `json:"address"`
	LastUpdate string  `json:"lastUpdate"`
}

type GeoCode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Address struct {
	CountryCode string `json:"countryCode"`
}

type Meta struct {
	Count int   `json:"count"`
	Links Links `json:"links"`
}

type Links struct {
	Self string `json:"self"`
}
