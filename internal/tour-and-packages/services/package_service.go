package services

import (
	"context"
	"microservices-travel-backend/internal/tour-and-packages/domain/models"
	"microservices-travel-backend/internal/tour-and-packages/domain/ports"
)

type PackageService struct {
	packageRepository ports.PackageRepositoryPort
}

func NewPackageService(repository ports.PackageRepositoryPort) *PackageService {
	return &PackageService{
		packageRepository: repository,
	}
}

func (s *PackageService) CreatePackage(ctx context.Context, pkg models.Package) (models.Package, error) {
	return models.Package{}, nil
}

func (s *PackageService) GetPackageByID(ctx context.Context, id string) (models.Package, error) {
	return models.Package{}, nil
}

func (s *PackageService) GetAllPackages(ctx context.Context) ([]models.Package, error) {
	return []models.Package{}, nil
}

func (s *PackageService) UpdatePackage(ctx context.Context, id string, updates map[string]interface{}) (models.Package, error) {
	return models.Package{}, nil
}

func (s *PackageService) DeletePackage(ctx context.Context, id string) error {
	return nil
}
