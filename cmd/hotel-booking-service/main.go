package main

import (
	"log"
	"microservices-travel-backend/internal/hotel-booking/app/usecase"
	"microservices-travel-backend/internal/hotel-booking/domain/amadeus"
	"microservices-travel-backend/internal/hotel-booking/infrastructure/handlers"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "hotel-service: ", log.LstdFlags|log.Lshortfile)

	client := hotels.NewAmadeusClient(os.Getenv("AMADEUS_API_KEY"), os.Getenv("AMADEUS_SECRET"), os.Getenv("REDIS_URL"))

	amadeusService := amadeus.NewAmadeusService(client)

	amadeusUsecase := usecase.NewHotelUsecase(amadeusService) 

	router := mux.NewRouter()

	handlers.NewHotelHandler(router, amadeusUsecase)
	// Set Port and Start Server
	serverPort := "5100"
	logger.Printf("Starting server on port %s...\n", "5100")
	if err := http.ListenAndServe(":"+serverPort, router); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}

}
