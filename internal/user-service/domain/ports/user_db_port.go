package ports

import "microservices-travel-backend/internal/user-service/domain/models"

type UserRepositoryPort interface {
	Create(user models.User) (*models.User, error)
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	ResetPassword(token string, newPassword string) error
	GetAll() ([]models.User, error)
	Update(id string, user models.User) (*models.User, error)
	Delete(id string) error
}
