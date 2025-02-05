package usecase

import (
	"context"
	"microservices-travel-backend/internal/user-service/domain/user"
)

type AuthUsecase interface {
	RegisterUser(ctx context.Context, userDetails *user.User) error
	LoginUser(ctx context.Context, userDetails *user.User) (string, error)
	LogoutUser(ctx context.Context) error
	ResetPassword(ctx context.Context, token, newPassword string) error
	ForgotPassword(ctx context.Context, email string) error
}
