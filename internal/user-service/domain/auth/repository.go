package auth

import "context"

type AuthRepository interface {
	UpdatePassword(ctx context.Context, userID string, newPasswordHash string) error
}
