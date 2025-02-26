package hotels

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/amadeusHotelModels"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels/util"

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
	fetchTokenUrl := util.BuildURL(amadeusBaseURL, "/v1/security/oauth2/token", nil)

	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", tm.APIKey)
	formData.Set("client_secret", tm.APISecret)

	payload := formData.Encode()

	req, err := http.NewRequest("POST", fetchTokenUrl, strings.NewReader(payload))
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

func getCachedData[T any](cache *redis.Client, cacheKey string, fetchFunc func() (T, error), cacheDuration time.Duration) (T, error) {
	var zeroValue T

	cachedData, err := cache.Get(ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		var data T
		if json.Unmarshal([]byte(cachedData), &data) == nil {
			log.Println("Cache hit: Returning data from Redis")
			return data, nil
		}
	}

	log.Println("Cache miss: Fetching data from Amadeus API")
	data, err := fetchFunc()
	if err != nil {
		return zeroValue, err
	}

	jsonData, _ := json.Marshal(data)
	cache.Set(ctx, cacheKey, jsonData, cacheDuration)

	return data, nil
}

func (c *AmadeusClient) FetchHotelOffers(hotelIDs []string, adults int) ([]amadeusHotelModels.HotelOffer, error) {
	const batchSize = 50
	var allOffers []amadeusHotelModels.HotelOffer

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

func (c *AmadeusClient) fetchBatchHotelOffers(hotelIDs []string, adults int) ([]amadeusHotelModels.HotelOffer, error) {
	cacheKey := fmt.Sprintf("hotel_offers:%s:adults:%d", strings.Join(hotelIDs, ","), adults)

	fetchFunc := func() ([]amadeusHotelModels.HotelOffer, error) {
		token, err := c.TokenManager.GetValidToken()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
		}

		params := map[string]string{
			"hotelIds": strings.Join(hotelIDs, ","),
			"adults":   strconv.Itoa(adults),
		}
		url := util.BuildURL(c.BaseURL, "/v3/shopping/hotel-offers", params)

		respBody, err := c.makeRequest("GET", url, token, nil)
		if err != nil {
			log.Printf("ERROR: Failed to fetch hotel offers: %v", err)
			return nil, err
		}

		var result amadeusHotelModels.HotelOffersResp
		if err := json.Unmarshal(respBody, &result); err != nil {
			log.Printf("ERROR: Failed to parse response JSON: %v", err)
			return nil, err
		}

		if len(result.Data) == 0 {
			log.Printf("WARNING: API returned empty data for hotels: %v", hotelIDs)
			return nil, fmt.Errorf("API returned empty data for the given hotels")
		}

		return result.Data, nil
	}

	offers, err := getCachedData(c.Cache, cacheKey, fetchFunc, 1*time.Minute)
	if err != nil {
		return nil, err
	}

	return offers, nil
}

