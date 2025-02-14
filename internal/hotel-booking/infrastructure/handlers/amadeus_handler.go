package handlers

import (
	"encoding/json"
	"io"
	"log"
	"microservices-travel-backend/internal/hotel-booking/app/usecase"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/models"
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
	r.HandleFunc("/hotels/book", handler.CreateHotelBooking).Methods("POST")
	r.HandleFunc("/hotels/ratings", handler.FetchHotelRatings).Methods("GET")
	r.HandleFunc("/hotels/hotelnamecomplete", handler.HotelNameAutoComplete).Methods("GET")
}

// Define a struct for the request body
type SearchHotelsRequest struct {
	CityCode     string `json:"cityCode"`
	CheckInDate  string `json:"checkInDate"`
	CheckOutDate string `json:"checkOutDate"`
	Rooms        int    `json:"rooms"`
	Persons      int    `json:"persons"`
}

func (h *HotelHandler) SearchHotels(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO: SearchHotels handler triggered")

	var req SearchHotelsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("ERROR: Failed to decode request body -", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("INFO: Received hotel search request - CityCode: %s, CheckIn: %s, CheckOut: %s, Rooms: %d, Persons: %d",
		req.CityCode, req.CheckInDate, req.CheckOutDate, req.Rooms, req.Persons)

	// Input validation
	if req.CityCode == "" || req.CheckInDate == "" || req.CheckOutDate == "" || req.Rooms <= 0 || req.Persons <= 0 {
		log.Println("WARN: Invalid request parameters detected")
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		return
	}

	usecaseReq := usecase.SearchHotelsRequest{
		CityCode:     req.CityCode,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		Rooms:        req.Rooms,
		Persons:      req.Persons,
	}

	log.Println("INFO: Calling hotel usecase to fetch hotels with offers")

	hotelsWithOffer, err := h.hotelUsecase.SearchHotels(usecaseReq)
	if err != nil {
		log.Println("ERROR: Error occurred while fetching hotel offers -", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("INFO: Successfully retrieved %d hotel offers", len(hotelsWithOffer))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hotelsWithOffer)

	log.Println("INFO: SearchHotels response sent successfully")
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

func (h *HotelHandler) CreateHotelBooking(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var bookingRequest models.HotelBookingRequest
	err = json.Unmarshal(requestBody, &bookingRequest)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	bookingResponse, err := h.hotelUsecase.CreateHotelBooking(bookingRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bookingResponse)
}

func (h *HotelHandler) FetchHotelRatings(w http.ResponseWriter, r *http.Request) {
	hotelIDsParam := r.URL.Query().Get("hotelIds")
	log.Println(hotelIDsParam)

	if hotelIDsParam == "" {
		http.Error(w, "hotelIds parameter are required", http.StatusBadRequest)
		return
	}

	hotelIDs := strings.Split(hotelIDsParam, ",")

	offers, err := h.hotelUsecase.FetchHotelRatings(hotelIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(offers)
}

func (h *HotelHandler) HotelNameAutoComplete(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	subtype := r.URL.Query().Get("subType")

	if keyword == "" || subtype == "" {
		http.Error(w, "keyword and subtype parameters are required", http.StatusBadRequest)
		return
	}

	offers, err := h.hotelUsecase.HotelNameAutoComplete(keyword, subtype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(offers)
}
