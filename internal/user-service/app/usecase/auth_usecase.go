package usecase

import (
	"context"
	"log"
	"microservices-travel-backend/internal/user-service/domain/email"
	emailDomain "microservices-travel-backend/internal/user-service/domain/email"
	"microservices-travel-backend/internal/user-service/domain/user"
	"microservices-travel-backend/internal/user-service/interfaces/service"
	"fmt"
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
	resetToken, user, err := u.authService.ForgotPassword(ctx, email)
	if err != nil {
		return err
	}

	resetLink := fmt.Sprintf("https://yourwebsite.com/reset-password?token=%s", resetToken)
	emailMessage := emailDomain.Email{
		To:      user.Email,
		Subject: "Password Reset Request",
		Body:    fmt.Sprintf("Click the link below to reset your password:\n\n%s", resetLink),
	}

	go func() {
		err := u.emailService.SendEmail(emailMessage)
		if err != nil {
			log.Printf("Error sending password reset email: %v", err)
		}
	}()

	return nil
}
