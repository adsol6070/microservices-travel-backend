package services

import (
	"context"
	"errors"
	"microservices-travel-backend/internal/blog-service/domain/models"
	"microservices-travel-backend/internal/blog-service/domain/ports"
	"microservices-travel-backend/pkg/utils"
	"time"
)

type BlogService struct {
	blogRepo ports.BlogRepositoryPort
}

func NewBlogService(blogRepo ports.BlogRepositoryPort) *BlogService {
	return &BlogService{blogRepo: blogRepo}
}

func (s *BlogService) CreateBlog(ctx context.Context, blogDetails *models.Blog) (*models.Blog, error) {
	blogDetails.Slug = utils.GenerateUniqueSlug(ctx, blogDetails.Title, s.blogRepo.SlugExists)
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

func (s *BlogService) GetBlogByCategory(ctx context.Context, category string) (*models.Blog, error) {
	blogCategoryTitle := utils.GenerateSlugToTitle(category)
	blog, err := s.blogRepo.GetByID(ctx, blogCategoryTitle)
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
	existingBlog.Slug = utils.GenerateUniqueSlug(ctx, updatedDetails.Title, s.blogRepo.SlugExists)

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
