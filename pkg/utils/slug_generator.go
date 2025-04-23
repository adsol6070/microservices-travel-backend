package utils

import (
	"context"
	"strings"

	"github.com/gosimple/slug"
)

// SlugExistsFunc defines a function type that checks if a slug exists
type SlugExistsFunc func(ctx context.Context, slug string) bool

func GenerateUniqueSlug(ctx context.Context, title string) string {
	return slug.Make(title) // Directly generate a slug from the title
}

func GenerateSlugToTitle(slug string) string {
	categoryConverted := strings.ReplaceAll(slug, "-", " ")
	return categoryConverted
}
