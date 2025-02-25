package googlePlaceModels

type PlaceDetailsResponse struct {
	Name                     string               `json:"name"`
	ID                       string               `json:"id"`
	Types                    []string             `json:"types"`
	NationalPhoneNumber      string               `json:"nationalPhoneNumber"`
	InternationalPhoneNumber string               `json:"internationalPhoneNumber"`
	FormattedAddress         string               `json:"formattedAddress"`
	AddressComponents        []AddressComponent   `json:"addressComponents"`
	PlusCode                 PlusCode             `json:"plusCode"`
	Location                 Coordinates          `json:"location"`
	Viewport                 Viewport             `json:"viewport"`
	Rating                   float64              `json:"rating"`
	GoogleMapsURI            string               `json:"googleMapsUri"`
	WebsiteURI               string               `json:"websiteUri"`
	UTCOffsetMinutes         int                  `json:"utcOffsetMinutes"`
	ADRFormatAddress         string               `json:"adrFormatAddress"`
	BusinessStatus           string               `json:"businessStatus"`
	UserRatingCount          int                  `json:"userRatingCount"`
	IconMaskBaseURI          string               `json:"iconMaskBaseUri"`
	IconBackgroundColor      string               `json:"iconBackgroundColor"`
	DisplayName              DisplayName          `json:"displayName"`
	PrimaryTypeDisplayName   DisplayName          `json:"primaryTypeDisplayName"`
	PrimaryType              string               `json:"primaryType"`
	ShortFormattedAddress    string               `json:"shortFormattedAddress"`
	EditorialSummary         EditorialSummary     `json:"editorialSummary"`
	Reviews                  []Review             `json:"reviews"`
	Photos                   []PhotoDetail        `json:"photos"`
	GoodForChildren          bool                 `json:"goodForChildren"`
	PaymentOptions           PaymentOptions       `json:"paymentOptions"`
	AccessibilityOptions     AccessibilityOptions `json:"accessibilityOptions"`
	PureServiceAreaBusiness  bool                 `json:"pureServiceAreaBusiness"`
	AddressDescriptor        AddressDescriptor    `json:"addressDescriptor"`
	GoogleMapsLinks          GoogleMapsLinks      `json:"googleMapsLinks"`
}

type AddressComponent struct {
	LongText     string   `json:"longText"`
	ShortText    string   `json:"shortText"`
	Types        []string `json:"types"`
	LanguageCode string   `json:"languageCode"`
}

type PlusCode struct {
	GlobalCode   string `json:"globalCode"`
	CompoundCode string `json:"compoundCode"`
}

type Viewport struct {
	Low  Coordinates `json:"low"`
	High Coordinates `json:"high"`
}

type EditorialSummary struct {
	Text         string `json:"text"`
	LanguageCode string `json:"languageCode"`
}

type Review struct {
	Name                           string            `json:"name"`
	RelativePublishTimeDescription string            `json:"relativePublishTimeDescription"`
	Rating                         int               `json:"rating"`
	Text                           ReviewText        `json:"text"`
	OriginalText                   ReviewText        `json:"originalText"`
	AuthorAttribution              AuthorAttribution `json:"authorAttribution"`
	PublishTime                    string            `json:"publishTime"`
	FlagContentURI                 string            `json:"flagContentUri"`
	GoogleMapsURI                  string            `json:"googleMapsUri"`
}

type ReviewText struct {
	Text         string `json:"text"`
	LanguageCode string `json:"languageCode"`
}

type PhotoDetail struct {
	Name               string                    `json:"name"`
	WidthPx            int                       `json:"widthPx"`
	HeightPx           int                       `json:"heightPx"`
	AuthorAttributions []AuthorAttributionDetail `json:"authorAttributions"`
	FlagContentURI     string                    `json:"flagContentUri"`
	GoogleMapsURI      string                    `json:"googleMapsUri"`
}

type AuthorAttributionDetail struct {
	DisplayName string `json:"displayName"`
	URI         string `json:"uri"`
	PhotoURI    string `json:"photoUri"`
}

type PaymentOptions struct {
	CardsAccepted     []string `json:"cardsAccepted"`
	CashAccepted      bool     `json:"cashAccepted"`
	AcceptsDebitCards bool     `json:"acceptsDebitCards"`
	AcceptsCashOnly   bool     `json:"acceptsCashOnly"`
	AcceptsNfc        bool     `json:"acceptsNfc"`
	DigitalWallets    []string `json:"digitalWalletsAccepted"`
}

type AccessibilityOptions struct {
	WheelchairAccessibleParking  bool `json:"wheelchairAccessibleParking"`
	WheelchairAccessibleEntrance bool `json:"wheelchairAccessibleEntrance"`
	WheelchairAccessible         bool `json:"wheelchairAccessible"`
}

type AddressDescriptor struct {
	Landmarks  []Landmark `json:"landmarks"`
	Areas      []Area     `json:"areas"`
	Street     string     `json:"street"`
	City       string     `json:"city"`
	State      string     `json:"state"`
	PostalCode string     `json:"postalCode"`
	Country    string     `json:"country"`
}

type Landmark struct {
	Name                 string      `json:"name"`
	PlaceID              string      `json:"placeId"`
	DisplayName          DisplayName `json:"displayName"`
	Types                []string    `json:"types"`
	StraightLineDistance float64     `json:"straightLineDistanceMeters"`
	TravelDistance       float64     `json:"travelDistanceMeters"`
}

type Area struct {
	Name        string      `json:"name"`
	PlaceID     string      `json:"placeId"`
	DisplayName DisplayName `json:"displayName"`
	Containment string      `json:"containment"`
}

type GoogleMapsLinks struct {
	DirectionsURI   string `json:"directionsUri"`
	PlaceURI        string `json:"placeUri"`
	WriteAReviewURI string `json:"writeAReviewUri"`
	ReviewsURI      string `json:"reviewsUri"`
	PhotosURI       string `json:"photosUri"`
}
