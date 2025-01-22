package repositories

import (
	"fmt"
	"microservices-travel-backend/internal/flight-booking/domain/models"
)

type DynamoDBBookingRepository struct {
}

func NewDynamoDBRepository() *DynamoDBBookingRepository {
	return &DynamoDBBookingRepository{}
}

func (r *DynamoDBBookingRepository) CreateFlight(flight *models.Flight) (*models.Flight, error) {
	fmt.Println("Flight created in DynamoDB")
	return flight, nil
}

func (r *DynamoDBBookingRepository) GetFlightByID(id string) (*models.Flight, error) {
	return &models.Flight{}, nil
}

func (r *DynamoDBBookingRepository) UpdateFlight(id string, flight *models.Flight) (*models.Flight, error) {
	return flight, nil
}

func (r *DynamoDBBookingRepository) DeleteFlight(id string) error {
	return nil
}
