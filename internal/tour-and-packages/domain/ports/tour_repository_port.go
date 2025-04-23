package ports

import (
	"context"
	"microservices-travel-backend/internal/tour-and-packages/domain/models"
)

type TourRepositoryPort interface {
	Create(ctx context.Context, tour models.Tour) (models.Tour, error)
	GetByID(ctx context.Context, id string) (models.Tour, error)
	GetAll(ctx context.Context) ([]models.Tour, error)
	Update(ctx context.Context, id string, updates map[string]interface{}) (models.Tour, error)
	Delete(ctx context.Context, id string) error
}
