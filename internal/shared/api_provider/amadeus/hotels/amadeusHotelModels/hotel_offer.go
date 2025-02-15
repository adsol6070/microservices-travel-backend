package amadeusHotelModels

type HotelOffersResponse struct {
	Data     []HotelOffer `json:"data"`
	Warnings []Warning    `json:"warnings"`
}

type HotelOffer struct {
	Type      string  `json:"type"`
	Hotel     Hotel   `json:"hotel"`
	Available bool    `json:"available"`
	Offers    []Offer `json:"offers"`
	Self      string  `json:"self"`
}

type Hotel struct {
	Type      string  `json:"type"`
	HotelID   string  `json:"hotelId"`
	ChainCode string  `json:"chainCode"`
	DupeID    string  `json:"dupeId"`
	Name      string  `json:"name"`
	CityCode  string  `json:"cityCode"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Offer struct {
	ID                  string              `json:"id"`
	CheckInDate         string              `json:"checkInDate"`
	CheckOutDate        string              `json:"checkOutDate"`
	RateCode            string              `json:"rateCode"`
	RateFamilyEstimated RateFamilyEstimated `json:"rateFamilyEstimated"`
	Room                Room                `json:"room"`
	Guests              Guests              `json:"guests"`
	Price               Price               `json:"price"`
	Policies            Policies            `json:"policies"`
	Self                string              `json:"self"`
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
