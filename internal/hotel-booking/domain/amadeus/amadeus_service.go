package amadeus

import (
	"encoding/json"
	"errors"
	"log"
	"microservices-travel-backend/internal/hotel-booking/app/dto/request"
	"microservices-travel-backend/internal/hotel-booking/app/dto/response"
	"microservices-travel-backend/internal/hotel-booking/domain/models"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/amadeusHotelModels"
	"microservices-travel-backend/internal/shared/api_provider/google/places"
	"microservices-travel-backend/internal/shared/api_provider/google/places/googlePlaceModels"
	"strconv"
	"sync"
)

type AmadeusService struct {
	client *hotels.AmadeusClient
	places *places.PlacesClient
}

func NewAmadeusService(client *hotels.AmadeusClient, places *places.PlacesClient) *AmadeusService {
	return &AmadeusService{
		client: client,
		places: places,
	}
}

func (a *AmadeusService) FetchHotelOffers(hotelIDs []string, adults int) ([]amadeusHotelModels.HotelOffer, error) {
	offers, err := a.client.FetchHotelOffers(hotelIDs, adults)
	if err != nil {
		return nil, err
	}
	return offers, nil
}

func (a *AmadeusService) SearchHotels(req request.HotelSearchRequest) ([]models.EnrichedHotelOffer, error) {
	hotels, err := a.client.FetchHotelsByCity(req.CityCode)
	if err != nil {
		log.Println("ERROR: Failed to fetch hotels -", err)
		return nil, err
	}

	if len(hotels) == 0 {
		log.Println("WARN: No hotels found for the given city code")
		return nil, errors.New("no hotels found for the given city code")
	}

	// Extract hotel IDs
	hotelIDs := extractHotelIDs(hotels)
	log.Printf("INFO: Extracted %d hotel IDs", len(hotelIDs))

	if len(hotelIDs) == 0 {
		log.Println("WARN: No valid hotels found")
		return nil, errors.New("no valid hotels found")
	}

	// Fetch hotel offers using extracted hotel IDs
	hotelOffers, err := a.client.FetchHotelOffers(hotelIDs, req.Persons)
	if err != nil {
		log.Println("ERROR: Failed to fetch hotel offers -", err)
		return nil, err
	}

	if len(hotelOffers) == 0 {
		log.Println("WARN: No hotel offers available")
		return nil, errors.New("no hotel offers available")
	}

	// Filter hotel offers based on search criteria
	filteredHotels := filterHotelsByRequest(hotelOffers, req)

	if len(filteredHotels) == 0 {
		log.Println("WARN: No hotels matched the search criteria")
		return nil, errors.New("no matching hotels found")
	}

	var enrichedHotels []models.EnrichedHotelOffer

	for _, hotel := range filteredHotels {

		textQuery := googlePlaceModels.TextQueryRequest{
			TextQuery: hotel.Hotel.Name,
			LocationBias: googlePlaceModels.LocationBias{
				Circle: googlePlaceModels.Circle{
					Center: googlePlaceModels.Coordinates{
						Latitude:  hotel.Hotel.Latitude,
						Longitude: hotel.Hotel.Longitude,
					},
					Radius: 5000,
				},
			},
		}

		placeDetails, err := a.places.SearchPlaces(
			textQuery,
			"places.id,places.displayName,places.rating,places.photos",
		)

		if err != nil || len(placeDetails.Places) == 0 {
			log.Printf("WARN: No Google Places data found for %s", hotel.Hotel.Name)
			continue
		}

		// Extract first place details
		place := placeDetails.Places[0]
		var photoURL string

		// Fetch photo if available
		if len(place.Photos) > 0 {
			photoResp, err := a.places.GetPlacePhoto(place.Photos[0].Name, 400, 400)
			if err == nil {
				photoURL = photoResp.PhotoURI
			}
		}

		enrichedHotel := models.EnrichedHotelOffer{
			HotelID:           hotel.Hotel.HotelID,
			Name:              hotel.Hotel.Name,
			Latitude:          hotel.Hotel.Latitude,
			Longitude:         hotel.Hotel.Longitude,
			PhotoURL:          photoURL,
			Offers:            hotel.Offers,
			GooglePlaceID:     place.ID,
			Rating:            place.Rating,
			GoogleDisplayName: place.DisplayName.Text,
		}

		enrichedHotels = append(enrichedHotels, enrichedHotel)
	}

	if len(enrichedHotels) == 0 {
		log.Println("WARN: No enriched hotels found")
		return nil, errors.New("no enriched hotels found")
	}

	return enrichedHotels, nil
}

func extractHotelIDs(hotels []amadeusHotelModels.Hotel) []string {
	var hotelIDs []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, hotel := range hotels {
		wg.Add(1)
		go func(h amadeusHotelModels.Hotel) {
			defer wg.Done()
			if h.HotelID != "" {
				mu.Lock()
				hotelIDs = append(hotelIDs, h.HotelID)
				mu.Unlock()
			}
		}(hotel)
	}

	wg.Wait()
	return hotelIDs
}

func filterHotelsByRequest(hotelOffers []amadeusHotelModels.HotelOffer, req request.HotelSearchRequest) []amadeusHotelModels.HotelOffer {
	var filteredHotels []amadeusHotelModels.HotelOffer

	for _, hotelOffer := range hotelOffers {
		if !hotelOffer.Available {
			continue
		}

		for _, offer := range hotelOffer.Offers {
			if offer.CheckIn == req.CheckInDate && offer.CheckOut == req.CheckOutDate {
				if req.Rooms > 0 && req.Rooms == len(hotelOffer.Offers) {
					filteredHotels = append(filteredHotels, hotelOffer)
					break
				}
			}
		}
	}

	return filteredHotels
}

