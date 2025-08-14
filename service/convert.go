package service

import (
	"context"
	"fmt"

	"github.com/pavankalyan767/exchange-rate-service/types"
)

func (s *ExchangeRateServiceImpl) Convert(ctx context.Context, request types.ConvertRequest) (float64, error) {

	fetch_request := &types.FetchExchangeRateRequest{
		BaseCurrency:   request.From,
		TargetCurrency: request.To,
	}

	rate_response, err := s.FetchExchangeRate(ctx, *fetch_request)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch exchange rate: %w", err)
	}

	final_amount := request.Amount * rate_response
	return final_amount, nil

}
