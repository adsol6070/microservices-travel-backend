package amadeusHotelModels

type HotelBookingRequest struct {
	Data HotelOrderData `json:"data"`
}

type HotelOrderData struct {
	Type             string            `json:"type"`
	Guests           []Guest           `json:"guests"`
	TravelAgent      TravelAgent       `json:"travelAgent"`
	RoomAssociations []RoomAssociation `json:"roomAssociations"`
	Payment          PaymentDetails    `json:"payment"`
}

type Guest struct {
	TID       int    `json:"tid"`
	Title     string `json:"title"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type TravelAgent struct {
	Contact Contact `json:"contact"`
}

type Contact struct {
	Email string `json:"email"`
}

type RoomAssociation struct {
	HotelOfferID    string               `json:"hotelOfferId"`
	GuestReferences []GuestBookReference `json:"guestReferences"`
}

type GuestBookReference struct {
	GuestReference string `json:"guestReference"`
}

type PaymentDetails struct {
	Method      string `json:"method"`
	PaymentCard Card   `json:"paymentCard"`
}

type Card struct {
	PaymentCardInfo PaymentCardInfo `json:"paymentCardInfo"`
}

type PaymentCardInfo struct {
	VendorCode string `json:"vendorCode"`
	CardNumber string `json:"cardNumber"`
	ExpiryDate string `json:"expiryDate"`
	HolderName string `json:"holderName"`
}