package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/pavankalyan767/exchange-rate-service/transport"
)

type ExchangeRateService interface {
	FetchExchangeRate(ctx context.Context, request FetchExchangeRateRequest) (float64, error)
	ConvertCurrency(ctx context.Context, amount float64, baseCurrency string, targetCurrency string, date *time.Time) (float64, error)
	UpdateExchangeRate(ctx context.Context, baseCurrency string, targetCurrency string, rate float64) error
}

type ExchangeRateServiceImpl struct{}

type ExchangeRateAPIResponse struct {
	Result         string  `json:"result"`
	ConversionRate float64 `json:"conversion_rate"`
	BaseCode       string  `json:"base_code"`
	TargetCode     string  `json:"target_code"`
	Timestamp      int64   `json:"time_last_update_unix"`
}

func (s *ExchangeRateServiceImpl) FetchExchangeRate(ctx context.Context, request FetchExchangeRateRequest) (float64, error) {
	baseCurrency := request.BaseCurrency
	targetCurrency := request.TargetCurrency
	if baseCurrency == "" || targetCurrency == "" {
		return 0, errors.New("base currency and target currency must be provided")
	}

	// Create a new HTTP client and set the base URL and bearer token from environment variables.
	client := &http.Client{Timeout: 10 * time.Second}
	base_url := os.Getenv("BASE_API_URL")
	fmt.Println("Base API URL:", base_url)
	if base_url == "" {
		return 0, errors.New("base API URL is not set")
	}

	api_key := os.Getenv("API_KEY")
	if api_key == "" {
		return 0, errors.New("api key is not set")
	}

	request_url := fmt.Sprintf("%s/%s/pair/%s/%s", base_url, api_key, baseCurrency, targetCurrency)
	req, err := http.NewRequestWithContext(ctx, "GET", request_url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("failed to fetch exchange rate from API")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Log the response body for debugging purposes.
	var apiResponse ExchangeRateAPIResponse
	err = json.Unmarshal(body, &apiResponse)

	if err != nil {
		return 0, fmt.Errorf("error parsing API response: %v", err)
	}
	log.Println("API Response:", apiResponse)

	rate := apiResponse.ConversionRate
	return rate, nil
}

func (s *ExchangeRateServiceImpl) ConvertCurrency(ctx context.Context, amount float64, baseCurrency string, targetCurrency string, date *time.Time) (float64, error) {
	return 0, errors.New("not implemented")
}

func (s *ExchangeRateServiceImpl) UpdateExchangeRate(ctx context.Context, baseCurrency string, targetCurrency string, rate float64) error {
	return errors.New("not implemented")
}
