package main

import (
	"log"
	"microservices-travel-backend/internal/flight-booking/adapters/handlers"
	"microservices-travel-backend/internal/flight-booking/adapters/repositories"
	"microservices-travel-backend/internal/flight-booking/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repo, err := repositories.NewPostgresRepository()

	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	service := services.NewFlightService(repo)

	flightHandler := handlers.NewFlightHandler(service)

	router := mux.NewRouter()

	flightHandler.RegisterRoutes(router)

	port := ":9090"
	log.Printf("Starting hotel-booking service on port %s...", port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
