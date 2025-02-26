package services

import (
	"errors"
	"microservices-travel-backend/internal/blog-service/domain/models"
	"microservices-travel-backend/internal/blog-service/domain/ports"
	"time"
)

type BlogService struct {
	blogRepo ports.BlogRepositoryPort
}

func NewBlogService(blogRepo ports.BlogRepositoryPort) *BlogService {
	return &BlogService{blogRepo: blogRepo}
}

// CreateBlog creates a new blog post
func (s *BlogService) CreateBlog(blog models.Blog) (*models.Blog, error) {
	if blog.Title == "" || blog.Content == "" {
		return nil, errors.New("title and content cannot be empty")
	}

	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	createdBlog, err := s.blogRepo.Create(blog)
	if err != nil {
		return nil, err
	}

	return createdBlog, nil
}

// GetBlogByID retrieves a blog post by its ID
func (s *BlogService) GetBlogByID(id string) (*models.Blog, error) {
	blog, err := s.blogRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// GetAllBlogs retrieves all blog posts
func (s *BlogService) GetAllBlogs() ([]models.Blog, error) {
	blogs, err := s.blogRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

// GetBlogsByAuthor retrieves all blog posts written by a specific author
func (s *BlogService) GetBlogsByAuthor(authorID string) ([]models.Blog, error) {
	blogs, err := s.blogRepo.GetByAuthor(authorID)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

// UpdateBlog updates a blog post
func (s *BlogService) UpdateBlog(id string, updatedBlog models.Blog) (*models.Blog, error) {
	existingBlog, err := s.blogRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("blog not found")
	}

	existingBlog.Title = updatedBlog.Title
	existingBlog.Content = updatedBlog.Content
	existingBlog.UpdatedAt = time.Now()

	updatedBlogData, err := s.blogRepo.Update(id, *existingBlog)
	if err != nil {
		return nil, err
	}

	return updatedBlogData, nil
}

// DeleteBlog deletes a blog post by its ID
func (s *BlogService) DeleteBlog(id string) error {
	return s.blogRepo.Delete(id)
}