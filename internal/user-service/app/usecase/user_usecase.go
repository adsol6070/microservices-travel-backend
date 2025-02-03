package usecase

import (
	"context"
	"microservices-travel-backend/internal/user-service/domain/user"
	"microservices-travel-backend/internal/user-service/interfaces/service"
)

type UserUsecaseImpl struct {
	userService service.UserService
}

func NewUserUsecase(userService service.UserService) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		userService: userService,
	}
}

func (u *UserUsecaseImpl) CreateUser(ctx context.Context, userDetails *user.User) (*user.User, error) {
	createdUser, err := u.userService.CreateUser(ctx, userDetails)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (u *UserUsecaseImpl) GetUser(ctx context.Context, userID string) (*user.User, error) {
	user, err := u.userService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecaseImpl) UpdateUser(ctx context.Context, userID string, updatedDetails *user.User) (*user.User, error) {
	updatedUser, err := u.userService.UpdateUser(ctx, userID, updatedDetails)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *UserUsecaseImpl) DeleteUser(ctx context.Context, userID string) error {
	err := u.userService.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecaseImpl) GetUsers(ctx context.Context) ([]*user.User, error) {
	users, err := u.userService.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
