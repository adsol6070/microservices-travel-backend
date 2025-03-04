package ports

import (
	"context"
	"microservices-travel-backend/internal/blog-service/domain/models"
)

type BlogCategoryServicePort interface {
	CreateCategory(ctx context.Context, category *models.BlogCategory) (*models.BlogCategory, error)
	GetCategoryByID(ctx context.Context, id string) (*models.BlogCategory, error)
	GetAllCategories(ctx context.Context) ([]*models.BlogCategory, error)
	UpdateCategory(ctx context.Context, id string, category *models.BlogCategory) (*models.BlogCategory, error)
	DeleteCategory(ctx context.Context, id string) error
}
