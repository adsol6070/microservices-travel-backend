package ports

import (
	"context"
	"microservices-travel-backend/internal/blog-service/domain/models"
)

type BlogServicePort interface {
	CreateBlog(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	GetBlogByID(ctx context.Context, id string) (*models.Blog, error)
	GetAllBlogs(ctx context.Context) ([]*models.Blog, error)
	GetBlogsByAuthor(ctx context.Context, authorID string) ([]*models.Blog, error)
	UpdateBlog(ctx context.Context, id string, blog *models.Blog) (*models.Blog, error)
	DeleteBlog(ctx context.Context, id string) error
}
