package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type BlogTag struct {
	BlogID uuid.UUID `json:"blog_id" gorm:"primaryKey"`
	Tag    string    `json:"tag" gorm:"primaryKey"`
}

type Blog struct {
	ID              uuid.UUID       `json:"id" gorm:"primaryKey"`
	Title           string          `json:"title" validate:"required,min=5,max=150"`
	Slug            string          `json:"slug" validate:"required"`
	Author          string          `json:"author" validate:"required"`
	CategoryID      uuid.UUID       `json:"category_id" validate:"required"`
	CategoryName    string          `json:"category_name"`
	Content         string          `json:"content" validate:"required,min=50"`
	Excerpt         string          `json:"excerpt"`
	MetaTitle       string          `json:"meta_title" validate:"max=150"`
	MetaDescription string          `json:"meta_description" validate:"max=300"`
	Thumbnail       string          `json:"thumbnail"`
	Tags            *pq.StringArray `json:"tags" gorm:"type:text[]"`
	Status          string          `json:"status" validate:"required,oneof=draft published archived"`
	PublishedAt     *time.Time      `json:"published_at,omitempty" gorm:"type:timestamptz"`
	ScheduledAt     *time.Time      `json:"scheduled_at,omitempty" gorm:"type:timestamptz"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type CreateBlogRequest struct {
	Title           string     `json:"title" validate:"required,min=5,max=150"`
	Slug            string     `json:"slug" validate:"required"`
	Author          string     `json:"author" validate:"required"`
	CategoryID      uuid.UUID  `json:"category_id" validate:"required"`
	Content         string     `json:"content" validate:"required,min=50"`
	Tags            []string   `json:"tags" validate:"omitempty"`
	Excerpt         string     `json:"excerpt,omitempty"`
	MetaTitle       string     `json:"meta_title,omitempty" validate:"max=150"`
	MetaDescription string     `json:"meta_description,omitempty" validate:"max=300"`
	Thumbnail       string     `json:"thumbnail,omitempty"`
	Status          string     `json:"status" validate:"required,oneof=draft published archived"`
	PublishedAt     *time.Time `json:"published_at,omitempty" gorm:"type:timestamptz"`
	ScheduledAt     *time.Time `json:"scheduled_at,omitempty" gorm:"type:timestamptz"`
}

type UpdateBlogRequest struct {
	Title           *string    `json:"title,omitempty" validate:"omitempty,min=5,max=150"`
	Slug            *string    `json:"slug,omitempty" validate:"omitempty"`
	Author          *string    `json:"author" validate:"required"`
	CategoryID      *uuid.UUID `json:"category_id,omitempty" validate:"omitempty"`
	Content         *string    `json:"content,omitempty" validate:"omitempty,min=50"`
	Tags            []string   `json:"tags" validate:"omitempty"`
	Excerpt         *string    `json:"excerpt,omitempty"`
	MetaTitle       *string    `json:"meta_title,omitempty" validate:"omitempty,max=150"`
	MetaDescription *string    `json:"meta_description,omitempty" validate:"omitempty,max=300"`
	Thumbnail       *string    `json:"thumbnail,omitempty"`
	Status          *string    `json:"status,omitempty" validate:"omitempty,oneof=draft published archived"`
	PublishedAt     *time.Time `json:"published_at,omitempty"`
	ScheduledAt     *time.Time `json:"scheduled_at,omitempty"`
}

type DeleteBlogRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
