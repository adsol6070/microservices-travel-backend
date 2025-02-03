package handlers

import (
	"encoding/json"
	"microservices-travel-backend/internal/user-service/domain/user"
	"microservices-travel-backend/internal/user-service/interfaces/usecase"
	"net/http"

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
	r.HandleFunc("/users/{id}", handler.GetUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", handler.GetUser).Methods(http.MethodDelete)
	r.HandleFunc("/users", handler.GetUser).Methods(http.MethodGet)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDetails user.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userDetails); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userUsecase.CreateUser(r.Context(), &userDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := h.userUsecase.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var updatedDetails user.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedDetails); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser, err := h.userUsecase.UpdateUser(r.Context(), userID, &updatedDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	err := h.userUsecase.DeleteUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUsecase.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
