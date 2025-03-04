package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"microservices-travel-backend/internal/blog-service/domain/models"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQLBlogCategoryRepository struct {
	db *gorm.DB
}

func NewPostgreSQLBlogCategoryRepository() (*PostgreSQLBlogCategoryRepository, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	sslMode := os.Getenv("DATABASE_SSLMODE")

	var dsn string
	if databaseName == "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/?sslmode=%s",
			databaseUsername, databasePassword, databaseURL, databasePort, sslMode)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			databaseUsername, databasePassword, databaseURL, databasePort, databaseName, sslMode)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database")

	return &PostgreSQLBlogCategoryRepository{db: db}, nil
}

func (r *PostgreSQLBlogCategoryRepository) Create(ctx context.Context, category *models.BlogCategory) (*models.BlogCategory, error) {
	if category.ID == "" {
		category.ID = uuid.New().String()
	}
	if err := r.db.WithContext(ctx).Table("blog_categories").Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *PostgreSQLBlogCategoryRepository) GetByID(ctx context.Context, id string) (*models.BlogCategory, error) {
	var category models.BlogCategory
	if err := r.db.WithContext(ctx).Table("blog_categories").First(&category, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (r *PostgreSQLBlogCategoryRepository) GetAll(ctx context.Context) ([]*models.BlogCategory, error) {
	var categories []*models.BlogCategory
	if err := r.db.WithContext(ctx).Table("blog_categories").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *PostgreSQLBlogCategoryRepository) Update(ctx context.Context, id string, category *models.BlogCategory) (*models.BlogCategory, error) {
	if err := r.db.WithContext(ctx).Table("blog_categories").Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *PostgreSQLBlogCategoryRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Table("blog_categories").Delete(&models.BlogCategory{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgreSQLBlogCategoryRepository) NameExists(ctx context.Context, name string) bool {
	var count int64
	err := r.db.WithContext(ctx).Table("blog_categories").Where("name = ?", name).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}
