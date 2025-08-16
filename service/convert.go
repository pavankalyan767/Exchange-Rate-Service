package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pavankalyan767/exchange-rate-service/internal"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

func (s *ExchangeRateServiceImpl) Convert(ctx context.Context, request types.ConvertRequest) (float64, error) {
	fmt.Println("inside convert service method")

	cache := s.cache
	if cache == nil {
		return 0, fmt.Errorf("cache is not initialized")
	}

	allowed_currencies := internal.AllowedCurrencies
	date := time.Now().Format(internal.DateFormat)
	_, check1 := allowed_currencies[request.BaseCurrency]
	if !check1 {
		return 0, fmt.Errorf("base currency %s is not allowed", request.BaseCurrency)
	}

	_, check2 := allowed_currencies[request.TargetCurrency]
	if !check2 {
		return 0, fmt.Errorf("target currency %s is not allowed", request.TargetCurrency)
	}

	if request.BaseCurrency == "USD" {
		key := request.BaseCurrency + request.TargetCurrency
		rate, exists := cache.GetLiveRate(date, key)
		if exists {
			return rate * request.Amount, nil
		}
	} else {
		key1 := "USD" + request.BaseCurrency
		key2 := "USD" + request.TargetCurrency
		rate1, exists1 := cache.GetLiveRate(date, key1)
		rate2, exists2 := cache.GetLiveRate(date, key2)
		if exists1 && exists2 {
			rate := rate2 / rate1
			return rate * request.Amount, nil
		}
	}
	return 0, fmt.Errorf("exchange rate not found for %s to %s", request.BaseCurrency, request.TargetCurrency)

}
