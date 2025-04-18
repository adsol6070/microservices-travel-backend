package user

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email" validate:"required,email"`
	Name      string    `json:"name" validate:"required,min=2,max=100"`
	Password  string    `json:"password" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdateRequest struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=8"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,min=8"`
}
