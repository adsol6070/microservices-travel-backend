package main

import (
	"log"
	"net/http"
	"os"

	"microservices-travel-backend/internal/hotel-booking/app/usecase"
	"microservices-travel-backend/internal/hotel-booking/domain/amadeus"
	"microservices-travel-backend/internal/hotel-booking/domain/google"
	"microservices-travel-backend/internal/hotel-booking/infrastructure/handlers"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels"
	"microservices-travel-backend/internal/shared/api_provider/google/places"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	logger := log.New(os.Stdout, "hotel-service: ", log.LstdFlags|log.Lshortfile)

	client := hotels.NewAmadeusClient(os.Getenv("AMADEUS_API_KEY"), os.Getenv("AMADEUS_SECRET"), os.Getenv("REDIS_URL"))
	googlePlacesClient := places.NewPlacesClient(os.Getenv("GOOGLE_PLACES_API_KEY"), os.Getenv("REDIS_URL"))

	amadeusService := amadeus.NewAmadeusService(client, googlePlacesClient)
	placeService := google.NewGooglePlacesService(googlePlacesClient)

	amadeusUsecase := usecase.NewHotelUsecase(amadeusService)
	placeUsecase := usecase.NewGooglePlacesUsecase(placeService)

	router := mux.NewRouter()

	handlers.NewHotelHandler(router, amadeusUsecase)
	handlers.NewPlaceHandler(router, placeUsecase)

	// Apply CORS middleware to the router
	serverPort := "5100"
	logger.Printf("Starting server on port %s...\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, corsMiddleware(router)); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
