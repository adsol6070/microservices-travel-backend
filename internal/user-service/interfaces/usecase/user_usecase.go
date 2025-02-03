package usecase

import (
	"context"
	"microservices-travel-backend/internal/user-service/domain/user"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, userDetails *user.User) (*user.User, error)
	GetUser(ctx context.Context, userID string) (*user.User, error)
	UpdateUser(ctx context.Context, userID string, updatedDetails *user.User) (*user.User, error)
	DeleteUser(ctx context.Context, userID string) error
	GetUsers(ctx context.Context) ([]*user.User, error)
}
