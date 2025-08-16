package transport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/schema"
	"github.com/pavankalyan767/exchange-rate-service/service"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

func ConvertEndpoint(svc service.ExchangeRateService) endpoint.Endpoint {
	fmt.Println("inside convert endpoint")
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(types.ConvertRequest)
		ctx := context.Background()
		fmt.Print("the request in convert endpoint: ", req)
		amount, err := svc.Convert(ctx, &req)
		if err != nil {
			a := &types.ConvertResponse{ConvertedAmount: amount, Error: err.Error()}
			return a, nil
		}
		return types.ConvertResponse{ConvertedAmount: amount}, nil
	}
}

func DecodeConvertRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("inside decode convert request")
	var request types.ConvertRequest
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&request, r.URL.Query()); err != nil {
		return nil, fmt.Errorf("error decoding convert request: %w", err)
	}

	// Read the entire body into a byte slice

	fmt.Println("the request after decoding it", request)
	return request, nil
}
