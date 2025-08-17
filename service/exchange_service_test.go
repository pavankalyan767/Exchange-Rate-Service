package service_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-kit/log"
	"github.com/pavankalyan767/exchange-rate-service/cache"
	"github.com/pavankalyan767/exchange-rate-service/internal"
	"github.com/pavankalyan767/exchange-rate-service/service"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

func setupServiceWithMockRates() *service.ExchangeRateServiceImpl {
	logger := log.NewLogfmtLogger(os.Stderr)
	fiatCache := cache.NewCache(1*time.Minute, 10*time.Second, logger)
	cryptoCache := cache.NewCache(1*time.Minute, 10*time.Second, logger)

	today := time.Now().Format(internal.DateFormat)
	yesterday := time.Now().AddDate(0, 0, -1).Format(internal.DateFormat)

	// Fiat rates
	fiatRates := map[string]float64{
		"USDINR": 83.0,
		"USDEUR": 0.91,
		"USDJPY": 148.2,
	}

	// Cache for today and yesterday
	fiatCache.Set(today, fiatRates, 1*time.Minute)
	fiatCache.Set(yesterday, fiatRates, 1*time.Minute)

	// Crypto rates
	cryptoRates := map[string]float64{
		"BTCUSD": 30000.0,
		"ETHUSD": 1800.0,
	}

	cryptoCache.Set(today, cryptoRates, 24*time.Hour)

	return service.NewExchangeRateServiceImpl(fiatCache, cryptoCache)
}



func TestConvert_ValidFiatToFiat(t *testing.T) {
	svc := setupServiceWithMockRates()

	req := &types.ConvertRequest{
		BaseCurrency:   "USD",
		TargetCurrency: "INR",
		Amount:         100,
		Date:           time.Now().Format(internal.DateFormat),
	}

	result, err := svc.Convert(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := 100 * 83.0
	if result != expected {
		t.Errorf("expected %.2f, got %.2f", expected, result)
	}
}



func TestConvert_InvalidCurrency(t *testing.T) {
	svc := setupServiceWithMockRates()

	req := &types.ConvertRequest{
		BaseCurrency:   "XXX",
		TargetCurrency: "INR",
		Amount:         50,
		Date:           time.Now().Format(internal.DateFormat),
	}

	_, err := svc.Convert(context.Background(), req)
	if err == nil {
		t.Fatalf("expected error for invalid currency")
	}
}



func TestFetchRate_ValidCryptoToCrypto(t *testing.T) {
	svc := setupServiceWithMockRates()

	req := &types.FetchRateRequest{
		BaseCurrency:   "BTC",
		TargetCurrency: "ETH",
		Date:           time.Now().Format(internal.DateFormat),
	}

	rate, err := svc.FetchRate(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := 30000.0 / 1800.0
	if rate != expected {
		t.Errorf("expected %.2f, got %.2f", expected, rate)
	}
}



func TestHistory_ValidRange(t *testing.T) {
	svc := setupServiceWithMockRates()

	today := time.Now().Format(internal.DateFormat)
	yesterday := time.Now().AddDate(0, 0, -1).Format(internal.DateFormat)

	req := &types.HistoryRequest{
		BaseCurrency:   "USD",
		TargetCurrency: "INR",
		From:           yesterday,
		To:             today,
	}

	rates, err := svc.History(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(rates) != 2 {
		t.Errorf("expected 2 rates, got %d", len(rates))
	}
}



func TestHistory_InvalidRange(t *testing.T) {
	svc := setupServiceWithMockRates()

	tomorrow := time.Now().AddDate(0, 0, 1).Format(internal.DateFormat)

	req := &types.HistoryRequest{
		BaseCurrency:   "USD",
		TargetCurrency: "INR",
		From:           tomorrow,
		To:             tomorrow,
	}

	_, err := svc.History(context.Background(), req)
	if err == nil {
		t.Fatalf("expected error for missing rate data in future")
	}
}



