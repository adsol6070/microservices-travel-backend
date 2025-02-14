package handlers

import (
	"encoding/json"
	"microservices-travel-backend/internal/hotel-booking/app/usecase"
	"microservices-travel-backend/internal/shared/api_provider/google/places/models"
	"net/http"

	"github.com/gorilla/mux"
)

type PlaceHandler struct {
	placeUsecase *usecase.GooglePlacesUsecase
}

func NewPlaceHandler(r *mux.Router, placeUsecase *usecase.GooglePlacesUsecase) {
	handler := &PlaceHandler{
		placeUsecase: placeUsecase,
	}

	r.HandleFunc("/places/search", handler.SearchPlaces).Methods("POST")
	r.HandleFunc("/places/{placeID}/photos/{photoID}", handler.GetPlacePhoto).Methods("GET")
	r.HandleFunc("/places/{placeID}", handler.GetPlaceDetail).Methods("GET")
}

func (h *PlaceHandler) SearchPlaces(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body
	var req models.TextQueryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate that textQuery is not empty
	if req.TextQuery == "" {
		http.Error(w, "textQuery parameter is required", http.StatusBadRequest)
		return
	}

	// Call the use case to search for places
	places, err := h.placeUsecase.SearchPlaces(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers and send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(places)
}

func (h *PlaceHandler) GetPlacePhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	placeID := vars["placeID"]
	photoID := vars["photoID"]

	maxHeight := 400 // Default max height
	maxWidth := 400  // Default max width

	photo, err := h.placeUsecase.GetPlacePhoto(placeID, photoID, maxHeight, maxWidth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(photo)
}

func (h *PlaceHandler) GetPlaceDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	placeID := vars["placeID"]

	placeDetail, err := h.placeUsecase.GetPlaceDetail(placeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(placeDetail)
}
