package ports

import (
	"context"
	"microservices-travel-backend/internal/user-service/domain/models"
)

// UserDBPort defines the interface for user database operations
type UserDBPort interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
	AuthenticateUser(ctx context.Context, email, password string) (*models.User, error)
}