package repositories

import (
	"context"
	"microservices-travel-backend/internal/tour-and-packages/domain/models"

	"gorm.io/gorm"
)

type PackageRepository struct {
	db *gorm.DB
}

func NewPackageRepository(db *gorm.DB) *PackageRepository {
	return &PackageRepository{db: db}
}

func (r *PackageRepository) Create(ctx context.Context, pkg models.Package) (models.Package, error) {

}

func (r *PackageRepository) GetByID(ctx context.Context, id string) (models.Package, error) {

}

func (r *PackageRepository) GetAll(ctx context.Context) ([]models.Package, error) {

}

func (r *PackageRepository) Update(ctx context.Context, id string, updates map[string]interface{}) (models.Package, error) {

}

func (r *PackageRepository) Delete(ctx context.Context, id string) error {

}
