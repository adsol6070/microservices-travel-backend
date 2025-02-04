package services

import (
	"microservices-travel-backend/internal/user-service/domain/models"
	"microservices-travel-backend/internal/user-service/domain/ports"
)

type TokenService struct {
	repo ports.TokenService
}

func NewTokenService(repo ports.TokenService) *TokenService {
	return &TokenService{repo: repo}
}

func (s *TokenService) StoreToken(token models.Token) (*models.Token, error) {
	return s.repo.StoreToken(token)
}

func (s *TokenService) GetToken(userID string) (*models.Token, error) {
	return s.repo.GetToken(userID)
}

func (s *TokenService) DeleteToken(userID string) error {
	return s.repo.DeleteToken(userID)
}
