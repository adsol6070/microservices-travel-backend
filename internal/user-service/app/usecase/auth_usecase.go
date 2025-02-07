package usecase

import (
	"context"
	"log"
	"microservices-travel-backend/internal/user-service/domain/email"
	"microservices-travel-backend/internal/user-service/domain/user"
	"microservices-travel-backend/internal/user-service/interfaces/service"
)

type AuthUsecaseImpl struct {
	authService  service.AuthService
	emailService *email.EmailService
}

func NewAuthUsecase(authService service.AuthService, emailService *email.EmailService) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{
		authService:  authService,
		emailService: emailService,
	}
}

func (u *AuthUsecaseImpl) RegisterUser(ctx context.Context, userDetails *user.User) error {
	err := u.authService.Register(ctx, userDetails)
	if err != nil {
		return err
	}

	emailMessage := email.Email{
		To:      userDetails.Email,
		Subject: "Welcome to Our Platform!",
		Body:    "Welcome " + userDetails.Name + "! You have successfully registered.",
	}

	go func() {
		err = u.emailService.SendEmail(emailMessage)
		if err != nil {
			log.Printf("Error sending registration email: %v", err)
		}
	}()

	return nil
}

func (u *AuthUsecaseImpl) LoginUser(ctx context.Context, userDetails *user.User) (string, error) {
	return u.authService.Login(ctx, userDetails)
}

func (u *AuthUsecaseImpl) LogoutUser(ctx context.Context) error {
	return nil
}

func (u *AuthUsecaseImpl) ResetPassword(ctx context.Context, email, newPassword string) error {
	return u.authService.ResetPassword(ctx, email, newPassword)
}

func (u *AuthUsecaseImpl) ForgotPassword(ctx context.Context, email string) error {
	return u.authService.ForgotPassword(ctx, email)
}
