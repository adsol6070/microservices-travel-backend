package amadeusHotelModels

import (
	"encoding/json"
	"strconv"
)

import "time"

type Hotel struct {
	Type       string          `json:"type,omitempty"`
	HotelID    string          `json:"hotelId"`
	ChainCode  string          `json:"chainCode"`
	DupeID     json.RawMessage `json:"dupeId"`
	Name       string          `json:"name"`
	CityCode   *string         `json:"cityCode,omitempty"`
	IATACode   *string         `json:"iataCode,omitempty"`
	GeoCode    *GeoCode        `json:"geoCode,omitempty"`
	Latitude   *float64        `json:"latitude,omitempty"`
	Longitude  *float64        `json:"longitude,omitempty"`
	Address    *Address        `json:"address,omitempty"`
	Distance   *Distance       `json:"distance,omityempty"`
	LastUpdate string          `json:"lastUpdate,omitempty"`
}

type CustomTime struct {
	time.Time
}

func (h *Hotel) GetDupeID() string {
	var dupeIDStr string
	if err := json.Unmarshal(h.DupeID, &dupeIDStr); err != nil {
		return dupeIDStr
	}

	var dupeIDNum int
	if err := json.Unmarshal(h.DupeID, &dupeIDNum); err != nil {
		return strconv.Itoa(dupeIDNum)
	}

	return ""
}

type GeoCode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Address struct {
	CountryCode string `json:"countryCode"`
}

type Distance struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type Guests struct {
	Adults int `json:"adults"`
}

type Price struct {
	Base         string     `json:"base"`
	Currency     string     `json:"currency"`
	SellingTotal *string    `json:"sellingTotal"`
	Taxes        *[]Tax     `json:"taxes"`
	Total        string     `json:"total"`
	Variations   Variations `json:"variations"`
}

type Description struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type GuestReference struct {
	GuestReference string `json:"guestReference"`
}

type RoomAssociation struct {
	HotelOfferID    string           `json:"hotelOfferId"`
	GuestReferences []GuestReference `json:"guestReferences"`
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
	ID        *int   `json:"id"`
	Title     string `json:"title"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type BaseHotelOffer struct {
	ID       string   `json:"id"`
	CheckIn  string   `json:"checkInDate"`
	CheckOut string   `json:"checkOutDate"`
	RateCode string   `json:"rateCode"`
	Guests   Guests   `json:"guests"`
	Policies Policies `json:"policies"`
	Price    Price    `json:"price"`
	Room     Room     `json:"room"`
}

type Offer struct {
	BaseHotelOffer
	RateFamilyEstimated RateFamilyEstimated `json:"rateFamilyEstimated"`
	Self                string              `json:"self"`
}

type BookedHotelOffer struct {
	BaseHotelOffer
	Type      string      `json:"type"`
	Category  string      `json:"category"`
	RoomCount int         `json:"roomQuantity"`
	Room      RoomDetails `json:"room"`
}

type Policies struct {
	AdditionalDetails *[]AdditionalDetail `json:"additionalDetails"`
	Cancellations     []Cancellation      `json:"cancellations"`
	PaymentType       string              `json:"paymentType"`
	Refundable        Refundable          `json:"refundable"`
}

type AdditionalDetail struct {
	Description []Description `json:"description"`
}

type Cancellation struct {
	Amount         string `json:"amount"`
	Deadline       string `json:"deadline"`
	PolicyType     string `json:"policyType"`
	NumberOfNights *int   `json:"numberOfNights"`
}

type Refundable struct {
	CancellationRefund string `json:"cancellationRefund"`
}
