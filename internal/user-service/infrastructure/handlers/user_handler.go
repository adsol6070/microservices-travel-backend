package handlers

import (
	"encoding/json"
	"net/http"

	"microservices-travel-backend/internal/user-service/domain/user"
	"microservices-travel-backend/internal/user-service/interfaces/usecase"
	"microservices-travel-backend/pkg/response"
	validator "microservices-travel-backend/pkg/validation"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(r *mux.Router, userUsecase usecase.UserUsecase) {
	handler := &UserHandler{
		userUsecase: userUsecase,
	}

	r.HandleFunc("/users", handler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", handler.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/users", handler.GetUsers).Methods(http.MethodGet)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDetails user.User
	if err := json.NewDecoder(r.Body).Decode(&userDetails); err != nil {
		response.BadRequest(w, "Invalid request payload")
		return
	}

	if err := validator.ValidateStruct(userDetails); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	createdUser, err := h.userUsecase.CreateUser(r.Context(), &userDetails)
	if err != nil {
		response.InternalServerError(w, "Failed to create user")
		return
	}

	response.Success(w, http.StatusCreated, "User created successfully", createdUser)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	user, err := h.userUsecase.GetUser(r.Context(), userID)
	if err != nil {
		response.NotFound(w, "User not found")
		return
	}

	response.Success(w, http.StatusOK, "User retrieved successfully", user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	var updatedDetails user.User
	if err := json.NewDecoder(r.Body).Decode(&updatedDetails); err != nil {
		response.BadRequest(w, "Invalid request payload")
		return
	}

	if err := validator.ValidateStruct(updatedDetails); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	updatedUser, err := h.userUsecase.UpdateUser(r.Context(), userID, &updatedDetails)
	if err != nil {
		response.InternalServerError(w, "Failed to update user")
		return
	}

	response.Success(w, http.StatusOK, "User updated successfully", updatedUser)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	if err := h.userUsecase.DeleteUser(r.Context(), userID); err != nil {
		response.NotFound(w, "User not found")
		return
	}

	response.Success(w, http.StatusNoContent, "User deleted successfully", nil)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.GetUsers(r.Context())
	if err != nil {
		response.InternalServerError(w, "Could not retrieve users")
		return
	}

	response.Success(w, http.StatusOK, "Users retrieved successfully", users)
}