func (c *AmadeusClient) FetchHotelsByGeocode(latitude, longitude string) ([]amadeusHotelModels.Hotel, error) {
	if latitude == "" || longitude == "" {
		return nil, errors.New("latitude and longitude cannot be empty")
	}

	cacheKey := fmt.Sprintf("hotels_by_geocode:%s:%s", latitude, longitude)

	fetchFunc := func() ([]amadeusHotelModels.Hotel, error) {
		params := map[string]string{
			"latitude":  latitude,
			"longitude": longitude,
		}

		token, err := c.TokenManager.GetValidToken()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
		}

		url := util.BuildURL(c.BaseURL, "/v1/reference-data/locations/hotels/by-geocode", params)

		respBody, err := c.makeRequest("GET", url, token, nil)
		if err != nil {
			return nil, err
		}

		var result amadeusHotelModels.HotelsByGeoCodeResp
		if err := json.Unmarshal(respBody, &result); err != nil {
			return nil, fmt.Errorf("failed to parse hotel search response: %w", err)
		}

		return result.Data, nil
	}

	hotels, err := getCachedData(c.Cache, cacheKey, fetchFunc, 1*time.Minute)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func (c *AmadeusClient) FetchHotelsByID(hotelIDs []string) ([]amadeusHotelModels.Hotel, error) {
	if len(hotelIDs) == 0 {
		return nil, errors.New("hotel IDs cannot be empty")
	}

	cacheKey := fmt.Sprintf("hotels_by_id:%s", strings.Join(hotelIDs, ","))

	fetchFunc := func() ([]amadeusHotelModels.Hotel, error) {
		hotelIDParam := strings.Join(hotelIDs, ",")
		params := map[string]string{
			"hotelIds": hotelIDParam,
		}

		token, err := c.TokenManager.GetValidToken()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
		}

		url := util.BuildURL(c.BaseURL, "/v1/reference-data/locations/hotels/by-hotels", params)

		respBody, err := c.makeRequest("GET", url, token, nil)
		if err != nil {
			return nil, err
		}

		var result amadeusHotelModels.HotelsByIDResp
		if err := json.Unmarshal(respBody, &result); err != nil {
			return nil, fmt.Errorf("failed to parse hotel search response: %w", err)
		}

		return result.Data, nil
	}

	hotels, err := getCachedData(c.Cache, cacheKey, fetchFunc, 1*time.Minute)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func (c *AmadeusClient) FetchHotelsByCity(cityCode string) ([]amadeusHotelModels.Hotel, error) {
	if cityCode == "" {
		return nil, errors.New("city code cannot be empty")
	}

	cacheKey := fmt.Sprintf("hotels_by_city:%s", strings.ToUpper(cityCode))

	fetchFunc := func() ([]amadeusHotelModels.Hotel, error) {
		params := map[string]string{
			"cityCode": strings.ToUpper(cityCode),
		}

		token, err := c.TokenManager.GetValidToken()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
		}

		url := util.BuildURL(c.BaseURL, "/v1/reference-data/locations/hotels/by-city", params)

		respBody, err := c.makeRequest("GET", url, token, nil)
		if err != nil {
			return nil, err
		}

		var result amadeusHotelModels.HotelsByCityResp
		if err := json.Unmarshal(respBody, &result); err != nil {
			log.Printf("DEBUG: Raw response body: %s", string(respBody))
			return nil, fmt.Errorf("failed to parse hotel search response: %w", err)
		}

		for i := range result.Data {
			result.Data[i].DupeID = json.RawMessage([]byte(`"` + result.Data[i].GetDupeID() + `"`))
		}

		return result.Data, nil
	}

	hotels, err := getCachedData(c.Cache, cacheKey, fetchFunc, 1*time.Minute)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func (c *AmadeusClient) CreateHotelBooking(requestBody amadeusHotelModels.AmadeusHotelOrderRequest) (*amadeusHotelModels.AmadeusBookingData, error) {
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	url := util.BuildURL(c.BaseURL, "/v2/booking/hotel-orders", map[string]string{})
	respBody, err := c.makeRequest("POST", url, token, requestBody)
	if err != nil {
		return nil, err
	}

	var result amadeusHotelModels.AmadeusHotelOrderResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse hotel booking response: %w", err)
	}

	return &result.Data, nil
}

func (c *AmadeusClient) FetchHotelRatings(hotelIDs []string) (*amadeusHotelModels.HotelSentimentResponse, error) {
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	params := map[string]string{
		"hotelIds": strings.Join(hotelIDs, ","),
	}
	url := util.BuildURL(c.BaseURL, "/v2/e-reputation/hotel-sentiments", params)

	respBody, err := c.makeRequest("GET", url, token, nil)
	if err != nil {
		return nil, err
	}

	var result amadeusHotelModels.HotelSentimentResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse hotel sentiment response: %w", err)
	}

	return &result, nil
}

func (c *AmadeusClient) HotelNameAutoComplete(keyword string, subType string) (*amadeusHotelModels.HotelNameResponse, error) {
	token, err := c.TokenManager.GetValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Amadeus token: %w", err)
	}

	params := map[string]string{
		"keyword": keyword,
		"subType": subType,
	}
	url := util.BuildURL(c.BaseURL, "/v1/reference-data/locations/hotel", params)

	respBody, err := c.makeRequest("GET", url, token, nil)
	if err != nil {
		return nil, err
	}

	var result amadeusHotelModels.HotelNameResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse hotel name response: %w", err)
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("request failed, status: %s, body: %s", resp.Status, string(respBody))
	}

	return respBody, nil
}
