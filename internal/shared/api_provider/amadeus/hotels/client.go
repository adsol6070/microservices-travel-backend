package hotels

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/models"
)

const amadeusBaseURL = "https://test.api.amadeus.com"

type TokenManager struct {
	Token      string
	Expiry     time.Time
	Mutex      sync.Mutex
	APIKey     string
	APISecret  string
	HTTPClient *http.Client
}

type AmadeusClient struct {
	BaseURL      string
	TokenManager *TokenManager
}

func NewAmadeusClient(apiKey, secret string) *AmadeusClient {
	return &AmadeusClient{
		BaseURL: amadeusBaseURL,
		TokenManager: &TokenManager{
			APIKey:     apiKey,
			APISecret:  secret,
			HTTPClient: &http.Client{Timeout: 10 * time.Second},
		},
	}
}

func (tm *TokenManager) GetValidToken() (string, error) {
	tm.Mutex.Lock()
	defer tm.Mutex.Unlock()

	if time.Now().Before(tm.Expiry) && tm.Token != "" {
		return tm.Token, nil
	}
	return tm.fetchNewToken()
}

func (tm *TokenManager) fetchNewToken() (string, error) {
	url := fmt.Sprintf("%s/v1/security/oauth2/token", amadeusBaseURL)
	payload := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s",
		tm.APIKey, tm.APISecret)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := tm.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send token request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token, status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read token response body: %w", err)
	}

	var authRes struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.Unmarshal(body, &authRes); err != nil {
		return "", fmt.Errorf("failed to parse token response JSON: %w", err)
	}

	tm.Token = authRes.AccessToken
	tm.Expiry = time.Now().Add(time.Duration(authRes.ExpiresIn) * time.Second)

	return tm.Token, nil
}

func (c *AmadeusClient) FetchHotelOffers(hotelIDs []string, adults int) ([]models.HotelOffer, error) {
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	url := fmt.Sprintf("%s/v3/shopping/hotel-offers?hotelIds=%s&adults=%d", c.BaseURL, strings.Join(hotelIDs, ","), adults)

	var result models.HotelOffersResponse
	if err := c.makeRequest("GET", url, token, nil, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (c *AmadeusClient) HotelSearch(cityCode string) ([]models.HotelData, error) {
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	url := fmt.Sprintf("%s/v1/reference-data/locations/hotels/by-city?cityCode=%s", c.BaseURL, strings.ToUpper(cityCode))

	var result models.HotelListResponse
	if err := c.makeRequest("GET", url, token, nil, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (c *AmadeusClient) CreateHotelBooking(requestBody models.HotelBookingRequest) (*models.HotelOrderResponse, error) {
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	url := fmt.Sprintf("%s/v2/booking/hotel-orders", c.BaseURL)
	var result models.HotelOrderResponse
	if err := c.makeRequest("POST", url, token, requestBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *AmadeusClient) FetchHotelRatings(hotelIDs []string) (*models.HotelSentimentResponse, error) {
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	url := fmt.Sprintf("%s/v2/e-reputation/hotel-sentiments?hotelIds=%s", c.BaseURL, strings.Join(hotelIDs, ","))

	var result models.HotelSentimentResponse
	if err := c.makeRequest("GET", url, token, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *AmadeusClient) makeRequest(method, url, token string, body interface{}, result interface{}) error {
	var requestBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		requestBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.TokenManager.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("request failed, status: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("failed to parse response JSON: %w", err)
	}

	return nil
}

func (c *AmadeusClient) HotelNameAutoComplete(keyword string, subType string) (*models.HotelNameResponse, error) {
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	url := fmt.Sprintf("%s/v1/reference-data/locations/hotel?keyword=%s&subType=%s", c.BaseURL, keyword, subType)

	var result models.HotelNameResponse
	if err := c.makeRequest("GET", url, token, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
