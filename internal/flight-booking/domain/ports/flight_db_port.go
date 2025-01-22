package ports

import "microservices-travel-backend/internal/flight-booking/domain/models"

type FlightDB interface {
	CreateFlight(hotel *models.Flight) (*models.Flight, error)
	GetFlightByID(id string) (*models.Flight, error)
	UpdateFlight(id string, hotel *models.Flight) (*models.Flight, error)
	DeleteFlight(id string) error
}
