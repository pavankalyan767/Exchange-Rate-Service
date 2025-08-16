package transport

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/pavankalyan767/exchange-rate-service/service"
)

type Endpoints struct {
	HistoryEndpoint endpoint.Endpoint
	FetchEndpoint   endpoint.Endpoint
	ConvertEndpoint endpoint.Endpoint
}

func MakeEndpoints(s service.ExchangeRateService) Endpoints {
	return Endpoints{
		HistoryEndpoint: HistoryEndpoint(s),
		FetchEndpoint:   FetchEndpoint(s),
		ConvertEndpoint: ConvertEndpoint(s),
	}
}
