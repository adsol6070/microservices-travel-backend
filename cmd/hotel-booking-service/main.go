package main

import (
	"log"
	"microservices-travel-backend/internal/hotel-booking/adapters/handlers"
	"microservices-travel-backend/internal/hotel-booking/adapters/repositories"
	// "microservices-travel-backend/internal/hotel-booking/infrastructure"
	"microservices-travel-backend/internal/hotel-booking/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// config.LoadConfig("dev")
	repo := repositories.NewDynamoDBRepository()
	// repo := repositories.PostgresBookingRepository()

	service := services.NewHotelService(repo)

	hotelHandler := handlers.NewHotelHandler(service)

	router := mux.NewRouter()

	hotelHandler.RegisterRoutes(router)

	port := ":5000"
	log.Printf("Starting hotel-booking service on port %s...", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
