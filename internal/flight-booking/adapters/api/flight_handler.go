package api

import (
	"encoding/json"
	"fmt"
	"microservices-travel-backend/internal/flight-booking/domain/models"
	"microservices-travel-backend/internal/flight-booking/domain/ports"
	"net/http"
	"strings"
)

type FlightHandler struct {
	service ports.FlightService
}

func NewFlightHandler(service ports.FlightService) *FlightHandler {
	return &FlightHandler{service: service}
}

// CreateFlight handles creating a new flight
func (h *FlightHandler) CreateFlight(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var flight models.Flight

	// Parse JSON body

	err := json.NewDecoder(r.Body).Decode(&flight)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdFlight, err := h.service.CreateFlight(&flight)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdFlight)
}

// GetFlightByID handles fetching a flight by ID
func (h *FlightHandler) GetFlightByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/flights/")

	flight, err := h.service.GetFlightByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flight)
}

// UpdateFlight handles updating an existing flight
func (h *FlightHandler) UpdateFlight(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/flights/")

	var flight models.Flight
	err := json.NewDecoder(r.Body).Decode(&flight)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedFlight, err := h.service.UpdateFlight(id, &flight)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedFlight)
}

// DeleteFlight handles deleting a flight by ID
func (h *FlightHandler) DeleteFlight(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/flights/")

	err := h.service.DeleteFlight(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
