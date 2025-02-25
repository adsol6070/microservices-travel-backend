package request

type HotelOrderRequest struct {
	Type        string      `json:"type" validate:"required"`
	Guests      []Guest     `json:"guests" validate:"required,dive"`
	TravelAgent TravelAgent `json:"travelAgent"`
	RoomAssoc   []RoomAssoc `json:"roomAssociations" validate:"required,dive"`
	Payment     Payment     `json:"payment"`
}

type Guest struct {
	TID       int    `json:"tid" validate:"required"`
	Title     string `json:"title" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

type TravelAgent struct {
	Contact Contact `json:"contact"`
}

type Contact struct {
	Email string `json:"email"`
}

type RoomAssoc struct {
	GuestRefs    []GuestRef `json:"guestReferences" validate:"required,dive"`
	HotelOfferID string     `json:"hotelOfferId" validate:"required"`
}

type GuestRef struct {
	GuestReference string `json:"guestReference"`
}

type Payment struct {
	Method      string      `json:"method"`
	PaymentCard PaymentCard `json:"paymentCard"`
}

type PaymentCard struct {
	PaymentCardInfo PaymentCardInfo `json:"paymentCardInfo"`
}

type PaymentCardInfo struct {
	VendorCode string `json:"vendorCode"`
	CardNumber string `json:"cardNumber"`
	ExpiryDate string `json:"expiryDate"`
	HolderName string `json:"holderName"`
}
