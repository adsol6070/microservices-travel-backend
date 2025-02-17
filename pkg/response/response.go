package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIResponse defines the structure for all JSON responses.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// JSON sends a standardized JSON response.
func JSON(w http.ResponseWriter, statusCode int, success bool, message string, data interface{}, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := APIResponse{
		Success: success,
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	// Encode response and handle encoding errors
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Success sends a standard success response.
func Success(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	JSON(w, statusCode, true, message, data, nil)
}

// SuccessWithMeta sends a success response with additional metadata (e.g., pagination).
func SuccessWithMeta(w http.ResponseWriter, statusCode int, message string, data interface{}, meta interface{}) {
	JSON(w, statusCode, true, message, data, meta)
}

// ErrorResponse sends a structured error response.
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, false, message, nil, nil)
}

// BadRequest sends a 400 Bad Request response.
func BadRequest(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusBadRequest, message)
}

// Unauthorized sends a 401 Unauthorized response.
func Unauthorized(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusUnauthorized, message)
}

// Forbidden sends a 403 Forbidden response.
func Forbidden(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusForbidden, message)
}

// NotFound sends a 404 Not Found response.
func NotFound(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusNotFound, message)
}

// InternalServerError sends a 500 Internal Server Error response.
func InternalServerError(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusInternalServerError, message)
}
