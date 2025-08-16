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
	FetchFiatRate(ctx context.Context, request types.FetchRateRequest) (float64, error)
	Convert(ctx context.Context, request types.ConvertRequest) (float64, error)

	History(ctx context.Context, request types.HistoryRequest) (float64, error)
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



func (s *ExchangeRateServiceImpl) getRateForCurrencies(base, target,date string) (float64, error) {
	// If the date is not provided, use the current date.
	if date == "" {
		date = time.Now().Format(internal.DateFormat)
	}

	baseIsFiat := internal.IsFiatCurrency(base)
	targetIsFiat := internal.IsFiatCurrency(target)

	// Case 1: Both currencies are fiat.
	if baseIsFiat && targetIsFiat {
		// If base is USD, it's a direct lookup in the fiat cache.
		if base == "USD" {
			key := "USD" + target
			rate, exists := s.fiatcache.GetRateWithDate(date, key)
			if exists {
				return rate, nil
			}
		} else {
			// Cross-rate calculation using USD as the base.
			rateUSDTarget, existsTarget := s.fiatcache.GetRateWithDate(date, "USD"+target)
			rateUSDBase, existsBase := s.fiatcache.GetRateWithDate(date, "USD"+base)
			if existsTarget && existsBase {
				if rateUSDBase == 0 {
					return 0, fmt.Errorf("invalid rate for %s, cannot divide by zero", base)
				}
				return rateUSDTarget / rateUSDBase, nil
			}
		}
	}

	// Case 2: Both currencies are crypto.
	if !baseIsFiat && !targetIsFiat {
		// Cross-rate calculation using USD as the base.
		rateUSDTarget, existsTarget := s.cryptocache.GetRateWithDate(date, "USD"+target)
		rateUSDBase, existsBase := s.cryptocache.GetRateWithDate(date, "USD"+base)
		if existsTarget && existsBase {
			if rateUSDBase == 0 {
				return 0, fmt.Errorf("invalid rate for %s, cannot divide by zero", base)
			}
			return rateUSDTarget / rateUSDBase, nil
		}
	}

	// Case 3: Mixed currencies (fiat to crypto).
	if baseIsFiat && !targetIsFiat {
		rateFiatUSD, existsFiat := s.fiatcache.GetRateWithDate(date, "USD"+base)
		rateCryptoUSD, existsCrypto := s.cryptocache.GetRateWithDate(date, "USD"+target)
		if existsFiat && existsCrypto {
			if rateCryptoUSD == 0 {
				return 0, fmt.Errorf("invalid rate for %s, cannot divide by zero", target)
			}
			// Rate is (1 / USD-to-Fiat) * USD-to-Crypto
			return (1 / rateFiatUSD) * rateCryptoUSD, nil
		}
	}

	// Case 4: Mixed currencies (crypto to fiat).
	if !baseIsFiat && targetIsFiat {
		rateCryptoUSD, existsCrypto := s.cryptocache.GetRateWithDate(date, "USD"+base)
		rateFiatUSD, existsFiat := s.fiatcache.GetRateWithDate(date, "USD"+target)
		if existsCrypto && existsFiat {
			if rateFiatUSD == 0 {
				return 0, fmt.Errorf("invalid rate for %s, cannot divide by zero", target)
			}
			// Rate is USD-to-Crypto / USD-to-Fiat
			return rateCryptoUSD / rateFiatUSD, nil
		}
	}

	return 0, fmt.Errorf("exchange rate not found for %s to %s on %s", base, target, date)
}
