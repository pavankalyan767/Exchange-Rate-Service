package transport

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/schema"
	"github.com/pavankalyan767/exchange-rate-service/service"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

func HistoryEndpoint(svc service.ExchangeRateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(types.HistoryRequest)
		ctx := context.Background()
		rates, err := svc.History(ctx, &req)
		if err != nil {
			a := &types.HistoryResponse{Rates: rates, Error: err.Error()}
			return a, nil
		}
		return types.HistoryResponse{Rates: rates}, nil
	}
}

func DecodeHistoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request types.HistoryRequest

	// Read the entire body into a byte slice
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&request, r.URL.Query()); err != nil {
		return nil, err
	}

	// Now decode from the byte slice

	return request, nil
}
