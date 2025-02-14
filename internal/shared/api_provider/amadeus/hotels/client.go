package hotels

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/models"

	"github.com/redis/go-redis/v9"
)

const amadeusBaseURL = "https://test.api.amadeus.com"

var ctx = context.Background()

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
	Cache        *redis.Client
}

func NewAmadeusClient(apiKey, secret string, redisAddr string) *AmadeusClient {
	// Ensure Redis address is formatted correctly
	if strings.HasPrefix(redisAddr, "redis://") {
		redisAddr = strings.TrimPrefix(redisAddr, "redis://")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr, // Correct format: "host:port"
		Password: "",
		DB:       0,
	})

	// Check Redis connection
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("❌ Failed to connect to Redis:", err)
	} else {
		fmt.Println("✅ Successfully connected to Redis at", redisAddr)
	}

	return &AmadeusClient{
		BaseURL: amadeusBaseURL,
		TokenManager: &TokenManager{
			APIKey:     apiKey,
			APISecret:  secret,
			HTTPClient: &http.Client{Timeout: 10 * time.Second},
		},
		Cache: redisClient,
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
	const batchSize = 50
	var allOffers []models.HotelOffer

	for i := 0; i < len(hotelIDs); i += batchSize {
		end := i + batchSize
		if end > len(hotelIDs) {
			end = len(hotelIDs)
		}

		batch := hotelIDs[i:end]
		offers, err := c.fetchBatchHotelOffers(batch, adults)
		if err != nil {
			log.Printf("WARNING: Skipping failed batch %v: %v", batch, err)
			continue
		}

		allOffers = append(allOffers, offers...)
	}

	return allOffers, nil
}

// Fetch offers for a batch of up to 50 hotel IDs
func (c *AmadeusClient) fetchBatchHotelOffers(hotelIDs []string, adults int) ([]models.HotelOffer, error) {
	cacheKey := fmt.Sprintf("hotel_offers:%s:adults:%d", strings.Join(hotelIDs, ","), adults)

	// Try fetching from Redis cache first
	cachedData, err := c.Cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var offers []models.HotelOffer
		if json.Unmarshal([]byte(cachedData), &offers) == nil {
			fmt.Println("Cache hit: Returning data from Redis")
			return offers, nil
		}
	}

	fmt.Println("Cache miss: Fetching data from Amadeus API")

	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	// Construct API request for up to 50 hotel IDs
	url := fmt.Sprintf("%s/v3/shopping/hotel-offers?hotelIds=%s&adults=%d",
		c.BaseURL, strings.Join(hotelIDs, ","), adults)

	// Fetch raw response body
	respBody, err := c.makeRequest("GET", url, token, nil)
	if err != nil {
		log.Printf("ERROR: Failed to fetch hotel offers: %v", err)
		return nil, err // Return error if the request fails
	}

	// Parse the response
	var result models.HotelOffersResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		log.Printf("ERROR: Failed to parse response JSON: %v", err)
		return nil, err // Return error if JSON parsing fails
	}

	// ✅ Fix: Check if Data is empty instead of checking for a "Status" field
	if len(result.Data) == 0 {
		log.Printf("WARNING: API returned empty data for hotels: %v", hotelIDs)
		return nil, fmt.Errorf("API returned empty data for the given hotels")
	}

	// Store in cache if we got any successful results
	jsonData, _ := json.Marshal(result.Data)
	c.Cache.Set(ctx, cacheKey, jsonData, 1*time.Minute)

	return result.Data, nil
}

func (c *AmadeusClient) HotelSearch(cityCode string) ([]models.HotelData, error) {
	cacheKey := fmt.Sprintf("hotel_search:%s", strings.ToUpper(cityCode))

	// Check cache
	cachedData, err := c.Cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var hotels []models.HotelData
		if json.Unmarshal([]byte(cachedData), &hotels) == nil {
			fmt.Println("Cache hit: Returning data from Redis")
			return hotels, nil
		}
	}

	fmt.Println("Cache miss: Fetching data from Amadeus API")

	// Fetch token
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	// Prepare request URL
	url := fmt.Sprintf("%s/v1/reference-data/locations/hotels/by-city?cityCode=%s", c.BaseURL, strings.ToUpper(cityCode))

	// Make API request and get raw response
	respBody, err := c.makeRequest("GET", url, token, nil)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var result models.HotelListResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse hotel search response: %w", err)
	}

	// Cache result
	jsonData, _ := json.Marshal(result.Data)
	c.Cache.Set(ctx, cacheKey, jsonData, 10*time.Minute)

	return result.Data, nil
}

func (c *AmadeusClient) CreateHotelBooking(requestBody models.HotelBookingRequest) (*models.HotelOrderResponse, error) {
	// Get a valid Amadeus token
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	// Prepare API URL
	url := fmt.Sprintf("%s/v2/booking/hotel-orders", c.BaseURL)

	// Make API request and get raw response
	respBody, err := c.makeRequest("POST", url, token, requestBody)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var result models.HotelOrderResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse hotel booking response: %w", err)
	}

	return &result, nil
}

func (c *AmadeusClient) FetchHotelRatings(hotelIDs []string) (*models.HotelSentimentResponse, error) {
	// Get a valid Amadeus token
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	// Construct API URL with hotel IDs
	url := fmt.Sprintf("%s/v2/e-reputation/hotel-sentiments?hotelIds=%s", c.BaseURL, strings.Join(hotelIDs, ","))

	// Make API request and get raw response
	respBody, err := c.makeRequest("GET", url, token, nil)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var result models.HotelSentimentResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse hotel sentiment response: %w", err)
	}

	return &result, nil
}

func (c *AmadeusClient) makeRequest(method, url, token string, body interface{}) ([]byte, error) {
	var requestBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		requestBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.TokenManager.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-successful responses
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("request failed, status: %s, body: %s", resp.Status, string(respBody))
	}

	// Return raw response body to be processed outside
	return respBody, nil
}

func (c *AmadeusClient) HotelNameAutoComplete(keyword string, subType string) (*models.HotelNameResponse, error) {
	// Get a valid Amadeus token
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	// Construct API URL
	url := fmt.Sprintf("%s/v1/reference-data/locations/hotel?keyword=%s&subType=%s", c.BaseURL, keyword, subType)

	// Make API request and get raw response
	respBody, err := c.makeRequest("GET", url, token, nil)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var result models.HotelNameResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse hotel name response: %w", err)
	}

	return &result, nil
}
