package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID uuid.UUID `json:"booking_id"`
	UserID uuid.UUID `json:"user_id"`
	PackageID uuid.UUID `json:"package_id"`
	Package Package `json:"package"`
	BookingDate time.Time `json:"booking_date"`
	TotalAmount float64 `json:"total_amount"`
	Status string `json:"status"`
	PaymentID uuid.UUID `json:"payment_id"`
	Payment Payment `json:"payment"`
	CreatedAt time.Time
	UpdatedAt time.Time
}