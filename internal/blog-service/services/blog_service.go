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

func (s *BlogService) CreateBlog(ctx context.Context, blogRequest models.CreateBlogRequest) (*models.Blog, error) {
	createdBlog, err := s.blogRepo.Create(ctx, &blogRequest)
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

func (s *BlogService) GetBlogByCategory(ctx context.Context, category string) ([]*models.Blog, error) {
	blog, err := s.blogRepo.GetByCategory(ctx, category)
	if err != nil {
		return nil, errors.New("blog not found")
	}
	return blog, nil
}

func (s *BlogService) GetAllBlogs(ctx context.Context) ([]models.Blog, error) {
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

func (s *BlogService) UpdateBlog(ctx context.Context, blogID string, updatedDetails models.UpdateBlogRequest) (*models.Blog, error) {
	updates := map[string]interface{}{
		"updated_at": time.Now(),
	}

	if updatedDetails.Title != nil {
		updates["title"] = *updatedDetails.Title
		updates["slug"] = utils.GenerateUniqueSlug(ctx, *updatedDetails.Title)
	}
	if updatedDetails.Content != nil {
		updates["content"] = *updatedDetails.Content
	}
	if updatedDetails.Excerpt != nil {
		updates["excerpt"] = *updatedDetails.Excerpt
	}
	if updatedDetails.MetaTitle != nil {
		updates["meta_title"] = *updatedDetails.MetaTitle
	}
	if updatedDetails.MetaDescription != nil {
		updates["meta_description"] = *updatedDetails.MetaDescription
	}
	if updatedDetails.CategoryID != nil {
		updates["category_id"] = *updatedDetails.CategoryID
	}
	if updatedDetails.Tags != nil {
		updates["tags"] = updatedDetails.Tags
	}
	if updatedDetails.Thumbnail != nil {
		updates["thumbnail"] = *updatedDetails.Thumbnail
	}
	if updatedDetails.Status != nil {
		updates["status"] = *updatedDetails.Status
	}
	if updatedDetails.PublishedAt != nil {
		updates["published_at"] = *updatedDetails.PublishedAt
	}
	if updatedDetails.ScheduledAt != nil {
		updates["scheduled_at"] = *updatedDetails.ScheduledAt
	}

	updatedBlog, err := s.blogRepo.Update(ctx, blogID, updates)
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

func (s *BlogService) UpdateBlogStatus(ctx context.Context, blogID string, status string) (*models.Blog, error) {
	updates := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}

	updatedBlog, err := s.blogRepo.Update(ctx, blogID, updates)
	if err != nil {
		return nil, errors.New("failed to update blog status")
	}

	return updatedBlog, nil
}
