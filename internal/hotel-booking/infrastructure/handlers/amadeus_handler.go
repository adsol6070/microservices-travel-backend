package handlers

import (
	"encoding/json"
	"io"
	"log"
	"microservices-travel-backend/internal/hotel-booking/app/dto/request"
	"microservices-travel-backend/internal/hotel-booking/app/usecase"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/amadeusHotelModels"
	"microservices-travel-backend/pkg/response"
	validator "microservices-travel-backend/pkg/validation"
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

	hotelRouter := r.PathPrefix("/hotels").Subrouter()

	hotelRouter.HandleFunc("/search", handler.SearchHotels).Methods(http.MethodPost)
	hotelRouter.HandleFunc("/detail", handler.HotelDetails).Methods(http.MethodPost)
	hotelRouter.HandleFunc("/offers", handler.FetchHotelOffers).Methods(http.MethodGet)
	hotelRouter.HandleFunc("/book", handler.CreateHotelBooking).Methods(http.MethodPost)
	hotelRouter.HandleFunc("/ratings", handler.FetchHotelRatings).Methods(http.MethodGet)
	hotelRouter.HandleFunc("/hotelnamecomplete", handler.HotelNameAutoComplete).Methods(http.MethodGet)
}

func (h *HotelHandler) SearchHotels(w http.ResponseWriter, r *http.Request) {
	var req request.HotelSearchRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("ERROR: Failed to decode request body -", err)
		response.BadRequest(w, "Invalid request body")
		return
	}

	if err := validator.ValidateStruct(req); err != nil {
		log.Println("ERROR: Validation failed -", err)
		response.BadRequest(w, err.Error())
		return
	}

	hotelsWithOffer, err := h.hotelUsecase.SearchHotels(req)
	if err != nil {
		log.Println("ERROR: Error occurred while fetching hotel offers -", err)
		response.InternalServerError(w, "Failed to fetch hotel offers")
		return
	}

	response.Success(w, http.StatusOK, "Hotels fetched successfully", hotelsWithOffer)
}

func (h *HotelHandler) HotelDetails(w http.ResponseWriter, r *http.Request) {
	var req request.HotelDetailsRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("ERROR: Failed to decode request body -", err)
		response.BadRequest(w, "Invalid request body")
		return
	}

	if err := validator.ValidateStruct(req); err != nil {
		log.Println("ERROR: Validation failed -", err)
		response.BadRequest(w, err.Error())
		return
	}

	hotelDetails, err := h.hotelUsecase.HotelDetails(req)
	if err != nil {
		log.Println("ERROR: Error occured while fetching hotel details -", err)
		response.InternalServerError(w, "Failed to fetch hotel details")
	}

	response.Success(w, http.StatusOK, "Hotel Details fetched successfully", hotelDetails)
}

func (h *HotelHandler) FetchHotelOffers(w http.ResponseWriter, r *http.Request) {
	hotelIDsParam := r.URL.Query().Get("hotelIds")
	adultsParam := r.URL.Query().Get("adults")

	if hotelIDsParam == "" || adultsParam == "" {
		response.BadRequest(w, "hotelIds and adults parameters are required")
		return
	}

	hotelIDs := strings.Split(hotelIDsParam, ",")
	adults, err := strconv.Atoi(adultsParam)
	if err != nil || adults <= 0 {
		response.BadRequest(w, "Invalid value for adults")
		return
	}

	offers, err := h.hotelUsecase.FetchHotelOffers(hotelIDs, adults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Success(w, http.StatusOK, "Hotel offers retrieved successfully", offers)
}

func (h *HotelHandler) CreateHotelBooking(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.BadRequest(w, "Failed to read request body")
		return
	}
	defer r.Body.Close()

	var bookingRequest amadeusHotelModels.HotelBookingRequest
	err = json.Unmarshal(requestBody, &bookingRequest)
	if err != nil {
		response.BadRequest(w, "Invalid request format")
		return
	}

	bookingResponse, err := h.hotelUsecase.CreateHotelBooking(bookingRequest)
	if err != nil {
		log.Println("ERROR: Failed to create hotel booking -", err)
		response.InternalServerError(w, "Failed to create hotel booking")
		return
	}

	response.Success(w, http.StatusCreated, "Hotel booking created successfully", bookingResponse)
}

func (h *HotelHandler) FetchHotelRatings(w http.ResponseWriter, r *http.Request) {
	hotelIDsParam := r.URL.Query().Get("hotelIds")
	log.Println(hotelIDsParam)

	if hotelIDsParam == "" {
		response.BadRequest(w, "hotelIds parameter is required")
		return
	}

	hotelIDs := strings.Split(hotelIDsParam, ",")
	ratings, err := h.hotelUsecase.FetchHotelRatings(hotelIDs)
	if err != nil {
		log.Println("ERROR: Failed to fetch hotel ratings -", err)
		response.InternalServerError(w, "Failed to fetch hotel ratings")
		return
	}

	response.Success(w, http.StatusOK, "Hotel ratings fetched successfully", ratings)
}

func (h *HotelHandler) HotelNameAutoComplete(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	subtype := r.URL.Query().Get("subType")

	if keyword == "" || subtype == "" {
		response.BadRequest(w, "keyword and subtype parameters are required")
		return
	}

	suggestions, err := h.hotelUsecase.HotelNameAutoComplete(keyword, subtype)
	if err != nil {
		log.Println("ERROR: Failed to fetch hotel name suggestions -", err)
		response.InternalServerError(w, "Failed to fetch hotel name suggestions")
		return
	}

	response.Success(w, http.StatusOK, "Hotel name suggestions retrieved", suggestions)
}
