package main

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

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/joho/godotenv"
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

type FetchExchangeRateRequest struct {
	BaseCurrency   string `json:"baseCurrency"`
	TargetCurrency string `json:"targetCurrency"`
}

type FetchExchangeRateResponse struct {
	Rate  float64 `json:"rate"`
	Error string  `json:"err,omitempty"`
}

func FetchExchangeRateEndPoint(svc ExchangeRateServiceImpl) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(FetchExchangeRateRequest)
		ctx := context.Background()
		rate, err := svc.FetchExchangeRate(ctx, req)
		if err != nil {
			return FetchExchangeRateResponse{rate, err.Error()}, nil
		}
		return FetchExchangeRateResponse{rate, ""}, nil
	}
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
	req, err := http.NewRequestWithContext(ctx,"GET", request_url, nil)
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

func main() {

	svc := ExchangeRateServiceImpl{}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fetchExchangeRateHandler := httptransport.NewServer(
		FetchExchangeRateEndPoint(svc),
		decodeFetchExchangeRateRequest,
		encodeResponse,
	)

	http.Handle("/fetch_exchange_rate", fetchExchangeRateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func decodeFetchExchangeRateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request FetchExchangeRateRequest

	// Read the entire body into a byte slice
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read request body: %w", err)
	}

	// Now decode from the byte slice
	if err := json.Unmarshal(body, &request); err != nil {
		return nil, err
	}

	fmt.Println("the request after decoding it", request)
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
