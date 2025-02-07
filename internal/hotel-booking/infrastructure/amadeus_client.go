package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"microservices-travel-backend/internal/hotel-booking/domain"
	"sync"
	"time"

	// "microservices-travel-backend/config/hotel-booking/"
	"net/http"
)

var (
	amadeusToken       string
	amadeusTokenExpiry time.Time
	tokenMutex         sync.Mutex
	// configData         = config.LoadConfig()
)

// AmadeusClient implements HotelProvider
type AmadeusClient struct{}

// GetAmadeusAuthToken fetches authentication token
func GetAmadeusAuthToken() (string, error) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	if time.Now().Before(amadeusTokenExpiry) {
		return amadeusToken, nil
	}

	// url := "https://test.api.amadeus.com/v1/security/oauth2/token"
	// payload := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s",
	// 	configData.AmadeusClientID, configData.AmadeusSecret)

	url := "https://test.api.amadeus.com/v1/security/oauth2/token"
	payload := fmt.Sprintf("grant_type=client_credentials&client_id=ldK8AEKr1ryNBhfpEMNkux4CwjydYqrX&client_secret=8DJFOdD0t7pbUQSf")

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body) // Replaced ioutil.ReadAll with io.ReadAll
	var authRes struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	json.Unmarshal(body, &authRes)

	amadeusToken = authRes.AccessToken
	amadeusTokenExpiry = time.Now().Add(time.Duration(authRes.ExpiresIn) * time.Second)

	return amadeusToken, nil
}

// SearchHotels fetches hotels using Amadeus API

func (c *AmadeusClient) SearchHotels(cityCode string) ([]domain.Hotel, error) {
	token, err := GetAmadeusAuthToken()
	if err != nil {
		return nil, err
	}

	// url := fmt.Sprintf("%s/v1/reference-data/locations/hotels/by-city?cityCode=%s", config.AmadeusAPIBaseURL, cityCode)

	url := fmt.Sprintf("https://test.api.amadeus.com/v1/reference-data/locations/hotels/by-city?cityCode=%s", cityCode)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result struct {
		Data []struct {
			HotelID     string `json:"hotelId"`
			ChainCode   string `json:"chainCode"`
			IATACode    string `json:"iataCode"`
			CountryCode string `json:"countryCode"`
			Name        string `json:"name"`
			GeoCode     struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geoCode"`
		} `json:"data"`
	}

	json.Unmarshal(body, &result)

	hotels := make([]domain.Hotel, 0, len(result.Data))
	for _, h := range result.Data {
		hotels = append(hotels, domain.Hotel{
			HotelID:     h.HotelID,
			ChainCode:   h.ChainCode,
			IATACode:    h.IATACode,
			CountryCode: h.CountryCode,
			Name:        h.Name,
			Latitude:    h.GeoCode.Latitude,
			Longitude:   h.GeoCode.Longitude,
		})
	}

	return hotels, nil
}
