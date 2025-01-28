package handlers

import (
	"encoding/json"
	"microservices-travel-backend/internal/user-service/domain/models"
	"microservices-travel-backend/internal/user-service/domain/ports"
	"net/http"

	"github.com/gorilla/mux"
)

// Mock users data.
var users = []models.User{
	{ID: "1", Username: "John Doe", Email: "john@example.com", Password: "password123"},
	{ID: "2", Username: "Jane Doe", Email: "jane@example.com", Password: "password123"},
}

// UserHandler handles user-related operations.
type UserHandler struct {
	service ports.UserServicePort
}

// NewUserHandler initializes a new UserHandler.
func NewUserHandler(service ports.UserServicePort) *UserHandler {
	return &UserHandler{service: service}
}

// RegisterAuthHandlers registers the auth-related routes.
func (h *UserHandler) RegisterAuthHandlers(router *mux.Router) {
	router.HandleFunc("/register", h.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", h.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", h.LogoutHandler).Methods("POST")
}

// RegisterHandler registers a new user.
func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields.
	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Name, Email, and Password are required", http.StatusBadRequest)
		return
	}

	// Check if the email is already registered.
	for _, existingUser := range users {
		if existingUser.Email == user.Email {
			http.Error(w, "Email is already registered", http.StatusConflict)
			return
		}
	}

	// Add the new user (ID generation for simplicity).
	user.ID = generateUniqueID()
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully", "user_id": user.ID})
}

// LoginHandler logs in a user and issues a mock token.
func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate credentials.
	for _, user := range users {
		if user.Email == credentials.Email && user.Password == credentials.Password {
			// Issue a mock token (in production, use JWT or other token systems).
			token := "mockToken123" // Replace with a real token generation logic.
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "token": token})
			return
		}
	}

	http.Error(w, "Invalid email or password", http.StatusUnauthorized)
}

// LogoutHandler handles user logout.
func (h *UserHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate token invalidation (implementation depends on your token system).
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}

// generateUniqueID generates a mock unique ID for users.
func generateUniqueID() string {
	return string(len(users) + 1) // This is a mock implementation.
}
