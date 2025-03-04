package services

import (
	"context"
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

func (s *BlogService) CreateBlog(ctx context.Context, blogDetails *models.Blog) (*models.Blog, error) {
	if err := s.validateBlogDetails(blogDetails); err != nil {
		return nil, err
	}

	blogDetails.CreatedAt = time.Now()
	blogDetails.UpdatedAt = blogDetails.CreatedAt

	createdBlog, err := s.blogRepo.Create(ctx, blogDetails)
	if err != nil {
		return nil, errors.New("failed to create blog")
	}

	return createdBlog, nil
}

func (s *BlogService) GetBlogByID(ctx context.Context, blogID string) (*models.Blog, error) {
	blog, err := s.blogRepo.GetByID(ctx, blogID)
	if err != nil {
		return nil, errors.New("blog not found")
	}
	return blog, nil
}

func (s *BlogService) GetAllBlogs(ctx context.Context) ([]*models.Blog, error) {
	blogs, err := s.blogRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve blogs")
	}
	return blogs, nil
}

func (s *BlogService) GetBlogsByAuthor(ctx context.Context, authorID string) ([]*models.Blog, error) {
	blogs, err := s.blogRepo.GetByAuthor(ctx, authorID)
	if err != nil {
		return nil, errors.New("failed to retrieve blogs by author")
	}
	return blogs, nil
}

func (s *BlogService) UpdateBlog(ctx context.Context, blogID string, updatedDetails *models.Blog) (*models.Blog, error) {
	existingBlog, err := s.blogRepo.GetByID(ctx, blogID)
	if err != nil {
		return nil, errors.New("blog not found")
	}

	if updatedDetails.Title != "" {
		existingBlog.Title = updatedDetails.Title
	}
	if updatedDetails.Content != "" {
		existingBlog.Content = updatedDetails.Content
	}

	existingBlog.UpdatedAt = time.Now()

	updatedBlog, err := s.blogRepo.Update(ctx, blogID, existingBlog)
	if err != nil {
		return nil, errors.New("failed to update blog")
	}

	return updatedBlog, nil
}

func (s *BlogService) DeleteBlog(ctx context.Context, blogID string) error {
	_, err := s.blogRepo.GetByID(ctx, blogID)
	if err != nil {
		return errors.New("blog not found")
	}

	if err := s.blogRepo.Delete(ctx, blogID); err != nil {
		return errors.New("failed to delete blog")
	}

	return nil
}

func (s *BlogService) validateBlogDetails(blog *models.Blog) error {
	if blog.Title == "" {
		return errors.New("title is required")
	}
	if blog.Content == "" {
		return errors.New("content is required")
	}
	return nil
}
