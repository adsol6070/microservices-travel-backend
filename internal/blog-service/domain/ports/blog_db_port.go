package ports

import (
	"context"
	"microservices-travel-backend/internal/blog-service/domain/models"
)

type BlogRepositoryPort interface {
	Create(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	GetByID(ctx context.Context, id string) (*models.Blog, error)
	GetAll(ctx context.Context) ([]*models.Blog, error)
	GetByAuthor(ctx context.Context, authorID string) ([]*models.Blog, error)
	Update(ctx context.Context, id string, blog *models.Blog) (*models.Blog, error)
	Delete(ctx context.Context, id string) error
}
