package models

import "time"

type BlogCategory struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=50"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBlogCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3,max=50"`
}

type UpdateBlogCategoryRequest struct {
	Name *string `json:"name,omitempty" validate:"omitempty,min=3,max=50"`
}

type DeleteBlogCategoryRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}