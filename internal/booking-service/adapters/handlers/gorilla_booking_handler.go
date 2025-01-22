package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"microservices-travel-backend/internal/booking-service/domain/models"
	"microservices-travel-backend/internal/booking-service/domain/ports"
	"net/http"
)

type BookingHandler struct {
	service ports.BookingService
}

func NewBookingHandler(service ports.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) RegisterRoutes(router *mux.Router) {
	// Hotel routes
	bookingRouter := router.PathPrefix("/bookings").Subrouter()
	bookingRouter.HandleFunc("/", h.CreateBooking).Methods(http.MethodPost)
	bookingRouter.HandleFunc("/", h.GetBookings).Methods(http.MethodGet)
	bookingRouter.HandleFunc("/{id}", h.GetBookingByID).Methods(http.MethodGet)
	bookingRouter.HandleFunc("/{id}", h.UpdateBooking).Methods(http.MethodPatch)
	bookingRouter.HandleFunc("/status/{id}", h.UpdateBookingStatus).Methods(http.MethodPatch)
	bookingRouter.HandleFunc("/{id}", h.DeleteBooking).Methods(http.MethodDelete)
}

func (h *BookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.service.GetAllBookings()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookings)
}

func (h *BookingHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	booking, err := h.service.GetBookingByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(booking)
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking

	// Parse JSON body
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call service to create booking
	err = h.service.CreateBooking(&booking)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the created booking
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func (h *BookingHandler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var booking models.Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedBooking, err := h.service.UpdateBooking(id, &booking)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBooking)
}

func (h *BookingHandler) UpdateBookingStatus(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var statusRequest struct {
		Status string `json:"bookingStatus"`
	}
	err := json.NewDecoder(r.Body).Decode(&statusRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateBookingStatus(id, statusRequest.Status)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{"Booking status updated successfully"})
}

func (h *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.service.DeleteBooking(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BookingHandler) GetBookingsByUserID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]

	bookings, err := h.service.GetBookingsByUserID(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookings)
}
