package models

import "time"

type Booking struct {
	BookingID     string    `json:"bookingID"`
	UserID        string    `json:"userID"`
	FlightID      string    `json:"flightID"`
	BookingStatus string    `json:"bookingStatus"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
