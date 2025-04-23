package services

import (
	"context"
	"microservices-travel-backend/internal/tour-and-packages/domain/models"
	"microservices-travel-backend/internal/tour-and-packages/domain/ports"
)

type TourService struct {
	tourRepository ports.TourRepositoryPort
}

func NewTourService(repository ports.TourRepositoryPort) *TourService {
	return &TourService{
		tourRepository: repository,
	}
}

func (s *TourService) CreateTour(ctx context.Context, tour models.Tour) (models.Tour, error) {
	return models.Tour{}, nil
}

func (s *TourService) GetTourByID(ctx context.Context, id string) (models.Tour, error) {
	return models.Tour{}, nil
}

func (s *TourService) GetAllTours(ctx context.Context) ([]models.Tour, error) {
	return []models.Tour{}, nil
}

func (s *TourService) UpdateTour(ctx context.Context, id string, updates map[string]interface{}) (models.Tour, error) {
	return models.Tour{}, nil
}

func (s *TourService) DeleteTour(ctx context.Context, id string) error {
	return nil
}
