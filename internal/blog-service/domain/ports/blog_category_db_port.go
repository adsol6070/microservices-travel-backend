package ports

import (
	"context"
	"microservices-travel-backend/internal/blog-service/domain/models"
)

type BlogCategoryRepositoryPort interface {
	Create(ctx context.Context, category *models.BlogCategory) (*models.BlogCategory, error)
	GetByID(ctx context.Context, id string) (*models.BlogCategory, error)
	GetAll(ctx context.Context) ([]*models.BlogCategory, error)
	Update(ctx context.Context, id string, category *models.BlogCategory) (*models.BlogCategory, error)
	Delete(ctx context.Context, id string) error
	CategoryExists(ctx context.Context, name string) bool
}
