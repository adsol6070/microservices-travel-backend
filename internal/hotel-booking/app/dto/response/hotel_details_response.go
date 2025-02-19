package response

type HotelDetails struct {
	HotelID       string               `json:"hotelId"`
	HotelName     string               `json:"hotelName"`
	CheckInDate   string               `json:"checkInDate"`
	CheckOutDate  string               `json:"checkOutDate"`
	Price         PriceDetails         `json:"price"`
	Location      LocationDetails      `json:"location"`
	Contact       ContactDetails       `json:"contact"`
	Rating        RatingDetails        `json:"rating"`
	Photos        []string             `json:"photos"`
	Accessibility AccessibilityOptions `json:"accessibilityOptions"`
	Payment       PaymentOptions       `json:"paymentOptions"`
	Amenities     Amenities            `json:"amenities"`
	Policies      HotelPolicies        `json:"policies"`
}

type PriceDetails struct {
	Currency     string  `json:"currency"`
	Base         float64 `json:"base"`
	TaxesAndFees float64 `json:"taxesAndFees"`
	Total        float64 `json:"total"`
}

type LocationDetails struct {
	CityCode      string      `json:"cityCode"`
	Address       Address     `json:"address"`
	Coordinates   Coordinates `json:"coordinates"`
	GoogleMapsURI string      `json:"googleMapsUri"`
	PlaceID       string      `json:"placeId"`
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ContactDetails struct {
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Website     string `json:"website"`
}

type RatingDetails struct {
	Stars            int     `json:"stars"`
	UserReviewsCount int     `json:"userReviewsCount"`
	AverageRating    float64 `json:"averageUserRating"`
}

type AccessibilityOptions struct {
	WheelchairAccessible bool `json:"wheelchairAccessible"`
	ElevatorAvailable    bool `json:"elevatorAvailable"`
	BrailleSignage       bool `json:"brailleSignage"`
	HearingAidSupport    bool `json:"hearingAidSupport"`
}

type PaymentOptions struct {
	CardsAccepted      []string           `json:"cardsAccepted"`
	CashAccepted       bool               `json:"cashAccepted"`
	DigitalWallets     []string           `json:"digitalWalletsAccepted"`
	CancellationPolicy CancellationPolicy `json:"cancellationPolicy"`
}

type CancellationPolicy struct {
	FreeCancellationUntil string  `json:"freeCancellationUntil"`
	CancellationFee       float64 `json:"cancellationFee"`
}

type Amenities struct {
	Wifi        bool    `json:"wifi"`
	Parking     Parking `json:"parking"`
	Pool        bool    `json:"pool"`
	Spa         bool    `json:"spa"`
	Gym         bool    `json:"gym"`
	Restaurant  bool    `json:"restaurant"`
	RoomService bool    `json:"roomService"`
	PetFriendly bool    `json:"petFriendly"`
}

type Parking struct {
	Available  bool    `json:"available"`
	CostPerDay float64 `json:"costPerDay"`
}

type HotelPolicies struct {
	CheckInTime    string `json:"checkInTime"`
	CheckOutTime   string `json:"checkOutTime"`
	SmokingAllowed bool   `json:"smokingAllowed"`
	MinCheckInAge  int    `json:"minCheckInAge"`
}
