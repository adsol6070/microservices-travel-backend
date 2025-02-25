package amadeusHotelModels

type HotelOffersReq struct {
	HotelIDs []string `json:"hotelIds"`
	Adults   int      `json:"adults"`
}

type HotelOffersResp struct {
	Data     []HotelOffer `json:"data"`
	Warnings []Warning    `json:"warnings,omitempty"`
}

type HotelOffer struct {
	Type      string  `json:"type"`
	Hotel     Hotel   `json:"hotel"`
	Available bool    `json:"available"`
	Offers    []Offer `json:"offers"`
	Self      string  `json:"self"`
}

type RateFamilyEstimated struct {
	Code string `json:"code"`
	Type string `json:"type"`
}

type Room struct {
	Type          string `json:"type"`
	TypeEstimated struct {
		Category string `json:"category"`
		Beds     int    `json:"beds"`
		BedType  string `json:"bedType"`
	} `json:"typeEstimated"`
	Description Description `json:"description"`
}
