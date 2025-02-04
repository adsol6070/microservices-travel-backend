package ports

import "microservices-travel-backend/internal/user-service/domain/models"

type TokenService interface {
	StoreToken(token models.Token) (*models.Token, error)
	GetToken(userID string) (*models.Token, error)
	DeleteToken(userID string) error
}
