package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Blog struct {
	ID              uuid.UUID      `json:"id" gorm:"primaryKey"`
	Title           string         `json:"title" validate:"required,min=5,max=150"`
	Slug            string         `json:"slug" validate:"required"`
	Content         string         `json:"content" validate:"required,min=50"`
	Excerpt         string         `json:"excerpt"`
	MetaTitle       string         `json:"meta_title" validate:"max=150"`
	MetaDescription string         `json:"meta_description" validate:"max=300"`
	AuthorID        uuid.UUID      `json:"author_id" validate:"required,uuid"`
	Category        string         `json:"category" validate:"required"`
	Tags            pq.StringArray `json:"tags" gorm:"type:text[]" validate:"omitempty"`
	Thumbnail       string         `json:"thumbnail"`
	Status          string         `json:"status" validate:"required,oneof=draft published archived"`
	PublishedAt     *time.Time     `json:"published_at,omitempty" gorm:"type:timestamptz"`
	ScheduledAt     *time.Time     `json:"scheduled_at,omitempty" gorm:"type:timestamptz"`
	IsPublished     bool           `json:"is_published"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// CreateBlogRequest - request struct for creating a blog
type CreateBlogRequest struct {
	Title           string         `json:"title" validate:"required,min=5,max=150"`
	Slug            string         `json:"slug" validate:"required"`
	Content         string         `json:"content" validate:"required,min=50"`
	Excerpt         string         `json:"excerpt,omitempty"`
	MetaTitle       string         `json:"meta_title,omitempty" validate:"max=150"`
	MetaDescription string         `json:"meta_description,omitempty" validate:"max=300"`
	AuthorID        uuid.UUID      `json:"author_id" validate:"required,uuid"`
	Category        string         `json:"category" validate:"required"`
	Tags            pq.StringArray `json:"tags" gorm:"type:text[]" validate:"omitempty"`
	Thumbnail       string         `json:"thumbnail,omitempty" validate:"omitempty,url"`
	Status          string         `json:"status" validate:"required,oneof=draft published archived"`
	PublishedAt     *time.Time     `json:"published_at,omitempty" gorm:"type:timestamptz"`
	ScheduledAt     *time.Time     `json:"scheduled_at,omitempty" gorm:"type:timestamptz"`
	IsPublished     bool           `json:"is_published"`
}

// UpdateBlogRequest - request struct for updating a blog
type UpdateBlogRequest struct {
	Title           *string        `json:"title,omitempty" validate:"omitempty,min=5,max=150"`
	Slug            *string        `json:"slug,omitempty" validate:"omitempty"`
	Content         *string        `json:"content,omitempty" validate:"omitempty,min=50"`
	Excerpt         *string        `json:"excerpt,omitempty"`
	MetaTitle       *string        `json:"meta_title,omitempty" validate:"omitempty,max=150"`
	MetaDescription *string        `json:"meta_description,omitempty" validate:"omitempty,max=300"`
	Category        *string        `json:"category,omitempty" validate:"omitempty"`
	Tags            pq.StringArray `json:"tags" validate:"omitempty"`
	Thumbnail       *string        `json:"thumbnail,omitempty" validate:"omitempty,url"`
	Status          *string        `json:"status,omitempty" validate:"omitempty,oneof=draft published archived"`
	PublishedAt     *time.Time     `json:"published_at,omitempty"`
	ScheduledAt     *time.Time     `json:"scheduled_at,omitempty"`
	IsPublished     *bool          `json:"is_published,omitempty"`
}

// DeleteBlogRequest - request struct for deleting a blog
type DeleteBlogRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"` // Use UUID instead of string
}
