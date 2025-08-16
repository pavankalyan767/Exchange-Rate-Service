package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// APIClient encapsulates all logic for making HTTP requests to the external API.
// It holds its own dependencies (baseURL, apiKey, and http client).
type APIClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

// NewAPIClient is a constructor that creates and returns a new APIClient instance.
// This is where all dependencies are injected.
func NewAPIClient(baseURL, apiKey string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		apiKey:  apiKey,
		// The http.Client is created once here and reused for all requests.
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

// Get performs a GET request and returns the response body as a byte slice.
// It uses the pre-initialized http.Client from the struct.
func (c *APIClient) Get(ctx context.Context, requestURL string) ([]byte, error) {
	fmt.Println("Request URL:", requestURL)

	
	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch exchange rate from API. Status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

// BuildLiveURL constructs the URL for a live exchange rate request.
// It is now a method of APIClient and uses the encapsulated baseURL and apiKey.
func (c *APIClient) BuildLiveURL(endpoint, baseCurrency string) string {
	return fmt.Sprintf("%s%s?access_key=%s&source=%s", c.baseURL, endpoint, c.apiKey, baseCurrency)
}

// BuildHistoryURL constructs the URL for a historical data request.
// It is now a method of APIClient and uses the encapsulated baseURL and apiKey.
func (c *APIClient) BuildHistoryURL(startDate, endDate string, currencies map[string]struct{}) string {
	var currency_slice []string
	for currency := range currencies {
		currency_slice = append(currency_slice, currency)
	}
	currenciesStr := strings.Join(currency_slice, ",")

	return fmt.Sprintf("%s/timeframe?access_key=%s&currencies=%s&start_date=%s&end_date=%s",
		c.baseURL, c.apiKey, currenciesStr, startDate, endDate)
}
