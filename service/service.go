package service

import (
	"context"
	"errors"
	"time"

	"github.com/pavankalyan767/exchange-rate-service/types"
)

type ExchangeRateService interface {
	FetchExchangeRate(ctx context.Context, request types.FetchExchangeRateRequest) (float64, error)
	ConvertCurrency(ctx context.Context, amount float64, baseCurrency string, targetCurrency string, date *time.Time) (float64, error)
	UpdateExchangeRate(ctx context.Context, baseCurrency string, targetCurrency string, rate float64) error
}

type ExchangeRateServiceImpl struct{}

func (s *ExchangeRateServiceImpl) ConvertCurrency(ctx context.Context, amount float64, baseCurrency string, targetCurrency string, date *time.Time) (float64, error) {
	return 0, errors.New("not implemented")
}

func (s *ExchangeRateServiceImpl) UpdateExchangeRate(ctx context.Context, baseCurrency string, targetCurrency string, rate float64) error {
	return errors.New("not implemented")
}
