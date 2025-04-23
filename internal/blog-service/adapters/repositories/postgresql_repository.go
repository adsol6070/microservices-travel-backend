package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"microservices-travel-backend/internal/blog-service/domain/models"
	"os"
	"path/filepath"
	"strings"
	"time"

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

func (r *PostgreSQLBlogRepository) Create(ctx context.Context, blogReq *models.CreateBlogRequest) (*models.Blog, error) {
	tx := r.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	blogID := uuid.New()

	blog := &models.Blog{
		ID:              blogID,
		Title:           blogReq.Title,
		Slug:            blogReq.Slug,
		Author:          blogReq.Author,
		CategoryID:      blogReq.CategoryID,
		Content:         blogReq.Content,
		Excerpt:         blogReq.Excerpt,
		MetaTitle:       blogReq.MetaTitle,
		MetaDescription: blogReq.MetaDescription,
		Thumbnail:       blogReq.Thumbnail,
		Status:          blogReq.Status,
		PublishedAt:     blogReq.PublishedAt,
		ScheduledAt:     blogReq.ScheduledAt,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := tx.Table("blogs").Create(blog).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, tag := range blogReq.Tags {
		blogTag := models.BlogTag{
			BlogID: blogID,
			Tag:    tag,
		}
		if err := tx.Table("blog_tags").Create(&blogTag).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return blog, nil
}

func (r *PostgreSQLBlogRepository) GetByID(ctx context.Context, id string) (*models.Blog, error) {
	var blog models.Blog
	query := `SELECT 
			b.*, 
			COALESCE(ARRAY_AGG(bt.tag) FILTER (WHERE bt.tag IS NOT NULL), ARRAY[]::TEXT[]) AS tags,
			c.name AS category_name
		FROM blogs b 
		LEFT JOIN blog_tags bt ON b.id = bt.blog_id 
		LEFT JOIN blog_categories c ON b.category_id = c.id
		WHERE b.id = $1 
		GROUP BY b.id, c.name;`

	if err := r.db.WithContext(ctx).Raw(query, id).Scan(&blog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &blog, nil
}

func (r *PostgreSQLBlogRepository) GetByCategory(ctx context.Context, category string) ([]*models.Blog, error) {
	var blog []*models.Blog
	if err := r.db.WithContext(ctx).Table("blogs").Where("category = ?", category).Find(&blog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return blog, nil
}

func (r *PostgreSQLBlogRepository) GetAll(ctx context.Context) ([]models.Blog, error) {
	var blogs []models.Blog

	query := `SELECT b.*, c.name AS category_name
	    FROM blogs b
		LEFT JOIN blog_categories c ON b.category_id = c.id 
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2`

	if err := r.db.WithContext(ctx).Raw(query, 10, 0).Scan(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *PostgreSQLBlogRepository) GetByAuthor(ctx context.Context, authorID string) ([]*models.Blog, error) {
	var blogs []*models.Blog
	if err := r.db.WithContext(ctx).Table("blogs").Where("author_id = ?", authorID).Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *PostgreSQLBlogRepository) Update(ctx context.Context, id string, blogUpdates map[string]interface{}) (*models.Blog, error) {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var blog models.Blog
	if err := tx.Table("blogs").First(&blog, "id = ?", id).Error; err != nil {
		tx.Rollback()
		log.Printf("Error fetching blog with ID %s: %v", id, err)
		return nil, err
	}

	var tags []string
	if rawTags, exists := blogUpdates["tags"]; exists {
		tags, _ = rawTags.([]string)
		delete(blogUpdates, "tags")
	}

	if err := tx.Table("blogs").Where("id = ?", id).Updates(blogUpdates).Error; err != nil {
		tx.Rollback()
		log.Printf("Error updating blog ID %s: %v", id, err)
		return nil, err
	}

	if len(tags) > 0 {
		if err := tx.Table("blog_tags").Where("blog_id = ?", id).Delete(nil).Error; err != nil {
			tx.Rollback()
			log.Printf("Error deleting existing tags for blog ID %s: %v", id, err)
			return nil, err
		}

		for _, tag := range tags {
			blogTag := models.BlogTag{
				BlogID: uuid.MustParse(id),
				Tag:    tag,
			}
			if err := tx.Table("blog_tags").Create(&blogTag).Error; err != nil {
				tx.Rollback()
				log.Printf("Error inserting new tag '%s' for blog ID %s: %v", tag, id, err)
				return nil, err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Transaction commit failed for blog ID %s: %v", id, err)
		tx.Rollback()
		return nil, err
	}

	return &blog, nil
}

func (r *PostgreSQLBlogRepository) Delete(ctx context.Context, id string) error {
	var blog models.Blog

	if err := r.db.WithContext(ctx).Table("blogs").Where("id = ?", id).First(&blog).Error; err != nil {
		return err
	}

	if blog.Thumbnail != "" {
		imagePath := strings.TrimPrefix(blog.Thumbnail, "http://localhost:7200/")

		if !strings.HasPrefix(imagePath, "uploads/") {
			return fmt.Errorf("invalid image path: %s", imagePath)
		}

		if err := os.Remove(imagePath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to delete image: %v", err)
		}
	}

	if err := r.db.WithContext(ctx).Unscoped().Table("blogs").Delete(&models.Blog{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLBlogRepository) SlugExists(ctx context.Context, slug string) bool {
	var count int64
	err := r.db.WithContext(ctx).Table("blogs").Where("slug = ?", slug).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}
