package handlers

import (
	"encoding/json"
	"microservices-travel-backend/internal/user-service/domain/user"
	"microservices-travel-backend/internal/user-service/interfaces/usecase"
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

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset successfully"})
}

func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset link sent"})
}
