package models

import (
	"time"

	"github.com/google/uuid"
)

type Tour struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
