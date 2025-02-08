package models

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

type Description struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Guests struct {
	Adults int `json:"adults"`
}

type Price struct {
	Currency   string     `json:"currency"`
	Base       string     `json:"base"`
	Total      string     `json:"total"`
	Variations Variations `json:"variations"`
}

type Variations struct {
	Average PriceDetail   `json:"average"`
	Changes []PriceChange `json:"changes"`
}

type PriceDetail struct {
	Base string `json:"base"`
}

type PriceChange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Base      string `json:"base"`
}

type Policies struct {
	Cancellations []Cancellation `json:"cancellations"`
	PaymentType   string         `json:"paymentType"`
	Refundable    Refundable     `json:"refundable"`
}

type Cancellation struct {
	NumberOfNights int    `json:"numberOfNights"`
	Deadline       string `json:"deadline"`
	Amount         string `json:"amount"`
	PolicyType     string `json:"policyType"`
}

type Refundable struct {
	CancellationRefund string `json:"cancellationRefund"`
}

type Warning struct {
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Source struct {
		Parameter string `json:"parameter"`
	} `json:"source"`
}
