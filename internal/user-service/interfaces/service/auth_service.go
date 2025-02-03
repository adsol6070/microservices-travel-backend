package service

import (
	"context"
	"microservices-travel-backend/internal/user-service/domain/user"
)

type AuthService interface {
	Register(ctx context.Context, userDetails *user.User) error
	Login(ctx context.Context, userDetails *user.User) (string, error)
	ResetPassword(ctx context.Context, email, newPassword string) error
	ForgotPassword(ctx context.Context, email string) error
}