package service

import (
	"context"
	"fmt"

	"github.com/pavankalyan767/exchange-rate-service/internal"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

func (s *ExchangeRateServiceImpl) Convert(ctx context.Context, req *types.ConvertRequest) (float64, error) {
	// Validate the input currencies and amount
	if !internal.IsAllowedCurrency(req.BaseCurrency) || !internal.IsAllowedCurrency(req.TargetCurrency) {
		return 0, fmt.Errorf("invalid currency: %s or %s", req.BaseCurrency, req.TargetCurrency)
	}
	if req.Amount <= 0 {
		return 0, fmt.Errorf("invalid amount: %f", req.Amount)
	}

	// Fetch the rate using the unified helper function
	rate, err := s.getRateForCurrencies(req.BaseCurrency, req.TargetCurrency,req.Date)
	if err != nil {
		return 0, fmt.Errorf("could not fetch rate: %v", err)
	}

	convertedAmount := req.Amount * rate

	return convertedAmount, nil
}