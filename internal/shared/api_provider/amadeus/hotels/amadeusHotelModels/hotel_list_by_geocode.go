package amadeusHotelModels

type HotelsByGeoCodeReq struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type HotelsByGeoCodeResp struct {
	Data []Hotel `json:"data"`
}