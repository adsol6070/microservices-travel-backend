package user

import (
	"context"
	"errors"
	"microservices-travel-backend/pkg/security"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, userDetails *User) (*User, error) {
	if err := s.validateUserDetails(userDetails); err != nil {
		return nil, err
	}

	exists, err := s.userRepo.Exists(ctx, userDetails.Email)
	if err != nil {
		return nil, errors.New("error checking user existence")
	}
	if exists {
		return nil, errors.New("user already exists with this email")
	}

	hashedPassword, err := security.HashPassword(userDetails.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	newUser := &User{
		ID:        uuid.New().String(),
		Email:     userDetails.Email,
		Name:      userDetails.Name,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := s.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	return createdUser, nil
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (*User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userID string, updatedDetails *User) (*User, error) {
	existingUser, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if updatedDetails.Name != "" {
		existingUser.Name = updatedDetails.Name
	}
	if updatedDetails.Email != "" && updatedDetails.Email != existingUser.Email {
		existingUser.Email = updatedDetails.Email
	}

	updatedUser, err := s.userRepo.Update(ctx, existingUser)
	if err != nil {
		return nil, errors.New("failed to update user")
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	existingUser, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	if err := s.userRepo.Delete(ctx, existingUser.ID); err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*User, error) {
	users, err := s.userRepo.ListAll(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve users")
	}
	return users, nil

}

func (s *UserService) validateUserDetails(user *User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email")
	}

	return nil
}

func isValidEmail(email string) bool {
	return len(email) > 3 && string(email[len(email)-1]) != "@" && len(email) < 254
}
