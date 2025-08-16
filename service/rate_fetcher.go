package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"time"

	"github.com/pavankalyan767/exchange-rate-service/cache"
	"github.com/pavankalyan767/exchange-rate-service/client"
	"github.com/pavankalyan767/exchange-rate-service/internal"
)

type RateFetcher struct {
	apiClient *client.APIClient
	cache     *cache.Cache
}

func NewRateFetcher(apiClient *client.APIClient, appCache *cache.Cache) *RateFetcher {
	return &RateFetcher{
		apiClient: apiClient,
		cache:     appCache,
	}
}

type LiveRateAPIResponse struct {
	Quotes map[string]float64 `json:"quotes"`
	Result float64            `json:"result,omitempty"`
}

type HistoricalRateAPIResponse struct {
	Quotes map[string]map[string]float64 `json:"quotes"`
}

func (rf *RateFetcher) LiveRate(ctx context.Context) error {

	endpoint := "live"
	baseCurrency := "USD"

	request_url := rf.apiClient.BuildLiveURL(endpoint, baseCurrency)

	resp, err := rf.apiClient.Get(ctx, request_url)
	if err != nil {
		return fmt.Errorf("error getting response from api client for live rates: %v", err)
	}

	var liveRateResponse LiveRateAPIResponse
	err = json.Unmarshal(resp, &liveRateResponse)
	if err != nil {
		return fmt.Errorf("error unmarshalling live rate response: %v", err)
	}
	if len(liveRateResponse.Quotes) == 0 {
		return errors.New("no quotes found in live rate response")
	}

	exchangeRate := make(map[string]float64)

	currencies := internal.AllowedCurrencies

	for currency := range currencies {
		if currency != baseCurrency {
			exchangeRate[baseCurrency+currency] = 0
		}
	}
	for key, value := range liveRateResponse.Quotes {
		exchangeRate[key] = value
	}

	today := time.Now().Format(internal.DateFormat)

    // Cache the entire map of today's rates using the date as the key.
    rf.cache.Set(today, exchangeRate, 24*time.Hour)
	fmt.Println("Live rates cached successfully")

	return nil

}

func (rf *RateFetcher) HistoricalRate(ctx context.Context) error {

	startDate := time.Now().AddDate(0, 0, -1*internal.LookbackDays).Format(internal.DateFormat)
	endDate := time.Now().AddDate(0,0,-1).Format(internal.DateFormat)
	request_url := rf.apiClient.BuildHistoryURL(startDate, endDate, internal.AllowedCurrencies)

	resp, err := rf.apiClient.Get(ctx, request_url)
	if err != nil {
		return fmt.Errorf("error getting response from api client for historical rates %w", err)
	}

	var historyRateResponse HistoricalRateAPIResponse
	if err := json.Unmarshal(resp, &historyRateResponse); err != nil {
		return fmt.Errorf("error unmarshalling historical rate response: %v", err)
	}

	Quotes := historyRateResponse.Quotes

	for date, rates := range Quotes {
		rf.cache.Set(date, rates, 24*90*time.Hour)
	}
	fmt.Println("Historical rates cached successfully")

	return nil

}
