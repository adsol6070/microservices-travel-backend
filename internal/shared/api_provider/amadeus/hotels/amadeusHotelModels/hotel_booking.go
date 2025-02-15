package amadeusHotelModels

type HotelBookingRequest struct {
	Data HotelOrderData `json:"data"`
}

type HotelBookingResponse struct {
	Data     HotelOrderData `json:"data"`
	Warnings []APIWarning   `json:"warnings"`
}

type HotelOrderData struct {
	Type              string             `json:"type"`
	ID                string             `json:"id"`
	Guests            []Guest            `json:"guests"`
	TravelAgent       TravelAgent        `json:"travelAgent"`
	RoomAssociations  []RoomAssociation  `json:"roomAssociations"`
	Payment           PaymentDetails     `json:"payment"`
	HotelBookings     []HotelBooking     `json:"hotelBookings"`
	AssociatedRecords []AssociatedRecord `json:"associatedRecords"`
	Self              string             `json:"self"`
}

type TravelAgent struct {
	Contact Contact `json:"contact"`
}

type Contact struct {
	Email string `json:"email"`
}

type HotelBooking struct {
	Type                     string              `json:"type"`
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

type GuestBookReference struct {
	GuestReference string `json:"guestReference"`
}

type HotelBookOffer struct {
	ID           string      `json:"id"`
	Type         string      `json:"type"`
	Category     string      `json:"category"`
	CheckInDate  string      `json:"checkInDate"`
	CheckOutDate string      `json:"checkOutDate"`
	Guests       Guests      `json:"guests"`
	Policies     Policies    `json:"policies"`
	Price        Price       `json:"price"`
	RateCode     string      `json:"rateCode"`
	Room         RoomDetails `json:"room"`
	RoomQuantity int         `json:"roomQuantity"`
}

type RoomDetails struct {
	Description Description `json:"description"`
	Type        string      `json:"type"`
}

type Tax struct {
	Amount           string `json:"amount"`
	Code             string `json:"code"`
	Currency         string `json:"currency"`
	Included         bool   `json:"included"`
	PricingFrequency string `json:"pricingFrequency"`
	PricingMode      string `json:"pricingMode"`
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
	PaymentCardInfo PaymentCardInfo `json:"paymentCardInfo"`
}

type PaymentCardInfo struct {
	VendorCode string `json:"vendorCode"`
	CardNumber string `json:"cardNumber"`
	ExpiryDate string `json:"expiryDate"`
	HolderName string `json:"holderName"`
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

type AssociatedRecord struct {
	Reference        string `json:"reference"`
	OriginSystemCode string `json:"originSystemCode"`
}

type APIWarning struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
}

