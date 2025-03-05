package models

import (
	"time"

	"github.com/lib/pq"
)

type Blog struct {
	ID          string         `json:"id"`
	Title       string         `json:"title" validate:"required,min=5,max=100"`
	Slug        string         `json:"slug"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`
	Content     string         `json:"content" validate:"required,min=20"`
	AuthorID    string         `json:"author_id" validate:"required,uuid"`
	Category    string         `json:"category" validate:"required"`
	Thumbnail   string         `json:"thumbnail"`
	PublishedAt *time.Time     `json:"published_at,omitempty"`
	Status      string         `json:"status" validate:"required,oneof=draft published"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type CreateBlogRequest struct {
	Title     string         `json:"title" validate:"required,min=5,max=100"`
	Content   string         `json:"content" validate:"required,min=20"`
	AuthorID  string         `json:"author_id" validate:"required,uuid"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
	Category  string         `json:"category,omitempty" validate:"omitempty"`
	Thumbnail string         `json:"thumbnail,omitempty" validate:"omitempty,url"`
	Status    string         `json:"status" validate:"required,oneof=draft published"`
}

// UpdateBlogRequest with validation tags
type UpdateBlogRequest struct {
	Title     *string        `json:"title,omitempty" validate:"omitempty,min=5,max=100"`
	Content   *string        `json:"content,omitempty" validate:"omitempty,min=20"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
	Category  *string        `json:"category,omitempty"`
	Thumbnail *string        `json:"thumbnail,omitempty" validate:"omitempty,url"`
	Status    *string        `json:"status,omitempty" validate:"omitempty,oneof=draft published"`
}

// DeleteBlogRequest with validation tags
type DeleteBlogRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}
