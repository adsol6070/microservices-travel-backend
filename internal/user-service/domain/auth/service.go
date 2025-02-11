package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"microservices-travel-backend/internal/user-service/domain/user"
	"microservices-travel-backend/pkg/security"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo user.UserRepository
	authRepo AuthRepository
	logger   *log.Logger
}

func NewAuthService(userRepo user.UserRepository, authRepo AuthRepository, logger *log.Logger) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		authRepo: authRepo,
		logger:   logger,
	}
}

func (s *AuthService) Register(ctx context.Context, userDetails *user.User) error {
	exists, err := s.userRepo.Exists(ctx, userDetails.Email)
	if err != nil {
		return fmt.Errorf("failed to check existing user: %w", err)
	}
	if exists {
		return errors.New("user already exists")
	}

	hashedPassword, err := security.HashPassword(userDetails.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	newUser := &user.User{
		ID:        uuid.New().String(),
		Email:     userDetails.Email,
		Name:      userDetails.Name,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.userRepo.Create(ctx, newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, userDetails *user.User) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, userDetails.Email)
	if err != nil {
		return "", fmt.Errorf("failed to fetch user: %w", err)
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	if !security.ValidatePassword(user.Password, userDetails.Password) {
		return "", errors.New("invalid credentials")
	}

	claims := map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
	}
	token, err := security.GenerateJWT(claims)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, token, newPassword string) error {
	claims, err := security.ValidateJWT(token)
	if err != nil {
		return fmt.Errorf("invalid or expired token: %w", err)
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return errors.New("invalid token payload")
	}

	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	hashedPassword, err := security.HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	err = s.authRepo.UpdatePassword(ctx, user.ID, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

func (s *AuthService) ForgotPassword(ctx context.Context, email string) (string, *user.User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	if user == nil {
		return "", nil, errors.New("user not found")
	}

	resetToken, err := security.GenerateJWT(map[string]interface{}{
		"userID": user.ID,
	})
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate reset token: %w", err)
	}

	return resetToken, user, nil
}
