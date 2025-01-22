package hotelbookingservice

import (
	"microservices-travel-backend/internal/flight-booking/adapters/api/"
	"microservices-travel-backend/internal/flight-booking/adapters/repositories"
	"microservices-travel-backend/internal/flight-booking/infrastructure"
	"microservices-travel-backend/internal/flight-booking/services"
)

func main() {
	config.LoadConfig("dev")
	repo := repositories.NewDynamoDBRepository()

	service := services.NewFlightService(repo)

	hotelHandler := http.NewFlightHandler(service)
}
