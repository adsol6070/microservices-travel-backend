package ports

import (
	"context"
	"microservices-travel-backend/internal/user-service/domain/models"
)

// UserServicePort defines the interface for user authentication services.
type UserServicePort interface {
	RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
	LoginUser(ctx context.Context, email, password string) (string, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userID string) error
}