package ports

import (
	"context"
	"microservices-travel-backend/internal/tour-and-packages/domain/models"
)

type TourServicePort interface {
	CreateTour(ctx context.Context, tour models.Tour) (models.Tour, error)
	GetTourByID(ctx context.Context, id string) (models.Tour, error)
	GetAllTours(ctx context.Context) ([]models.Tour, error)
	UpdateTour(ctx context.Context, id string, updates map[string]interface{}) (models.Tour, error)
	DeleteTour(ctx context.Context, id string) error
}
