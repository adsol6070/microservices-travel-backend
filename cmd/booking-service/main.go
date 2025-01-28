package main

import (
	"log"
	"microservices-travel-backend/internal/booking-service/adapters/handlers"
	"microservices-travel-backend/internal/booking-service/adapters/repositories"
	"microservices-travel-backend/internal/booking-service/infrastructure"
	"microservices-travel-backend/internal/booking-service/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig("dev")
	repo := repositories.NewDynamoDBRepository()

	service := services.BookingService(repo)

	bookingHandler := handlers.NewBookingHandler(service)

	router := mux.NewRouter()

	bookingHandler.RegisterRoutes(router)

	port := ":6000"
	log.Printf("Starting booking service port %s...", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
