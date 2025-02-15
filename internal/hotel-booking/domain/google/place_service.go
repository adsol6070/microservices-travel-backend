package google

import (
	"microservices-travel-backend/internal/shared/api_provider/google/places"
	"microservices-travel-backend/internal/shared/api_provider/google/places/googlePlaceModels"
)

type GooglePlacesService struct {
	client *places.PlacesClient
}

func NewGooglePlacesService(client *places.PlacesClient) *GooglePlacesService {
	return &GooglePlacesService{
		client: client,
	}
}

func (g *GooglePlacesService) SearchPlaces(requestBody googlePlaceModels.TextQueryRequest) (*googlePlaceModels.PlacesResponse, error) {
	places, err := g.client.SearchPlaces(requestBody, "places.displayName,places.id,places.photos")
	if err != nil {
		return nil, err
	}
	return places, nil
}

func (g *GooglePlacesService) GetPlacePhoto(photoName string, maxHeight, maxWidth int) (*googlePlaceModels.PhotoResponse, error) {
	placesPhotos, err := g.client.GetPlacePhoto(photoName, maxHeight, maxWidth)
	if err != nil {
		return nil, err
	}
	return placesPhotos, nil
}

func (g *GooglePlacesService) GetPlaceDetail(placeID string) (*googlePlaceModels.PlaceDetailsResponse, error) {

	placeDetail, err := g.client.GetPlaceDetails(placeID, "*")
	if err != nil {
		return nil, err
	}
	return placeDetail, nil
}
