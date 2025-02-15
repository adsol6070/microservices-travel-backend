package usecase

import (
	"microservices-travel-backend/internal/hotel-booking/domain/google"
	"microservices-travel-backend/internal/shared/api_provider/google/places/googlePlaceModels"
)

type GooglePlacesUsecase struct {
	service *google.GooglePlacesService
}

func NewGooglePlacesUsecase(service *google.GooglePlacesService) *GooglePlacesUsecase {
	return &GooglePlacesUsecase{
		service: service,
	}
}

func (u *GooglePlacesUsecase) SearchPlaces(requestBody googlePlaceModels.TextQueryRequest) (*googlePlaceModels.PlacesResponse, error) {
	places, err := u.service.SearchPlaces(requestBody)
	if err != nil {
		return nil, err
	}
	return places, nil
}

func (u *GooglePlacesUsecase) GetPlacePhoto(photoName string, maxHeight, maxWidth int) (*googlePlaceModels.PhotoResponse, error) {
	placesPhotos, err := u.service.GetPlacePhoto(photoName, maxHeight, maxWidth)
	if err != nil {
		return nil, err
	}
	return placesPhotos, nil
}

func (u *GooglePlacesUsecase) GetPlaceDetail(placeID string) (*googlePlaceModels.PlaceDetailsResponse, error) {
	placeDetails, err := u.service.GetPlaceDetail(placeID)
	if err != nil {
		return nil, err
	}
	return placeDetails, nil
}