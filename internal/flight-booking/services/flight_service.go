package services

import (
	"microservices-travel-backend/internal/flight-booking/domain/models"
	"microservices-travel-backend/internal/flight-booking/domain/ports"
)

type FlightService struct {
	db ports.FlightDB
}

func NewFlightService(db ports.FlightDB) *FlightService {
	return &FlightService{db: db}
}

func (h *FlightService) CreateFlight(flight *models.Flight) (*models.Flight, error) {
	createdFlight, err := h.db.CreateFlight(flight)
	if err != nil {
		return nil, err
	}
	return createdFlight, nil
}

func (h *FlightService) GetFlightByID(id string) (*models.Flight, error) {
	flight, err := h.db.GetFlightByID(id)
	if err != nil {
		return nil, err
	}
	return flight, nil
}

func (h *FlightService) UpdateFlight(id string, flight *models.Flight) (*models.Flight, error) {
	updatedFlight, err := h.db.UpdateFlight(id, flight)
	if err != nil {
		return nil, err
	}
	return updatedFlight, nil
}

func (h *FlightService) DeleteFlight(id string) error {
	err := h.db.DeleteFlight(id)
	if err != nil {
		return err
	}
	return nil
}
