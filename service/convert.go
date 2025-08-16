package service

import (
	"context"
	"fmt"

	"github.com/pavankalyan767/exchange-rate-service/internal"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

func (s *ExchangeRateServiceImpl) Convert(ctx context.Context, request types.ConvertRequest) (float64, error) {
	// Validate inputs
	allowedCurrencies := internal.AllowedCurrencies
	_, check1 := allowedCurrencies[request.BaseCurrency]
	_, check2 := allowedCurrencies[request.TargetCurrency]
	if !check1 || !check2 {
		return 0, fmt.Errorf("invalid currency provided")
	}

	// Use the helper function to get the base exchange rate.
	baseRate, err := s.getRateForCurrencies(request.BaseCurrency, request.TargetCurrency,"")
	if err != nil {
		return 0, err
	}

	// Apply the conversion amount and return.
	return baseRate * request.Amount, nil
}
