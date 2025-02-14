package google

import (
	"microservices-travel-backend/internal/shared/api_provider/google/places"
	"microservices-travel-backend/internal/shared/api_provider/google/places/models"
)

type GooglePlacesService struct {
	client *places.PlacesClient
}

func NewGooglePlacesService(client *places.PlacesClient) *GooglePlacesService {
	return &GooglePlacesService{
		client: client,
	}
}

func (g *GooglePlacesService) SearchPlaces(requestBody models.TextQueryRequest) (*models.PlacesResponse, error) {
	places, err := g.client.SearchPlaces(requestBody, "places.displayName,places.id,places.photos")
	if err != nil {
		return nil, err
	}
	return places, nil
}

func (g *GooglePlacesService) GetPlacePhoto(placeID, photoID string, maxHeight, maxWidth int) (*models.PhotoResponse, error) {
	placesPhotos, err := g.client.GetPlacePhoto(placeID, photoID, maxHeight, maxWidth)
	if err != nil {
		return nil, err
	}
	return placesPhotos, nil
}

func (g *GooglePlacesService) GetPlaceDetail(placeID string) (*models.PlaceDetailsResponse, error) {

	placeDetail, err := g.client.GetPlaceDetails(placeID, "*")
	if err != nil {
		return nil, err
	}
	return placeDetail, nil
}
