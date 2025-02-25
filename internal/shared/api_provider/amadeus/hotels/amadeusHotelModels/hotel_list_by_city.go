package amadeusHotelModels

type HotelsByCityReq struct {
	CityCode string `json:"cityCode"`
}

type HotelsByCityResp struct {
	Data []Hotel `json:"data"`
}
