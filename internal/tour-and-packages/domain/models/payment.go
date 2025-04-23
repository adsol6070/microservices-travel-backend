package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID uuid.UUID `json:"id"`
	BookingID uuid.UUID `json:"booking_id"`
	Amount float64 `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	TransactionID string `json:"transaction_id"`
	Status string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}