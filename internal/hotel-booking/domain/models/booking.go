package models

import (
	"time"
)

// BookingStatus Enum
type BookingStatus string

const (
	StatusPending                 BookingStatus = "pending"
	StatusConfirmed               BookingStatus = "confirmed"
	StatusAwaitingPayment         BookingStatus = "awaiting_payment"
	StatusPaymentReceived         BookingStatus = "payment_received"
	StatusCancelled               BookingStatus = "cancelled"
	StatusCancelledByGuest        BookingStatus = "cancelled_by_guest"
	StatusCancelledByHotel        BookingStatus = "cancelled_by_hotel"
	StatusCheckedIn               BookingStatus = "checked_in"
	StatusCheckedOut              BookingStatus = "checked_out"
	StatusNoShow                  BookingStatus = "no_show"
	StatusAwaitingConfirmation    BookingStatus = "awaiting_confirmation"
	StatusPendingReview           BookingStatus = "pending_review"
	StatusReviewed                BookingStatus = "reviewed"
	StatusExpired                 BookingStatus = "expired"
	StatusInProgress              BookingStatus = "in_progress"
	StatusUnderPaymentReview      BookingStatus = "under_payment_review"
	StatusUnderCancellationReview BookingStatus = "under_cancellation_review"
	StatusRefunded                BookingStatus = "refunded"
	StatusDisputed                BookingStatus = "disputed"
)

// Booking Model
type Booking struct {
	ID                      string        `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	HotelID                 string        `json:"hotel_id"`
	RoomID                  *string       `json:"room_id"`
	StartDate               time.Time     `json:"start_date"`
	EndDate                 time.Time     `json:"end_date"`
	CheckinDate             *time.Time    `json:"checkin_date"`
	CheckoutDate            *time.Time    `json:"checkout_date"`
	Status                  BookingStatus `gorm:"type:booking_status" json:"status"`
	CreatedAt               time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt               time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	TotalPrice              float64       `json:"total_price"`
	GuestCount              int           `json:"guest_count"`
	SpecialRequests         string        `json:"special_requests"`
	BranchID                *int          `json:"branch_id"`
	CancellationReason      *string       `json:"cancellation_reason"`
	CancellationDate        *time.Time    `json:"cancellation_date"`
	ExternalBookingID       *string       `json:"external_booking_id"`
	ExternalSyncAttempts    int           `json:"external_sync_attempts"`
	ExternalSyncLastAttempt *time.Time    `json:"external_sync_last_attempt"`
	ExternalSyncError       *string       `json:"external_sync_error"`
}
