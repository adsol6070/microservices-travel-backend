package ports

import (
	"context"
	"microservices-travel-backend/internal/tour-and-packages/domain/models"
)

type PackageServicePort interface {
	CreatePackage(ctx context.Context, pkg models.Package) (models.Package, error)
	GetPackageByID(ctx context.Context, id string) (models.Package, error)
	GetAllPackages(ctx context.Context) ([]models.Package, error)
	UpdatePackage(ctx context.Context, id string, updates map[string]interface{}) (models.Package, error)
	DeletePackage(ctx context.Context, id string) error
}
