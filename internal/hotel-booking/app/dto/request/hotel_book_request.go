package request

type HotelBookingRequest struct {
	Data HotelBookRequest `json:"data" validate:"required"`
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
	Contact struct {
		Email string `json:"email" validate:"required,email"`
	} `json:"contact"`
}

type GuestReference struct {
	GuestReference string `json:"guestReference" validate:"required"`
}

type RoomAssociation struct {
	GuestReferences []GuestReference `json:"guestReferences" validate:"required,dive"`
	HotelOfferID    string           `json:"hotelOfferId" validate:"required"`
}

type PaymentCardInfo struct {
	VendorCode string `json:"vendorCode" validate:"required"`
	CardNumber string `json:"cardNumber" validate:"required,len=16"`
	ExpiryDate string `json:"expiryDate" validate:"required"`
	HolderName string `json:"holderName" validate:"required"`
}

type PaymentCard struct {
	PaymentCardInfo PaymentCardInfo `json:"paymentCardInfo"`
}

type Payment struct {
	Method       string      `json:"method" validate:"required"`
	PaymentCard  PaymentCard `json:"paymentCard"`
}

type HotelBookRequest struct {
	Type             string            `json:"type" validate:"required"`
	Guests           []Guest           `json:"guests" validate:"required,dive"`
	TravelAgent      TravelAgent       `json:"travelAgent"`
	RoomAssociations []RoomAssociation `json:"roomAssociations" validate:"required,dive"`
	Payment          Payment           `json:"payment" validate:"required"`
}
