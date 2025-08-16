package service

import (
	"context"

	"github.com/pavankalyan767/exchange-rate-service/cache"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

type ExchangeRateService interface {
	FetchExchangeRate(ctx context.Context, request types.FetchRateRequest) (float64, error)
	Convert(ctx context.Context, request types.ConvertRequest) (float64, error)

	History(ctx context.Context, request types.HistoryRequest) (float64, error)
	
}

type ExchangeRateServiceImpl struct{
	cache *cache.Cache
}

func NewExchangeRateServiceImpl(c *cache.Cache) *ExchangeRateServiceImpl {
	return &ExchangeRateServiceImpl{
		cache: c,
	}
	
}



