package models

import (
	"time"

	"github.com/google/uuid"
)

type Package struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	TourID uuid.UUID `json:"tour_id"`
	Tour Tour `json:"tour"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Duration int `json:"duration"`
	IsAvailable bool `json:"is_available"`
	ImageURL string `json:"image_url"`
	CreatedAt time.Time
	UpdatedAt time.Time
}