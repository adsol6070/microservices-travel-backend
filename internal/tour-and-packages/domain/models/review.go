package models

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	PackageID uuid.UUID `json:"package_id"`
	Package Package `json:"package"`
	Rating int `json:"rating"`
	Comment string `json:"comment"`
	CreatedAt time.Time
	UpdatedAt time.Time
}