package service

import (
	"errors"
	"microservices-travel-backend/internal/user-service/domain/models"
	"microservices-travel-backend/internal/user-service/domain/ports"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo ports.UserRepositoryPort
}

var jwtKey = []byte("secret-key")

func NewUserService(userRepo ports.UserRepositoryPort) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Replace the plain password with the hashed one
	user.Password = string(hashedPassword)
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
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := &jwt.RegisteredClaims{
		Subject:   user.ID,
		Issuer:    "royal-dusk",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
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
