package ports

import (
	"context"
	"microservices-travel-backend/internal/tour-and-packages/domain/models"
)

type PackageRepositoryPort interface {
	Create(ctx context.Context, pkg models.Package) (models.Package, error)
	GetByID(ctx context.Context, id string) (models.Package, error)
	GetAll(ctx context.Context) ([]models.Package, error)
	Update(ctx context.Context, id string, updates map[string]interface{}) (models.Package, error)
	Delete(ctx context.Context, id string) error
}
