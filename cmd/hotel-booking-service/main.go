package hotelbookingservice

import (
	"microservices-travel-backend/internal/hotel-booking/adapters/http"
	"microservices-travel-backend/internal/hotel-booking/adapters/repositories"
	"microservices-travel-backend/internal/hotel-booking/infrastructure"
	"microservices-travel-backend/internal/hotel-booking/services"
)

func main() {
	config.LoadConfig("dev")
	repo := repositories.NewDynamoDBRepository()

	service := services.NewHotelService(repo)

	hotelHandler := http.NewHotelHandler(service)
}
