package handlers

import (
	"microservices-travel-backend/internal/tour-and-packages/domain/ports"
	"net/http"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	tourService ports.TourServicePort
}
	
func NewTourHandler(service ports.TourServicePort) *TourHandler {
	return &TourHandler{tourService: service}
}

func (h *TourHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/tours", h.CreateTour).Methods("POST")
	router.HandleFunc("/tours/{id}", h.GetTourByID).Methods("GET")
	router.HandleFunc("/tours", h.GetAllTours).Methods("GET")
	router.HandleFunc("/tours/{id}", h.UpdateTour).Methods("PATCH")
	router.HandleFunc("/tours/{id}", h.DeleteTour).Methods("DELETE")
}

func (h *TourHandler) CreateTour(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a tour
}

func (h *TourHandler) GetTourByID(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a tour by ID
}

func (h *TourHandler) GetAllTours(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting all tours
}

func (h *TourHandler) UpdateTour(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a tour
}

func (h *TourHandler) DeleteTour(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a tour
}
