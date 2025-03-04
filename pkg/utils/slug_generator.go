package utils

import (
	"context"
	"strconv"

	"github.com/gosimple/slug"
)

// SlugExistsFunc defines a function type that checks if a slug exists
type SlugExistsFunc func(ctx context.Context, slug string) bool

// GenerateUniqueSlug creates a unique slug for a given title
func GenerateUniqueSlug(ctx context.Context, title string, slugExists SlugExistsFunc) string {
	baseSlug := slug.Make(title) // Convert title to a slug
	slugToCheck := baseSlug
	count := 1

	// Ensure the slug is unique
	for slugExists(ctx, slugToCheck) {
		slugToCheck = baseSlug + "-" + strconv.Itoa(count)
		count++
	}

	return slugToCheck
}
