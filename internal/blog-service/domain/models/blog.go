package models

import "time"

type Blog struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Content     string    `json:"content"`
	AuthorID    string    `json:"author_id"`
	Tags        []string  `json:"tags"`
	Category    string    `json:"category"`
	Thumbnail   string    `json:"thumbnail"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateBlogRequest struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags,omitempty"`
	Category  string   `json:"category,omitempty"`
	Thumbnail string   `json:"thumbnail,omitempty"`
	Status    string   `json:"status"`
}

type UpdateBlogRequest struct {
	Title     *string   `json:"title,omitempty"`
	Content   *string   `json:"content,omitempty"`
	Tags      *[]string `json:"tags,omitempty"`
	Category  *string   `json:"category,omitempty"`
	Thumbnail *string   `json:"thumbnail,omitempty"`
	Status    *string   `json:"status,omitempty"`
}

type DeleteBlogRequest struct {
	ID string `json:"id"`
}
