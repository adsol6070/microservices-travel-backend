package ports

import "microservices-travel-backend/internal/blog-service/domain/models"

type BlogRepositoryPort interface {
	Create(blog models.Blog) (*models.Blog, error)
	GetByID(id string) (*models.Blog, error)
	GetAll() ([]models.Blog, error)
	GetByAuthor(authorID string) ([]models.Blog, error)
	Update(id string, blog models.Blog) (*models.Blog, error)
	Delete(id string) error
}
