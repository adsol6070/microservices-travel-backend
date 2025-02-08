package models

type HotelBookingResponse struct {
	Data     BookingData  `json:"data"`
	Warnings []APIWarning `json:"warnings"`
}

type BookingData struct {
	ID                string             `json:"id"`
	HotelBookings     []HotelBooking     `json:"hotelBookings"`
	Guests            []Guest            `json:"guests"`
	AssociatedRecords []AssociatedRecord `json:"associatedRecords"`
	Self              string             `json:"self"`
}

type HotelBooking struct {
	ID                       string              `json:"id"`
	BookingStatus            string              `json:"bookingStatus"`
	HotelProviderInformation []HotelProviderInfo `json:"hotelProviderInformation"`
	RoomAssociations         []RoomAssociation   `json:"roomAssociations"`
	HotelOffer               HotelOffer          `json:"hotelOffer"`
	Hotel                    HotelInfo           `json:"hotel"`
	Payment                  PaymentDetails      `json:"payment"`
	TravelAgentID            string              `json:"travelAgentId"`
}

type HotelProviderInfo struct {
	HotelProviderCode  string `json:"hotelProviderCode"`
	ConfirmationNumber string `json:"confirmationNumber"`
}

type RoomAssociation struct {
	HotelOfferID    string               `json:"hotelOfferId"`
	GuestReferences []GuestBookReference `json:"guestReferences"`
}

type HotelBookOffer struct {
	ID           string `json:"id"`
	Category     string `json:"category"`
	CheckInDate  string `json:"checkInDate"`
	CheckOutDate string `json:"checkOutDate"`
	Guests       struct {
		Adults int `json:"adults"`
	} `json:"guests"`
	Policies Policies `json:"policies"`
	Price    Price    `json:"price"`
	RateCode string   `json:"rateCode"`
	Room     Room     `json:"room"`
}

type PoliciesBook struct {
	Cancellations []Cancellation `json:"cancellations"`
	PaymentType   string         `json:"paymentType"`
	Refundable    struct {
		CancellationRefund string `json:"cancellationRefund"`
	} `json:"refundable"`
}

type CancellationBook struct {
	Amount     string `json:"amount"`
	Deadline   string `json:"deadline"`
	PolicyType string `json:"policyType"`
}

type PriceBook struct {
	Base         string `json:"base"`
	Currency     string `json:"currency"`
	SellingTotal string `json:"sellingTotal"`
	Taxes        []Tax  `json:"taxes"`
	Total        string `json:"total"`
}

type Tax struct {
	Amount           string `json:"amount"`
	Code             string `json:"code"`
	Currency         string `json:"currency"`
	Included         bool   `json:"included"`
	PricingFrequency string `json:"pricingFrequency"`
	PricingMode      string `json:"pricingMode"`
}

type RoomBook struct {
	Description struct {
		Lang string `json:"lang"`
		Text string `json:"text"`
	} `json:"description"`
	Type string `json:"type"`
}

type HotelInfo struct {
	HotelID   string `json:"hotelId"`
	ChainCode string `json:"chainCode"`
	Name      string `json:"name"`
	Self      string `json:"self"`
}

type PaymentDetails struct {
	Method      string `json:"method"`
	PaymentCard Card   `json:"paymentCard"`
}

type Card struct {
	PaymentCardInfo struct {
		VendorCode string `json:"vendorCode"`
		CardNumber string `json:"cardNumber"`
		ExpiryDate string `json:"expiryDate"`
		HolderName string `json:"holderName"`
	} `json:"paymentCardInfo"`
}

type Guest struct {
	TID       int    `json:"tid"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type GuestBookReference struct {
	GuestID int `json:"guestId"`
}

type AssociatedRecord struct {
	Reference        string `json:"reference"`
	OriginSystemCode string `json:"originSystemCode"`
}

type APIWarning struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
}
