package places

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"microservices-travel-backend/internal/shared/api_provider/google/places/models"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const googlePlacesBaseURL = "https://places.googleapis.com/v1"

var ctx = context.Background()

type PlacesClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	Cache      *redis.Client
	Mutex      sync.Mutex
}

func NewPlacesClient(apiKey, redisAddr string) *PlacesClient {
	if strings.HasPrefix(redisAddr, "redis://") {
		redisAddr = strings.TrimPrefix(redisAddr, "redis://")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("❌ Failed to connect to Redis:", err)
	} else {
		fmt.Println("✅ Successfully connected to Redis at", redisAddr)
	}

	return &PlacesClient{
		APIKey:     apiKey,
		BaseURL:    googlePlacesBaseURL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Cache:      redisClient,
	}
}

func (c *PlacesClient) makeRequest(method, url string, body interface{}, headers map[string]string, result interface{}) error {
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
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", c.APIKey)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
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

func (pc *PlacesClient) GetPlacePhoto(placeID, photoID string, maxHeight, maxWidth int) (*models.PhotoResponse, error) {
	cacheKey := fmt.Sprintf("place_photo:%s:%s", placeID, photoID)

	cachedData, err := pc.Cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var result models.PhotoResponse
		if json.Unmarshal([]byte(cachedData), &result) == nil {
			fmt.Println("Cache hit: Returning data from Redis")
			return &result, nil
		}
	}

	fmt.Println("Cache miss: Fetching photo from Google Places API")

	url := fmt.Sprintf("%s/places/%s/photos/%s/media?key=%s&maxHeightPx=%d&maxWidthPx=%d&skipHttpRedirect=true",
		pc.BaseURL, placeID, photoID, pc.APIKey, maxHeight, maxWidth)

	var result models.PhotoResponse
	if err := pc.makeRequest("GET", url, nil, nil, &result); err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(result)
	pc.Cache.Set(ctx, cacheKey, jsonData, 10*time.Minute)

	return &result, nil
}

func (pc *PlacesClient) SearchPlaces(requestBody models.TextQueryRequest, fieldMask string) (*models.PlacesResponse, error) {
	cacheKey := fmt.Sprintf("place_search:%s", requestBody)

	// Attempt to retrieve cached data
	cachedData, err := pc.Cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var result models.PlacesResponse
		if json.Unmarshal([]byte(cachedData), &result) == nil {
			fmt.Println("Cache hit: Returning data from Redis")
			return &result, nil
		}
	}

	fmt.Println("Cache miss: Fetching data from Google Places API")

	// Define API URL
	url := fmt.Sprintf("%s/places:searchText", pc.BaseURL)

	headers := map[string]string{
		"X-Goog-FieldMask": fieldMask,
	}

	var result models.PlacesResponse
	if err := pc.makeRequest("POST", url, requestBody, headers, &result); err != nil {
		return nil, fmt.Errorf("failed to fetch places: %w", err)
	}

	// Store in cache
	jsonData, _ := json.Marshal(result)
	pc.Cache.Set(ctx, cacheKey, jsonData, 10*time.Minute)

	return &result, nil
}

func (pc *PlacesClient) GetPlaceDetails(placeID, fieldMask string) (*models.PlaceDetailsResponse, error) {
	cacheKey := fmt.Sprintf("place_details:%s", placeID)

	// Attempt to retrieve cached data
	cachedData, err := pc.Cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var result models.PlaceDetailsResponse
		if json.Unmarshal([]byte(cachedData), &result) == nil {
			fmt.Println("Cache hit: Returning data from Redis")
			return &result, nil
		}
	}

	fmt.Println("Cache miss: Fetching place details from Google Places API")

	// Define API URL
	url := fmt.Sprintf("%s/places/%s", pc.BaseURL, placeID)

	headers := map[string]string{
		"X-Goog-FieldMask": fieldMask,
	}

	var result models.PlaceDetailsResponse
	if err := pc.makeRequest("GET", url, nil, headers, &result); err != nil {
		return nil, fmt.Errorf("failed to fetch place details: %w", err)
	}

	// Store in cache
	jsonData, _ := json.Marshal(result)
	pc.Cache.Set(ctx, cacheKey, jsonData, 10*time.Minute)

	return &result, nil
}
