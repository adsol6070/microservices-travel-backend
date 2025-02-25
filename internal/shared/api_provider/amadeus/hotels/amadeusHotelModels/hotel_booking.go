package amadeusHotelModels

import "microservices-travel-backend/internal/hotel-booking/app/dto/request"

type AmadeusHotelOrderRequest struct {
	Data request.HotelOrderRequest `json:"data"`
}

type AmadeusHotelOrderResponse struct {
	Data     AmadeusBookingData `json:"data"`
	Warnings []Warning          `json:"warnings,omitempty"`
}

type HotelOrderData struct {
	Type             string            `json:"type"`
	Guests           []Guest           `json:"guests"`
	TravelAgent      TravelAgent       `json:"travelAgent"`
	RoomAssociations []RoomAssociation `json:"roomAssociations"`
	Payment          PaymentDetails    `json:"payment"`
}

type TravelAgent struct {
	Contact Contact `json:"contact"`
}

type Contact struct {
	Email string `json:"email"`
}

type AmadeusBookingData struct {
	Type              string             `json:"type"`
	ID                string             `json:"id"`
	HotelBookings     []HotelBooking     `json:"hotelBookings"`
	Guests            []Guest            `json:"guests"`
	AssociatedRecords []AssociatedRecord `json:"associatedRecords"`
	Self              string             `json:"self"`
}

type HotelBooking struct {
	Type                     string              `json:"type"`
	ID                       string              `json:"id"`
	BookingStatus            string              `json:"bookingStatus"`
	HotelProviderInformation []HotelProviderInfo `json:"hotelProviderInformation"`
	RoomAssociations         []RoomAssociation   `json:"roomAssociations"`
	HotelOffer               BookedHotelOffer    `json:"hotelOffer"`
	Hotel                    HotelDetails        `json:"hotel"`
	Payment                  PaymentDetails      `json:"payment"`
	TravelAgentID            string              `json:"travelAgentId"`
}

type HotelProviderInfo struct {
	HotelProviderCode  string `json:"hotelProviderCode"`
	ConfirmationNumber string `json:"confirmationNumber"`
}

type Tax struct {
	Amount           string `json:"amount"`
	Code             string `json:"code"`
	Currency         string `json:"currency"`
	Included         bool   `json:"included"`
	PricingFrequency string `json:"pricingFrequency"`
	PricingMode      string `json:"pricingMode"`
}

type Variations struct {
	Changes []PriceChange `json:"changes"`
}

type PriceChange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Base      string `json:"base"`
	Currency  string `json:"currency"`
}

type RoomDetails struct {
	Description Description `json:"description"`
	Type        string      `json:"type"`
}

type HotelDetails struct {
	HotelID   string `json:"hotelId"`
	ChainCode string `json:"chainCode"`
	Name      string `json:"name"`
	Self      string `json:"self"`
}

type AssociatedRecord struct {
	Reference        string `json:"reference"`
	OriginSystemCode string `json:"originSystemCode"`
}

type Warning struct {
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Details string `json:"details,omitempty"`
}
