package services

import (
	"context"
	"errors"
	"microservices-travel-backend/internal/user-service/domain/models"
	"microservices-travel-backend/internal/user-service/domain/ports"
)

type UserService struct {
	repo ports.UserDBPort
}

func NewUserService(repo ports.UserDBPort) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	return s.repo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is empty")
	}
	return s.repo.DeleteUser(ctx, id)
}
