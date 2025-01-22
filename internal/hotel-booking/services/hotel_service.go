package services

import (
	"microservices-travel-backend/internal/hotel-booking/domain/models"
	"microservices-travel-backend/internal/hotel-booking/domain/ports"
)

type HotelService struct {
	db ports.HotelDB
}

func NewHotelService(db ports.HotelDB) *HotelService {
	return &HotelService{db: db}
}

func (h *HotelService) CreateHotel(hotel *models.Hotel) (*models.Hotel, error) {
	createdHotel, err := h.db.CreateHotel(hotel)
	if err != nil {
		return nil, err
	}
	return createdHotel, nil
}

func (h *HotelService) GetHotelByID(id string) (*models.Hotel, error) {
	hotel, err := h.db.GetHotelByID(id)
	if err != nil {
		return nil, err
	}
	return hotel, nil
}

func (h *HotelService) UpdateHotel(id string, hotel *models.Hotel) (*models.Hotel, error) {
	updatedHotel, err := h.db.UpdateHotel(id, hotel)
	if err != nil {
		return nil, err
	}
	return updatedHotel, nil
}

func (h *HotelService) DeleteHotel(id string) error {
	err := h.db.DeleteHotel(id)
	if err != nil {
		return err
	}
	return nil
}
