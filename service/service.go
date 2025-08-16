package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pavankalyan767/exchange-rate-service/cache"
	"github.com/pavankalyan767/exchange-rate-service/internal"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

type ExchangeRateService interface {
	FetchExchangeRate(ctx context.Context, request types.FetchRateRequest) (float64, error)
	Convert(ctx context.Context, request types.ConvertRequest) (float64, error)

	History(ctx context.Context, request types.HistoryRequest) (float64, error)
}

type ExchangeRateServiceImpl struct {
	cache *cache.Cache
}

func NewExchangeRateServiceImpl(c *cache.Cache) *ExchangeRateServiceImpl {
	return &ExchangeRateServiceImpl{
		cache: c,
	}

}

func (s *ExchangeRateServiceImpl) getRateForCurrencies(base, target string, date string) (float64, error) {
    cache := s.cache
    if cache == nil {
        return 0, fmt.Errorf("cache is not initialized")
    }

    // If the date is not provided, use the current date.
    if date == "" {
        date = time.Now().Format(internal.DateFormat)
    }

    // Check if the rate is directly available (if base currency is USD).
    if base == "USD" {
        key := base + target
        rate, exists := cache.GetRateWithDate(date, key)
        if exists {
            return rate, nil
        }
    } else {
        // If the base currency is not USD, convert it using the USD rates.
        key1 := "USD" + base
        key2 := "USD" + target
        rate1, exists1 := cache.GetRateWithDate(date, key1)
        rate2, exists2 := cache.GetRateWithDate(date, key2)
        if exists1 && exists2 {
            rate := rate2 / rate1
            return rate, nil
        }
    }

    return 0, fmt.Errorf("exchange rate not found for %s to %s", base, target)
}



