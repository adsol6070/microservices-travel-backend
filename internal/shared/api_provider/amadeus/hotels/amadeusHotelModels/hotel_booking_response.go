package amadeusHotelModels

// HotelOrderResponse represents the overall response structure.
type HotelOrderResponse struct {
	Data     HotelOrderResponseData `json:"data"`
	Warnings []WarningDetails       `json:"warnings"`
}

// HotelOrderData contains the main booking details.
type HotelOrderResponseData struct {
	OrderType         string                `json:"type"`
	OrderID           string                `json:"id"`
	HotelBookings     []HotelBookingDetails `json:"hotelBookings"`
	GuestList         []GuestInfo           `json:"guests"`
	AssociatedRecords []RecordDetails       `json:"associatedRecords"`
	SelfLink          string                `json:"self"`
}

// HotelBookingDetails represents individual hotel bookings.
type HotelBookingDetails struct {
	BookingType       string                      `json:"type"`
	BookingID         string                      `json:"id"`
	BookingStatus     string                      `json:"bookingStatus"`
	ProviderInfo      []HotelProviderInfoResponse `json:"hotelProviderInformation"`
	RoomAssignments   []RoomAssociation           `json:"roomAssociations"`
	HotelOfferDetails HotelResponseOffer          `json:"hotelOffer"`
	HotelInfo         HotelDetails                `json:"hotel"`
	PaymentDetails    PaymentInfo                 `json:"payment"`
	TravelAgentID     string                      `json:"travelAgentId"`
}

// HotelProviderInfo represents provider details.
type HotelProviderInfoResponse struct {
	ProviderCode       string `json:"hotelProviderCode"`
	ConfirmationNumber string `json:"confirmationNumber"`
}

// RoomAssociation represents the mapping of guests to room bookings.
type RoomAssociationResponse struct {
	HotelOfferID    string           `json:"hotelOfferId"`
	GuestReferences []GuestReference `json:"guestReferences"`
}

// GuestReference holds references for guest identification.
type GuestReference struct {
	Reference string `json:"guestReference"`
}

// HotelOffer contains details about the booked hotel offer.
type HotelResponseOffer struct {
	OfferID      string          `json:"id"`
	OfferType    string          `json:"type"`
	Category     string          `json:"category"`
	CheckInDate  string          `json:"checkInDate"`
	CheckOutDate string          `json:"checkOutDate"`
	Guests       GuestCount      `json:"guests"`
	Policies     BookingPolicies `json:"policies"`
	Pricing      PricingDetails  `json:"price"`
	RateCode     string          `json:"rateCode"`
	RoomDetails  RoomInfo        `json:"room"`
	RoomQuantity int             `json:"roomQuantity"`
}

// GuestCount represents the number of guests.
type GuestCount struct {
	Adults int `json:"adults"`
}

// BookingPolicies contains cancellation and refund policies.
type BookingPolicies struct {
	Cancellations []CancellationPolicy `json:"cancellations"`
	PaymentType   string               `json:"paymentType"`
	Refundable    RefundPolicy         `json:"refundable"`
}

// CancellationPolicy represents cancellation details.
type CancellationPolicy struct {
	Amount     string `json:"amount"`
	Deadline   string `json:"deadline"`
	PolicyType string `json:"policyType"`
}

// RefundPolicy represents refund eligibility.
type RefundPolicy struct {
	CancellationRefund string `json:"cancellationRefund"`
}

// PricingDetails holds pricing-related information.
type PricingDetails struct {
	BaseAmount   string          `json:"base"`
	Currency     string          `json:"currency"`
	SellingTotal string          `json:"sellingTotal"`
	Taxes        []TaxDetails    `json:"taxes"`
	TotalAmount  string          `json:"total"`
	Variations   PriceVariations `json:"variations"`
}

// TaxDetails represents tax breakdown.
type TaxDetails struct {
	Amount           string `json:"amount"`
	Code             string `json:"code"`
	Currency         string `json:"currency"`
	Included         bool   `json:"included"`
	PricingFrequency string `json:"pricingFrequency"`
	PricingMode      string `json:"pricingMode"`
}

// PriceVariations represents price changes.
type PriceVariations struct {
	Changes []PriceChange `json:"changes"`
}

// PriceChange represents price change over time.
type PriceChangeResponse struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Base      string `json:"base"`
	Currency  string `json:"currency"`
}

// RoomInfo contains room description.
type RoomInfo struct {
	Description RoomDescription `json:"description"`
	RoomType    string          `json:"type"`
}

// RoomDescription contains room description details.
type RoomDescription struct {
	Language string `json:"lang"`
	Text     string `json:"text"`
}

// HotelDetails represents hotel information.
type HotelDetails struct {
	HotelID   string `json:"hotelId"`
	ChainCode string `json:"chainCode"`
	Name      string `json:"name"`
	SelfURL   string `json:"self"`
}

// PaymentInfo holds payment details.
type PaymentInfo struct {
	Method      string      `json:"method"`
	PaymentCard PaymentCard `json:"paymentCard"`
}

// PaymentCard holds card-related information.
type PaymentCard struct {
	CardInfo CardDetails `json:"paymentCardInfo"`
}

// CardDetails contains payment card information.
type CardDetails struct {
	VendorCode string `json:"vendorCode"`
	CardNumber string `json:"cardNumber"`
	ExpiryDate string `json:"expiryDate"`
	HolderName string `json:"holderName"`
}

// GuestInfo holds guest details.
type GuestInfo struct {
	TID       int    `json:"tid"`
	GuestID   int    `json:"id"`
	Title     string `json:"title"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

// RecordDetails represents associated records.
type RecordDetails struct {
	Reference        string `json:"reference"`
	OriginSystemCode string `json:"originSystemCode"`
}

// WarningDetails represents warnings in the response.
type WarningDetails struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
}
