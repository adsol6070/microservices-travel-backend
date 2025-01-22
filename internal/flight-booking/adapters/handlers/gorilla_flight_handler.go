package handlers

import (
	"encoding/json"
	"fmt"
	"microservices-travel-backend/internal/flight-booking/domain/models"
	"microservices-travel-backend/internal/flight-booking/domain/ports"
	"net/http"

	"github.com/gorilla/mux"
)

type FlightHandler struct {
	service ports.FlightService
}

func NewFlightHandler(service ports.FlightService) *FlightHandler {
	return &FlightHandler{service: service}
}

func (h *FlightHandler) RegisterRoutes(r *mux.Router) {
	// Flight routes
	r.HandleFunc("/flights", h.CreateFlight).Methods(http.MethodPost)
	r.HandleFunc("/flights/{id}", h.GetFlightByID).Methods(http.MethodGet)
	r.HandleFunc("/flights/{id}", h.UpdateFlight).Methods(http.MethodPut)
	r.HandleFunc("/flights/{id}", h.DeleteFlight).Methods(http.MethodDelete)

	// Testing route
	r.HandleFunc("/test", h.TestRoute).Methods(http.MethodGet)
}

func (h *FlightHandler) CreateFlight(w http.ResponseWriter, r *http.Request) {
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

func (h *FlightHandler) GetFlightByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	flight, err := h.service.GetFlightByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flight)
}

func (h *FlightHandler) UpdateFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

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

func (h *FlightHandler) DeleteFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.DeleteFlight(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// TestRoute is a simple health check or testing route
func (h *FlightHandler) TestRoute(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple JSON message to verify the service is working
	response := map[string]string{
		"status":  "OK",
		"message": "Flight booking service is up and running!",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
