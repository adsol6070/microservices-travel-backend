package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"microservices-travel-backend/internal/hotel-booking/domain/amadeus"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	amadeusToken       string
	amadeusTokenExpiry time.Time
	tokenMutex         sync.Mutex
	httpClient         = &http.Client{Timeout: 10 * time.Second}
	baseURL            = "https://test.api.amadeus.com"
)

func GetValidAmadeusToken() (string, error) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	if time.Now().Before(amadeusTokenExpiry) && amadeusToken != "" {
		return amadeusToken, nil
	}

	return GetAmadeusAuthToken()
}

func GetAmadeusAuthToken() (string, error) {
	url := fmt.Sprintf("%s/v1/security/oauth2/token", baseURL)
	payload := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s",
		"ldK8AEKr1ryNBhfpEMNkux4CwjydYqrX", "8DJFOdD0t7pbUQSf")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error response from Amadeus: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var authRes struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.Unmarshal(body, &authRes); err != nil {
		return "", fmt.Errorf("error parsing JSON: %w", err)
	}

	amadeusToken = authRes.AccessToken
	amadeusTokenExpiry = time.Now().Add(time.Duration(authRes.ExpiresIn) * time.Second)

	return amadeusToken, nil
}

func FetchHotelOffers(hotelIDs []string, adults int) ([]amadeus.HotelOffer, error) {
	token, err := GetValidAmadeusToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get Amadeus token: %w", err)
	}

	query := fmt.Sprintf("%s/v3/shopping/hotel-offers?hotelIds=%s&adults=%d", baseURL,
		strings.Join(hotelIDs, ","), adults)

	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response from Amadeus: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var result amadeus.HotelOffersResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return result.Data, nil
}
