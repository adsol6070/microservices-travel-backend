package response

type HotelBookReceipt struct {
	Type               string          `json:"type"`
	ID                 string          `json:"id"`
	HotelBookings      []HotelBooking  `json:"hotelBookings"`
	Guests             []Guest         `json:"guests"`
	AssociatedRecords  []Record        `json:"associatedRecords"`
	Self               string          `json:"self"`
	Warnings           []Warning       `json:"warnings"`
}

type HotelBooking struct {
	Type                     string                   `json:"type"`
	ID                       string                   `json:"id"`
	BookingStatus            string                   `json:"bookingStatus"`
	HotelProviderInformation []HotelProviderInfo      `json:"hotelProviderInformation"`
	RoomAssociations         []RoomAssociation        `json:"roomAssociations"`
	HotelOffer               HotelOffer               `json:"hotelOffer"`
	Hotel                    HotelInfo                `json:"hotel"`
	Payment                  PaymentDetails           `json:"payment"`
	TravelAgentID            string                   `json:"travelAgentId"`
}

type HotelProviderInfo struct {
	HotelProviderCode   string `json:"hotelProviderCode"`
	ConfirmationNumber  string `json:"confirmationNumber"`
}

type RoomAssociation struct {
	HotelOfferID    string           `json:"hotelOfferId"`
	GuestReferences []GuestReference `json:"guestReferences"`
}

type GuestReference struct {
	GuestReference string `json:"guestReference"`
}

type HotelOffer struct {
	ID         string      `json:"id"`
	Type       string      `json:"type"`
	Category   string      `json:"category"`
	CheckIn    string      `json:"checkInDate"`
	CheckOut   string      `json:"checkOutDate"`
	Guests     GuestInfo   `json:"guests"`
	Policies   Policies    `json:"policies"`
	Price      Price       `json:"price"`
	RateCode   string      `json:"rateCode"`
	Room       RoomDetails `json:"room"`
	RoomQty    int         `json:"roomQuantity"`
}

type GuestInfo struct {
	Adults int `json:"adults"`
}

type Policies struct {
	Cancellations []Cancellation `json:"cancellations"`
	PaymentType   string         `json:"paymentType"`
	Refundable    RefundInfo     `json:"refundable"`
}

type Cancellation struct {
	Amount     string `json:"amount"`
	Deadline   string `json:"deadline"`
	PolicyType string `json:"policyType"`
}

type RefundInfo struct {
	CancellationRefund string `json:"cancellationRefund"`
}

type Price struct {
	Base        string   `json:"base"`
	Currency    string   `json:"currency"`
	SellingTotal string  `json:"sellingTotal"`
	Taxes       []Tax    `json:"taxes"`
	Total       string   `json:"total"`
	Variations  Variation `json:"variations"`
}

type Tax struct {
	Amount          string `json:"amount"`
	Code            string `json:"code"`
	Currency        string `json:"currency"`
	Included        bool   `json:"included"`
	PricingFrequency string `json:"pricingFrequency"`
	PricingMode      string `json:"pricingMode"`
}

type Variation struct {
	Changes []Change `json:"changes"`
}

type Change struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Base      string `json:"base"`
}

type RoomDetails struct {
	Description Description `json:"description"`
	Type        string      `json:"type"`
}

type Description struct {
	Lang string `json:"lang"`
	Text string `json:"text"`
}

type HotelInfo struct {
	HotelID   string `json:"hotelId"`
	ChainCode string `json:"chainCode"`
	Name      string `json:"name"`
	Self      string `json:"self"`
}

type PaymentDetails struct {
	Method      string       `json:"method"`
	PaymentCard PaymentCard  `json:"paymentCard"`
}

type PaymentCard struct {
	PaymentCardInfo CardInfo `json:"paymentCardInfo"`
}

type CardInfo struct {
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

type Record struct {
	Reference       string `json:"reference"`
	OriginSystemCode string `json:"originSystemCode"`
}

type Warning struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
}
