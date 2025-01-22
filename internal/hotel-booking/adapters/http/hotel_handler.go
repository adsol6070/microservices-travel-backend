package http

import (
	"encoding/json"
	"fmt"
	"microservices-travel-backend/internal/hotel-booking/domain/models"
	"microservices-travel-backend/internal/hotel-booking/domain/ports"
	"net/http"
	"strings"
)

type HotelHandler struct {
	service ports.HotelService
}

func NewHotelHandler(service ports.HotelService) *HotelHandler {
	return &HotelHandler{service: service}
}

// CreateHotel handles creating a new hotel
func (h *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var hotel models.Hotel

	// Parse JSON body
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdHotel, err := h.service.CreateHotel(&hotel)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdHotel)
}

// GetHotelByID handles fetching a hotel by ID
func (h *HotelHandler) GetHotelByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/hotels/")

	hotel, err := h.service.GetHotelByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hotel)
}

// UpdateHotel handles updating an existing hotel
func (h *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/hotels/")

	var hotel models.Hotel
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedHotel, err := h.service.UpdateHotel(id, &hotel)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedHotel)
}

// DeleteHotel handles deleting a hotel by ID
func (h *HotelHandler) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/hotels/")

	err := h.service.DeleteHotel(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
