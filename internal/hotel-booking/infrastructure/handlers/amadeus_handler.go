package handlers

import (
	"encoding/json"
	"microservices-travel-backend/internal/hotel-booking/app/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type HotelHandler struct {
	hotelUsecase *usecase.HotelUsecase
}

func NewHotelHandler(r *mux.Router, hotelUsecase *usecase.HotelUsecase) {
	handler := &HotelHandler{
		hotelUsecase: hotelUsecase,
	}

	r.HandleFunc("/hotels/search", handler.SearchHotels).Methods("GET")
	r.HandleFunc("/hotels/offers", handler.FetchHotelOffers).Methods("GET")
}

func (h *HotelHandler) SearchHotels(w http.ResponseWriter, r *http.Request) {
	cityCode := r.URL.Query().Get("cityCode")
	if cityCode == "" {
		http.Error(w, "City code is required", http.StatusBadRequest)
		return
	}

	hotels, err := h.hotelUsecase.SearchHotels(cityCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hotels)
}

func (h *HotelHandler) FetchHotelOffers(w http.ResponseWriter, r *http.Request) {
	hotelIDsParam := r.URL.Query().Get("hotelIds")
	adultsParam := r.URL.Query().Get("adults")

	if hotelIDsParam == "" || adultsParam == "" {
		http.Error(w, "hotelIds and adults parameters are required", http.StatusBadRequest)
		return
	}

	hotelIDs := strings.Split(hotelIDsParam, ",")
	adults, err := strconv.Atoi(adultsParam)
	if err != nil || adults <= 0 {
		http.Error(w, "Invalid value for adults", http.StatusBadRequest)
		return
	}

	offers, err := h.hotelUsecase.FetchHotelOffers(hotelIDs, adults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(offers)
}
