package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/schema"
	"github.com/pavankalyan767/exchange-rate-service/service"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

func FetchFiatRateEndpoint(svc *service.ExchangeRateServiceImpl) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(types.FetchRateRequest)
		ctx := context.Background()
		rate, err := svc.FetchRate(ctx, &req)
		if err != nil {
			a := &types.FetchRateResponse{Rate: rate, Error: err.Error()}
			return a, nil
		}
		return types.FetchRateResponse{rate, ""}, nil
	}
}

func DecodeFetchFiatRateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request types.FetchRateRequest

	// Read the entire body into a byte slice
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&request, r.URL.Query()); err != nil {
		return nil, err
	}

	// Now decode from the byte slice

	fmt.Println("the request after decoding it", request)
	return request, nil
}

func EncodeFetchFiatRateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
