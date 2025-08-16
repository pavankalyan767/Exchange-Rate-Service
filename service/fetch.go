package service

import (
	"context"
	"fmt"

	"github.com/pavankalyan767/exchange-rate-service/internal"
	"github.com/pavankalyan767/exchange-rate-service/types"
)



func (s *ExchangeRateServiceImpl) FetchRate(ctx context.Context, req *types.FetchRateRequest) (output float64, err error) {
	// Validate the input currencies
	if !internal.IsAllowedCurrency(req.BaseCurrency) || !internal.IsAllowedCurrency(req.TargetCurrency) {
		return 0, fmt.Errorf("invalid currency: %s or %s", req.BaseCurrency, req.TargetCurrency)
	}

	// Use a single helper function to get the rate for any currency pair.
	rate, err := s.getRateForCurrencies(req.BaseCurrency, req.TargetCurrency,req.Date)
	if err != nil {
		return 0, fmt.Errorf("could not fetch rate: %v", err)
	}

	return rate,nil
}
