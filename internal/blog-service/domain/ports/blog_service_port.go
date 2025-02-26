package ports

import "microservices-travel-backend/internal/blog-service/domain/models"

type BlogServicePort interface {
	CreateBlog(blog models.Blog) (*models.Blog, error)
	GetBlogByID(id string) (*models.Blog, error)
	GetAllBlogs() ([]models.Blog, error)
	GetBlogsByAuthor(authorID string) ([]models.Blog, error)
	UpdateBlog(id string, blog models.Blog) (*models.Blog, error)
	DeleteBlog(id string) error
}
