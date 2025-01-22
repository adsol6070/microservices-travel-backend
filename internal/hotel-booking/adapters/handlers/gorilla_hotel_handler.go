package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"microservices-travel-backend/internal/hotel-booking/domain/models"
	"microservices-travel-backend/internal/hotel-booking/domain/ports"
	"net/http"
)

type HotelHandler struct {
	service ports.HotelService
}

func NewHotelHandler(service ports.HotelService) *HotelHandler {
	return &HotelHandler{service: service}
}

func (h *HotelHandler) RegisterRoutes(router *mux.Router)  {
	// Hotel routes
	hotelRouter := router.PathPrefix("/hotels").Subrouter()
	hotelRouter.HandleFunc("/", h.GetHotels).Methods(http.MethodGet)
	hotelRouter.HandleFunc("/{id}", h.GetHotelByID).Methods(http.MethodGet)
	hotelRouter.HandleFunc("/", h.CreateHotel).Methods(http.MethodPost)
	hotelRouter.HandleFunc("/{id}", h.UpdateHotel).Methods(http.MethodPatch)
	hotelRouter.HandleFunc("/{id}", h.DeleteHotel).Methods(http.MethodDelete)
}

func (h *HotelHandler) GetHotels(w http.ResponseWriter, r *http.Request) {
	hotels, err := h.service.GetAllHotels()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hotels)
}

func (h *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {
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

func (h *HotelHandler) GetHotelByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	hotel, err := h.service.GetHotelByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hotel)
}

func (h *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

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

func (h *HotelHandler) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.service.DeleteHotel(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// TestRoute is a simple health check or testing route
func (h *HotelHandler) TestRoute(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple JSON message to verify the service is working
	response := map[string]string{
		"status":  "OK",
		"message": "Hotel booking service is up and running!",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
