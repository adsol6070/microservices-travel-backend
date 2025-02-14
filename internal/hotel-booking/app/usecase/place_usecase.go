package usecase

import (
	"microservices-travel-backend/internal/hotel-booking/domain/google"
	"microservices-travel-backend/internal/shared/api_provider/google/places/models"
)

type GooglePlacesUsecase struct {
	service *google.GooglePlacesService
}

func NewGooglePlacesUsecase(service *google.GooglePlacesService) *GooglePlacesUsecase {
	return &GooglePlacesUsecase{
		service: service,
	}
}

func (u *GooglePlacesUsecase) SearchPlaces(requestBody models.TextQueryRequest) (*models.PlacesResponse, error) {
	places, err := u.service.SearchPlaces(requestBody)
	if err != nil {
		return nil, err
	}
	return places, nil
}

func (u *GooglePlacesUsecase) GetPlacePhoto(placeID, photoID string, maxHeight, maxWidth int) (*models.PhotoResponse, error) {
	placesPhotos, err := u.service.GetPlacePhoto(placeID, photoID, maxHeight, maxWidth)
	if err != nil {
		return nil, err
	}
	return placesPhotos, nil
}

func (u *GooglePlacesUsecase) GetPlaceDetail(placeID string) (*models.PlaceDetailsResponse, error) {
	placeDetails, err := u.service.GetPlaceDetail(placeID)
	if err != nil {
		return nil, err
	}
	return placeDetails, nil
}