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

type PostgreSQLBlogRepository struct {
	db *gorm.DB
}

func NewPostgreSQLBlogRepository() (*PostgreSQLBlogRepository, error) {
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

	return &PostgreSQLBlogRepository{db: db}, nil
}

func (r *PostgreSQLBlogRepository) Create(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	if blog.ID == "" {
		blog.ID = uuid.New().String()
	}
	if err := r.db.WithContext(ctx).Create(blog).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *PostgreSQLBlogRepository) GetByID(ctx context.Context, id string) (*models.Blog, error) {
	var blog models.Blog
	if err := r.db.WithContext(ctx).First(&blog, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &blog, nil
}

func (r *PostgreSQLBlogRepository) GetAll(ctx context.Context) ([]*models.Blog, error) {
	var blogs []*models.Blog
	if err := r.db.WithContext(ctx).Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *PostgreSQLBlogRepository) GetByAuthor(ctx context.Context, authorID string) ([]*models.Blog, error) {
	var blogs []*models.Blog
	if err := r.db.WithContext(ctx).Where("author_id = ?", authorID).Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *PostgreSQLBlogRepository) Update(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	if err := r.db.WithContext(ctx).Save(blog).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *PostgreSQLBlogRepository) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&models.Blog{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
