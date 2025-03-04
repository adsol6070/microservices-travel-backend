package services

import (
	"context"
	"errors"
	"microservices-travel-backend/internal/blog-service/domain/models"
	"microservices-travel-backend/internal/blog-service/domain/ports"
	"time"
)

type BlogCategoryService struct {
	categoryRepo ports.BlogCategoryRepositoryPort
}

func NewBlogCategoryService(categoryRepo ports.BlogCategoryRepositoryPort) *BlogCategoryService {
	return &BlogCategoryService{categoryRepo: categoryRepo}
}

func (s *BlogCategoryService) CreateCategory(ctx context.Context, categoryDetails *models.BlogCategory) (*models.BlogCategory, error) {
	categoryDetails.CreatedAt = time.Now()
	categoryDetails.UpdatedAt = categoryDetails.CreatedAt

	createdCategory, err := s.categoryRepo.Create(ctx, categoryDetails)
	if err != nil {
		return nil, errors.New("failed to create category")
	}

	return createdCategory, nil
}

func (s *BlogCategoryService) GetCategoryByID(ctx context.Context, categoryID string) (*models.BlogCategory, error) {
	category, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}

func (s *BlogCategoryService) GetAllCategories(ctx context.Context) ([]*models.BlogCategory, error) {
	categories, err := s.categoryRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve categories")
	}
	return categories, nil
}

func (s *BlogCategoryService) UpdateCategory(ctx context.Context, categoryID string, updatedDetails *models.BlogCategory) (*models.BlogCategory, error) {
	existingCategory, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	existingCategory.UpdatedAt = time.Now()

	updatedCategory, err := s.categoryRepo.Update(ctx, categoryID, existingCategory)
	if err != nil {
		return nil, errors.New("failed to update category")
	}

	return updatedCategory, nil
}

func (s *BlogCategoryService) DeleteCategory(ctx context.Context, categoryID string) error {
	_, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return errors.New("category not found")
	}

	if err := s.categoryRepo.Delete(ctx, categoryID); err != nil {
		return errors.New("failed to delete category")
	}

	return nil
}
