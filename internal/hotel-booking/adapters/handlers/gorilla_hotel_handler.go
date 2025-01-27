package handlers

import (
	"encoding/json"
	"log"
	// "microservices-travel-backend/internal/hotel-booking/domain/models"
	"microservices-travel-backend/internal/hotel-booking/domain/ports"
	"net/http"

	"github.com/gorilla/mux"
)

type HotelHandler struct {
	service ports.HotelService
}

func NewHotelHandler(service ports.HotelService) *HotelHandler {
	return &HotelHandler{service: service}
}

func (h *HotelHandler) RegisterRoutes(router *mux.Router) {
	hotelRouter := router.PathPrefix("/hotels").Subrouter()
	hotelRouter.HandleFunc("/", h.SearchHotelsHandler).Methods(http.MethodGet)
	// hotelRouter.HandleFunc("/{id}", h.GetHotelByID).Methods(http.MethodGet)
	// hotelRouter.HandleFunc("/", h.CreateHotelHandler).Methods(http.MethodPost)
	// hotelRouter.HandleFunc("/{id}", h.UpdateHotelHandler).Methods(http.MethodPatch)
	// hotelRouter.HandleFunc("/{id}", h.DeleteHotelHandler).Methods(http.MethodDelete)

	// // Booking related routes
	// hotelRouter.HandleFunc("/bookings", h.CreateBookingHandler).Methods(http.MethodPost)
	// hotelRouter.HandleFunc("/bookings/{id}", h.UpdateBookingStatusHandler).Methods(http.MethodPatch)
	// hotelRouter.HandleFunc("/bookings/{id}/cancel", h.CancelBookingHandler).Methods(http.MethodPost)
	// hotelRouter.HandleFunc("/availability", h.GetHotelAvailabilityHandler).Methods(http.MethodGet)
}

func (h *HotelHandler) SearchHotelsHandler(w http.ResponseWriter, r *http.Request) {
	// Use service to fetch hotels based on search parameters
	hotels, err := h.service.FetchHotels()
	if err != nil {
		log.Printf("Error fetching hotels: %v", err)
		http.Error(w, "Failed to fetch hotels", http.StatusInternalServerError)
		return
	}

	// Return the hotels data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotels)
}
