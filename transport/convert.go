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

func ConvertEndPoint(svc service.ExchangeRateServiceImpl) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(types.ConvertRequest)
		ctx := context.Background()
		amount, err := svc.Convert(ctx, req)
		if err != nil {
			a := &types.ConvertResponse{Amount: amount, Error: err.Error()}
			return a, nil
		}
		return types.ConvertResponse{Amount: amount}, nil
	}
}

func DecodeConvertRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request types.ConvertRequest
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&request, r.URL.Query()); err != nil {
		return nil, err
	}

	// Read the entire body into a byte slice

	fmt.Println("the request after decoding it", request)
	return request, nil
}

func EncodeConvertResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
