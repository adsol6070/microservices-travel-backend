package main

import (
	"log"
	"microservices-travel-backend/internal/flight-booking/adapters/handlers"
	"microservices-travel-backend/internal/flight-booking/adapters/repositories"
	// "microservices-travel-backend/internal/flight-booking/infrastructure"
	"microservices-travel-backend/internal/flight-booking/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// config.LoadConfig("dev")
	repo := repositories.NewDynamoDBRepository()

	service := services.NewFlightService(repo)

	flightHandler := handlers.NewFlightHandler(service)

	router := mux.NewRouter()

	flightHandler.RegisterRoutes(router)

	port := ":9090"
	log.Printf("Starting hotel-booking service on port %s...", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
