package delivery

import (
	"encoding/json"
	"microservices-travel-backend/internal/hotel-booking/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type HotelHandler struct {
	hotelUsecase usecase.HotelUsecase
}

func NewHotelHandler(r *mux.Router, hotelUsecase usecase.HotelUsecase) {
	handler := &HotelHandler{
		hotelUsecase: hotelUsecase,
	}

	r.HandleFunc("/hotels/search", handler.SearchHotels).Methods("GET")
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
