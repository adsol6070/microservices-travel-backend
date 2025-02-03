package service

import (
	"context"
	"microservices-travel-backend/internal/user-service/domain/user"
)

type UserService interface {
	CreateUser(ctx context.Context, userDetails *user.User) (*user.User, error)
	GetUserByID(ctx context.Context, userID string) (*user.User, error)
	UpdateUser(ctx context.Context, userID string, updatedDetails *user.User) (*user.User, error)
	DeleteUser(ctx context.Context, userID string) error
	GetAllUsers(ctx context.Context) ([]*user.User, error)
}
