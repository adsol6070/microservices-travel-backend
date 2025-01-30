package main

import (
	"log"
	"microservices-travel-backend/internal/hotel-booking/adapters/handlers"
	"microservices-travel-backend/internal/hotel-booking/adapters/hotel_provider"
	"microservices-travel-backend/internal/hotel-booking/adapters/repositories"
	"microservices-travel-backend/internal/hotel-booking/domain/mapper"
	"microservices-travel-backend/internal/hotel-booking/domain/ports"
	"microservices-travel-backend/internal/hotel-booking/services"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	repo, err := repositories.NewPostgresRepository()

	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	expedia_provider := hotel_provider.NewExpediaAdapter("dfgdfgfdg")

	providers := []ports.HotelProvider{expedia_provider}

	hotelMapper := mapper.NewHotelMapper()

	service := services.NewHotelService(repo, providers, hotelMapper)

	hotelHandler := handlers.NewHotelHandler(service)

	router := mux.NewRouter()

	hotelHandler.RegisterRoutes(router)

	port := ":5000"
	baseURL := os.Getenv("HOTEL_API_BASE_URL")
	log.Printf("Starting Hotel Booking Service on port %s with base URL: %s", port, baseURL)

	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
