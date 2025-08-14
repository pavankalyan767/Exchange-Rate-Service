
package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/pavankalyan767/exchange-rate-service/service"
	"github.com/pavankalyan767/exchange-rate-service/types"
)



func FetchExchangeRateEndPoint(svc service.ExchangeRateServiceImpl) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(types.FetchExchangeRateRequest)
		ctx := context.Background()
		rate, err := svc.FetchExchangeRate(ctx, req)
		if err != nil {
			a := &types.FetchExchangeRateResponse{Rate: rate, Error: err.Error()}
			return a, nil
		}
		return types.FetchExchangeRateResponse{rate, ""}, nil
	}
}

func DecodeFetchExchangeRateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request types.FetchExchangeRateRequest

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

func EncodeFetchExchangeRateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
