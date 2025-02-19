package request

type HotelSearchRequest struct {
	CityCode     string `json:"cityCode" validate:"required,len=3"` // City code (e.g., "NYC" for New York)
	CheckInDate  string `json:"checkInDate" validate:"required"`    // Check-in date in YYYY-MM-DD format
	CheckOutDate string `json:"checkOutDate" validate:"required"`   // Check-out date in YYYY-MM-DD format
	Rooms        int    `json:"rooms" validate:"required,min=1"`    // Number of rooms requested
	Persons      int    `json:"persons" validate:"required,min=1"`  // Number of persons for the booking
}
