package amadeusHotelModels

type HotelsByIDReq struct {
	HotelIDs []string `json:"hotelIds"`
}

type HotelsByIDResp struct {
	Data []Hotel `json:"data"`
}
