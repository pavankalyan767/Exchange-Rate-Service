package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pavankalyan767/exchange-rate-service/internal"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

func (s *ExchangeRateServiceImpl) History(ctx context.Context, request types.HistoryRequest) (map[string]float64, error) {
	cache := s.fiatcache
	if cache == nil {
		return nil, fmt.Errorf("cache is not initialized")
	}

	// Validate the input currencies.
	allowed_currencies := internal.AllowedFiatCurrencies
	if _, check1 := allowed_currencies[request.BaseCurrency]; !check1 {
		return nil, fmt.Errorf("base currency %s is not allowed", request.BaseCurrency)
	}
	if _, check2 := allowed_currencies[request.TargetCurrency]; !check2 {
		return nil, fmt.Errorf("target currency %s is not allowed", request.TargetCurrency)
	}

	// Ensure 'from' and 'to' dates are provided.
	if request.From == "" || request.To == "" {
		return nil, fmt.Errorf("from and to dates must be provided")
	}

	// Parse the start and end dates from the request.
	from, err := time.Parse(internal.DateFormat, request.From)
	if err != nil {
		return nil, fmt.Errorf("invalid 'from' date format: %w", err)
	}
	to, err := time.Parse(internal.DateFormat, request.To)
	if err != nil {
		return nil, fmt.Errorf("invalid 'to' date format: %w", err)
	}

	// Ensure the 'from' date is not after the 'to' date.
	if from.After(to) {
		return nil, fmt.Errorf("'from' date cannot be after 'to' date")
	}

	// Initialize the map to store the historical rates.
	rates := make(map[string]float64)

	// Loop through each day from the 'from' date to the 'to' date.
	for d := from; !d.After(to); d = d.AddDate(0, 0, 1) {
		// Format the current date as a string for the getFiatRateForCurrencies function.
		dateString := d.Format(internal.DateFormat)

		// Get the exchange rate for the current day.
		rate, err := s.getRateForCurrencies(request.BaseCurrency, request.TargetCurrency, dateString)
		if err != nil {

			return nil, fmt.Errorf("failed to get rate for %s: %w", dateString, err)
		}
		rates[dateString] = rate
	}

	return rates, nil
}
