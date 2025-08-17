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
	FetchRate(ctx context.Context, request *types.FetchRateRequest) (float64, error)
	Convert(ctx context.Context, request *types.ConvertRequest) (float64, error)

	History(ctx context.Context, request *types.HistoryRequest) (map[string]float64, error)
}

type ExchangeRateServiceImpl struct {
	fiatcache   *cache.Cache
	cryptocache *cache.Cache
}

func NewExchangeRateServiceImpl(fiatcache, cryptocache *cache.Cache) *ExchangeRateServiceImpl {
	return &ExchangeRateServiceImpl{
		fiatcache:   fiatcache,
		cryptocache: cryptocache,
	}

}
func (s *ExchangeRateServiceImpl) getRateForCurrencies(base, target, date string) (float64, error) {
	// If the date is not provided, use the current date.
	if date == "" {
		date = time.Now().Format(internal.DateFormat)
	}

	baseIsFiat := internal.IsFiatCurrency(base)
	targetIsFiat := internal.IsFiatCurrency(target)

	// Case 1: Both currencies are fiat.
	// The rate is (USD->Target) / (USD->Base).
	if baseIsFiat && targetIsFiat {
		// Direct lookup for USD to any fiat.
		if base == internal.BaseCurrency {
			key := internal.BaseCurrency + target
			rate, exists := s.fiatcache.GetRateWithDate(date, key)
			if exists {
				return 1/rate, nil
			}
		}else if target==internal.BaseCurrency{
			key := internal.BaseCurrency + base
			rate,exists := s.fiatcache.GetRateWithDate(date,key)
			if exists{
				return 1/rate,nil
			}
		}else {
			// Cross-rate calculation for any fiat to any fiat (e.g., EUR to INR).
			rateUSDTarget, existsTarget := s.fiatcache.GetRateWithDate(date, internal.BaseCurrency+target)
			rateUSDBase, existsBase := s.fiatcache.GetRateWithDate(date, internal.BaseCurrency+base)
			if existsTarget && existsBase {
				if rateUSDBase == 0 {
					return 0, fmt.Errorf("invalid rate for %s, cannot divide by zero", base)
				}
				return rateUSDTarget / rateUSDBase, nil
			}
		}
	}

	// Case 2: Both currencies are crypto.
	// The rate is (Base->USD) / (Target->USD).
	if !baseIsFiat && !targetIsFiat {
		rateBaseUSD, existsBase := s.cryptocache.GetRateWithDate(date, base+internal.BaseCurrency)
		rateTargetUSD, existsTarget := s.cryptocache.GetRateWithDate(date, target+internal.BaseCurrency)
		if existsBase && existsTarget {
			if rateTargetUSD == 0 {
				return 0, fmt.Errorf("invalid rate for %s, cannot divide by zero", target)
			}
			return rateBaseUSD / rateTargetUSD, nil
		}
	}

	// Case 3: Mixed currencies (fiat to crypto).
	
	if baseIsFiat && !targetIsFiat {
		var rateUSDTarget float64
		var existsFiat bool
		if base != internal.BaseCurrency {
			rateUSDTarget, existsFiat = s.fiatcache.GetRateWithDate(date, internal.BaseCurrency+base)
		}else{
			rateUSDTarget=1
			existsFiat=true
		}
		
		
		rateTargetUSD, existsCrypto := s.cryptocache.GetRateWithDate(date, target+internal.BaseCurrency)
		if existsFiat && existsCrypto {
			if rateUSDTarget == 0 {
				return 0, fmt.Errorf("invalid rate for %s, cannot divide by zero", base)
			}
			return 1/(rateUSDTarget * rateTargetUSD), nil
		}
	}

	// Case 4: Mixed currencies (crypto to fiat).
	
	if !baseIsFiat && targetIsFiat {
		rateBaseUSD, existsCrypto := s.cryptocache.GetRateWithDate(date, base+internal.BaseCurrency)
		
		if target != internal.BaseCurrency {
		rateUSDTarget, existsFiat := s.fiatcache.GetRateWithDate(date, internal.BaseCurrency+target)
		if existsCrypto && existsFiat {
			return rateBaseUSD * rateUSDTarget, nil
		}}else{
			return rateBaseUSD,nil
		}
		
	}

	return 0, fmt.Errorf("exchange rate not found for %s to %s on %s", base, target, date)
}