func (a *AmadeusService) extractPhotoReferences(photos []googlePlaceModels.PhotoDetail) []string {
	var photoURLs []string

	for _, photo := range photos {
		photoResp, err := a.places.GetPlacePhoto(photo.Name, photo.HeightPx, photo.WidthPx)
		if err == nil {
			photoURLs = append(photoURLs, photoResp.PhotoURI)
		} else {
			log.Println("ERROR: Failed to fetch photo -", err)
		}
	}

	return photoURLs
}

func (a *AmadeusService) HotelDetails(req request.HotelDetailsRequest) (response.HotelDetails, error) {
	log.Println("INFO: Fetching hotel details for", req.HotelID)
	hotelOffers, err := a.client.FetchHotelOffers([]string{req.HotelID}, 1)
	if err != nil {
		log.Println("ERROR: Failed to fetch hotel offers -", err)
		return response.HotelDetails{}, err
	}

	if len(hotelOffers) == 0 {
		log.Println("ERROR: No hotel offers available for HotelID:", req.HotelID)
		return response.HotelDetails{}, errors.New("no hotel offers available")
	}

	offersJSON, err := json.MarshalIndent(hotelOffers, "", "  ")
	if err != nil {
		log.Println("ERROR: Failed to marshal hotel offers -", err)
	} else {
		log.Println("INFO: Hotel offers fetched successfully for HotelID:", req.HotelID, "Offers:", string(offersJSON))
	}

	log.Println("INFO: Fetching Google Place details for PlaceID:", req.GooglePlaceID)
	googlePlaceDetails, err := a.places.GetPlaceDetails(req.GooglePlaceID, "*")
	if err != nil {
		log.Println("ERROR: Failed to fetch Google Place details -", err)
		return response.HotelDetails{}, err
	}

	googlePlaceDetailsJSON, err := json.MarshalIndent(googlePlaceDetails, "", "  ")
	if err != nil {
		log.Println("ERROR: Failed to marshal googlePlace details -", err)
	} else {
		log.Println("INFO: GooglePlaceDetails fetched successfully for GooglePlaceID:", req.GooglePlaceID, "GooglePlaceDetails:", string(googlePlaceDetailsJSON))
	}

	offer := hotelOffers[0]
	if len(offer.Offers) == 0 {
		return response.HotelDetails{}, errors.New("no offers available for the selected hotel")
	}

	offerDetails := offer.Offers[0]

	basePrice, _ := strconv.ParseFloat(offerDetails.Price.Base, 64)
	totalPrice, _ := strconv.ParseFloat(offerDetails.Price.Total, 64)

	hotel := response.HotelDetails{
		HotelID:      offer.Hotel.HotelID,
		HotelName:    googlePlaceDetails.Name,
		CheckInDate:  offerDetails.CheckIn,
		CheckOutDate: offerDetails.CheckOut,
		Price: response.PriceDetails{
			Currency: offerDetails.Price.Currency,
			Base:     basePrice,
			Total:    totalPrice,
		},
		Location: response.LocationDetails{
			CityCode: *offer.Hotel.CityCode,
			Address: response.Address{
				Street:     googlePlaceDetails.AddressDescriptor.Street,
				City:       googlePlaceDetails.AddressDescriptor.City,
				State:      googlePlaceDetails.AddressDescriptor.State,
				PostalCode: googlePlaceDetails.AddressDescriptor.PostalCode,
				Country:    googlePlaceDetails.AddressDescriptor.Country,
			},
			Coordinates: response.Coordinates{
				Latitude:  *googlePlaceDetails.Location.Latitude,
				Longitude: *googlePlaceDetails.Location.Longitude,
			},
			GoogleMapsURI: googlePlaceDetails.GoogleMapsURI,
			PlaceID:       googlePlaceDetails.ID,
		},
		Contact: response.ContactDetails{
			PhoneNumber: googlePlaceDetails.InternationalPhoneNumber,
			Website:     googlePlaceDetails.WebsiteURI,
		},
		Rating: response.RatingDetails{
			Stars:            0, // Google Places API does not return star ratings
			UserReviewsCount: googlePlaceDetails.UserRatingCount,
			AverageRating:    googlePlaceDetails.Rating,
		},
		Photos: a.extractPhotoReferences(googlePlaceDetails.Photos), // Extracting photo URLs
		Accessibility: response.AccessibilityOptions{
			WheelchairAccessible: googlePlaceDetails.AccessibilityOptions.WheelchairAccessible,
		},
		Payment: response.PaymentOptions{
			CardsAccepted:  googlePlaceDetails.PaymentOptions.CardsAccepted,
			CashAccepted:   googlePlaceDetails.PaymentOptions.CashAccepted,
			DigitalWallets: googlePlaceDetails.PaymentOptions.DigitalWallets,
		},
		Policies: response.HotelPolicies{
			CheckInTime:    "",
			CheckOutTime:   "",
			SmokingAllowed: false,
			MinCheckInAge:  18,
		},
	}

	return hotel, nil
}

func (a *AmadeusService) CreateHotelBooking(requestBody amadeusHotelModels.HotelBookingReq) (*amadeusHotelModels.HotelOrderResponseData, error) {
	booking, err := a.client.CreateHotelBooking(requestBody)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (a *AmadeusService) FetchHotelRatings(hotelIDs []string) (*amadeusHotelModels.HotelSentimentResponse, error) {
	ratings, err := a.client.FetchHotelRatings(hotelIDs)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (a *AmadeusService) HotelNameAutoComplete(keyword string, subtype string) (*amadeusHotelModels.HotelNameResponse, error) {
	hotels, err := a.client.HotelNameAutoComplete(keyword, subtype)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
