package usecase

import (
	"context"
	"microservices-travel-backend/internal/user-service/domain/user"
	"microservices-travel-backend/internal/user-service/interfaces/service"
)

type AuthUsecaseImpl struct {
	authService service.AuthService
}

func NewAuthUsecase(authService service.AuthService) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{
		authService: authService,
	}
}

func (u *AuthUsecaseImpl) RegisterUser(ctx context.Context, userDetails *user.User) error {
	return u.authService.Register(ctx, userDetails)
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
