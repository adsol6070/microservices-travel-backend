package service

import (
	"errors"
	"microservices-travel-backend/internal/user-service/domain/models"
	"microservices-travel-backend/internal/user-service/domain/ports"
)

type UserService struct {
	userRepo ports.UserRepositoryPort
}

func NewUserService(userRepo ports.UserRepositoryPort) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *UserService) Login(creds models.Credentials) (string, error) {
	user, err := s.userRepo.GetByEmail(creds.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Validate password (in practice, compare hashed passwords)
	if user.Password != creds.Password {
		return "", errors.New("invalid credentials")
	}

	// Generate token (for simplicity, returning a dummy token here)
	token := "dummy-jwt-token" // You can use JWT library to generate a proper token
	return token, nil
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) UpdateUser(id string, user models.User) (*models.User, error) {
	// You can add validation here (e.g., ensuring the user exists)
	updatedUser, err := s.userRepo.Update(id, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(id string) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
