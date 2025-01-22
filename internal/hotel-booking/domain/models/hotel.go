package models
import "time"

type Asset struct {
	Type   string `json:"type"`
	S3Link string `json:"s3Link"`
}

// Hotel represents a hotel booking
type Hotel struct {
	HotelID           string    `json:"hotelID"`
	BookingID         string    `json:"bookingID"`
	UserID            string    `json:"userID"`
	CheckInDate       time.Time `json:"checkInDate"`
	CheckOutDate      time.Time `json:"checkOutDate"`
	IsCheckinFlexible bool      `json:"isCheckinFlexible"`
	BookingStatus     string    `json:"bookingStatus"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	PaymentStatus     string    `json:"paymentStatus"`
	RoomType          string    `json:"roomType"`
	NumberOfGuests    int       `json:"numberOfGuests"`
	SpecialRequests   string    `json:"specialRequests"`
	Assets            []Asset   `json:"assets"`
}
