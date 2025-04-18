package handlers

import (
	"encoding/json"
	"microservices-travel-backend/internal/user-service/domain/user"
	"microservices-travel-backend/internal/user-service/interfaces/usecase"
	"microservices-travel-backend/pkg/response"
	validator "microservices-travel-backend/pkg/validation"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(r *mux.Router, authUsecase usecase.AuthUsecase) {
	handler := &AuthHandler{
		authUsecase: authUsecase,
	}

	r.HandleFunc("/auth/register", handler.RegisterUser).Methods("POST")
	r.HandleFunc("/auth/login", handler.LoginUser).Methods("POST")
	r.HandleFunc("/auth/logout", handler.LogoutUser).Methods("POST")
	r.HandleFunc("/auth/reset-password", handler.ResetPassword).Methods("POST")
	r.HandleFunc("/auth/forgot-password", handler.ForgotPassword).Methods("POST")
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userDetails user.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userDetails); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.authUsecase.RegisterUser(r.Context(), &userDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userDetails user.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userDetails); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.authUsecase.LoginUser(r.Context(), &userDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *AuthHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User logged out successfully"})
}

func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var request user.ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.BadRequest(w, "Invalid request payload")
		return
	}

	if err := validator.ValidateStruct(request); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	if err := h.authUsecase.ForgotPassword(r.Context(), request.Email); err != nil {
		response.InternalServerError(w, "Failed to send password reset link")
		return
	}

	response.Success(w, http.StatusOK, "Password reset link sent", nil)
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var request user.ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.BadRequest(w, "Invalid request payload")
		return
	}

	if err := validator.ValidateStruct(request); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	if err := h.authUsecase.ResetPassword(r.Context(), request.Token, request.NewPassword); err != nil {
		response.InternalServerError(w, "Failed to reset password")
		return
	}

	response.Success(w, http.StatusOK, "Password reset successfully", nil)
}
