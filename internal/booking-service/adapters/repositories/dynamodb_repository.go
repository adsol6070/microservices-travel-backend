package repositories

import (
	"fmt"
	"microservices-travel-backend/internal/hotel-booking/domain/models"
)

type DynamoDBBookingRepository struct {
}

func NewDynamoDBRepository() *DynamoDBBookingRepository {
	return &DynamoDBBookingRepository{}
}

func (r *DynamoDBBookingRepository) CreateHotel(hotel *models.Hotel) (*models.Hotel, error) {
	fmt.Println("Hotel created in DynamoDB")
	return hotel, nil
}

func (r *DynamoDBBookingRepository) GetHotelByID(id string) (*models.Hotel, error) {
	return &models.Hotel{}, nil
}

func (r *DynamoDBBookingRepository) UpdateHotel(id string, hotel *models.Hotel) (*models.Hotel, error) {
	return hotel, nil
}

func (r *DynamoDBBookingRepository) DeleteHotel(id string) error {
	return nil
}
