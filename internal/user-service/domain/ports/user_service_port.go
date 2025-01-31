package ports

import "microservices-travel-backend/internal/user-service/domain/models"

type UserServicePort interface {
	CreateUser(user models.User) (*models.User, error)
	Login(creds models.Credentials) (string, error)
	ForgotPassword(email string) error
	ResetPassword(token string, newPassword string) error
	GetUserByID(id string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(id string, user models.User) (*models.User, error)
	DeleteUser(id string) error
}
