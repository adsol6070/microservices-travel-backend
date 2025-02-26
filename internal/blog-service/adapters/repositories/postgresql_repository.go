package repositories

import (
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

func (repo *PostgreSQLBlogRepository) Create(blog models.Blog) (*models.Blog, error) {
	if blog.ID == "" {
		blog.ID = uuid.New().String()
	}
	if err := repo.db.Create(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

func (repo *PostgreSQLBlogRepository) GetByID(id string) (*models.Blog, error) {
	var blog models.Blog
	if err := repo.db.First(&blog, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("blog not found")
		}
		return nil, err
	}
	return &blog, nil
}

func (repo *PostgreSQLBlogRepository) GetAll() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := repo.db.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (repo *PostgreSQLBlogRepository) GetByAuthor(authorID string) ([]models.Blog, error) {
	var blogs []models.Blog
	if err := repo.db.Where("author_id = ?", authorID).Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (repo *PostgreSQLBlogRepository) Update(id string, blog models.Blog) (*models.Blog, error) {
	var existingBlog models.Blog
	if err := repo.db.First(&existingBlog, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("blog not found")
		}
		return nil, err
	}

	if err := repo.db.Model(&existingBlog).Updates(blog).Error; err != nil {
		return nil, err
	}

	if err := repo.db.First(&existingBlog, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &existingBlog, nil
}

func (repo *PostgreSQLBlogRepository) Delete(id string) error {
	if err := repo.db.Delete(&models.Blog{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
