package repositories

import (
	"context"
	"microservices-travel-backend/internal/tour-and-packages/domain/models"

	"gorm.io/gorm"
)

type TourRepository struct {
	db *gorm.DB
}

func NewTourRepository(db *gorm.DB) *TourRepository {
	return &TourRepository{db: db}
}

func (r *TourRepository) Create(ctx context.Context, tour models.Tour) (models.Tour, error) {

}

func (r *TourRepository) GetByID(ctx context.Context, id string) (models.Tour, error) {

}

func (r *TourRepository) GetAll(ctx context.Context) ([]models.Tour, error) {

}

func (r *TourRepository) Update(ctx context.Context, id string, updates map[string]interface{}) (models.Tour, error) {

}

func (r *TourRepository) Delete(ctx context.Context, id string) error {

}