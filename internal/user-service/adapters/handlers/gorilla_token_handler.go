package handlers

import (
	"encoding/json"
	"net/http"

	"microservices-travel-backend/internal/user-service/domain/models"
	"microservices-travel-backend/internal/user-service/domain/ports"

	"github.com/gorilla/mux"
)

type TokenHandler struct {
	service ports.TokenService
}

func NewTokenHandler(service ports.TokenService) *TokenHandler {
	return &TokenHandler{service: service}
}

func (h *TokenHandler) StoreTokenHandler(w http.ResponseWriter, r *http.Request) {
	var token models.Token
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newToken, err := h.service.StoreToken(token)
	if err != nil {
		http.Error(w, "Failed to store token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newToken)
}

func (h *TokenHandler) GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	token, err := h.service.GetToken(userID)
	if err != nil {
		http.Error(w, "Token not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(token)
}

func (h *TokenHandler) DeleteTokenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	err := h.service.DeleteToken(userID)
	if err != nil {
		http.Error(w, "Failed to delete token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
