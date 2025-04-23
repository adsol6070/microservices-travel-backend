package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID             uuid.UUID            `json:"id" gorm:"type:uuid;primaryKey"`
	UserID         uuid.UUID            `json:"user_id" gorm:"type:uuid;not null" validate:"required,uuid"`
	BookingID      uuid.UUID            `json:"booking_id" gorm:"type:uuid;not null" validate:"required,uuid"`
	InvoiceNumber  string               `json:"invoice_number" gorm:"uniqueIndex;not null" validate:"required"`
	Currency       string               `json:"currency" validate:"required"`
	Amount         float64              `json:"amount" validate:"required,gt=0"`
	TaxAmount      float64              `json:"tax_amount"`
	DiscountAmount float64              `json:"discount_amount"`
	FinalAmount    float64              `json:"final_amount" validate:"required,gt=0"`
	Status         string               `json:"status" validate:"required,oneof=paid unpaid cancelled refunded"`
	PaymentMethod  string               `json:"payment_method,omitempty"`
	PaidAt         *time.Time           `json:"paid_at,omitempty"`
	Description    []InvoicePackageItem `json:"description" gorm:"foreignKey:InvoiceID;constraint:OnDelete:CASCADE;"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
	DeletedAt      gorm.DeletedAt       `json:"deleted_at,omitempty" gorm:"index"`
}

type InvoicePackageItem struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	InvoiceID   uuid.UUID `json:"invoice_id" gorm:"type:uuid;index;not null"`
	PackageName string    `json:"package_name" validate:"required"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity" validate:"required,min=1"`
	UnitPrice   float64   `json:"unit_price" validate:"required,gt=0"`
	TotalPrice  float64   `json:"total_price" validate:"required,gt=0"`
}

type UpdateInvoiceRequest struct {
	Status        *string    `json:"status,omitempty" validate:"omitempty,oneof=paid unpaid cancelled refunded"`
	PaymentMethod *string    `json:"payment_method,omitempty" validate:"omitempty"`
	PaidAt        *time.Time `json:"paid_at,omitempty"`
}
