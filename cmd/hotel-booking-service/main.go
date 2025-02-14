package main

import (
	"log"
	"microservices-travel-backend/internal/hotel-booking/app/usecase"
	"microservices-travel-backend/internal/hotel-booking/domain/amadeus"
	"microservices-travel-backend/internal/hotel-booking/domain/google"
	"microservices-travel-backend/internal/hotel-booking/infrastructure/handlers"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels"
	"microservices-travel-backend/internal/shared/api_provider/google/places"

	"net/http"
	"os"
 
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "hotel-service: ", log.LstdFlags|log.Lshortfile)

	client := hotels.NewAmadeusClient(os.Getenv("AMADEUS_API_KEY"), os.Getenv("AMADEUS_SECRET"), os.Getenv("REDIS_URL"))

	googlePlacesClient := places.NewPlacesClient(
		os.Getenv("GOOGLE_PLACES_API_KEY"),
		os.Getenv("REDIS_URL"),
	)

	amadeusService := amadeus.NewAmadeusService(client)

	placeService := google.NewGooglePlacesService(googlePlacesClient)

	amadeusUsecase := usecase.NewHotelUsecase(amadeusService) 

	placeUsecase := usecase.NewGooglePlacesUsecase(placeService) 

	router := mux.NewRouter()

	handlers.NewHotelHandler(router, amadeusUsecase)
	handlers.NewPlaceHandler(router, placeUsecase)
	// Set Port and Start Server
	serverPort := "5100"
	logger.Printf("Starting server on port %s...\n", "5100")
	if err := http.ListenAndServe(":"+serverPort, router); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}

}
