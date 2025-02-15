package amadeusHotelModels

type HotelNameResponse struct {
	Data []HotelNameData `json:"data"`
}

type HotelNameData struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	IATACode  string   `json:"iataCode"`
	SubType   string   `json:"subType"`
	Relevance int      `json:"relevance"`
	Type      string   `json:"type"`
	HotelIDs  []string `json:"hotelIds"`
	Address   AddressResponse  `json:"address"`
	GeoCode   GeoCodeResponse  `json:"geoCode"`
}

type AddressResponse struct {
	CityName    string `json:"cityName"`
	CountryCode string `json:"countryCode"`
}

type GeoCodeResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
