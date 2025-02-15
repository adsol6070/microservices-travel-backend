package amadeusHotelModels

type Guests struct {
	Adults int `json:"adults"`
}

type Description struct {
	Lang string `json:"lang"`
	Text string `json:"text"`
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
	Amount         string `json:"amount"`
	NumberOfNights int    `json:"numberOfNights"`
	Deadline       string `json:"deadline"`
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
