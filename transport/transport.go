package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/pavankalyan767/exchange-rate-service/service"
)

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
