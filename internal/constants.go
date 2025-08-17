package internal

var AllowedFiatCurrencies = map[string]struct{}{
	"USD": {},
	"INR": {},
	"EUR": {},
	"JPY": {},
	"GBP": {},
}

var AllowedCryptoCurrencies = map[string]struct{}{
	"BTC": {},
	"ETH": {},
	"USDT": {},
}

const (
	// LookbackDays is the maximum number of days for historical data.
	LookbackDays = 90

	// DateFormat is the required date format for historical requests.
	DateFormat = "2006-01-02"
	BaseCurrency = "USD"
)

func IsFiatCurrency(currency string) bool {
	_, exists := AllowedFiatCurrencies[currency]
	return exists
}

// IsCryptoCurrency checks if the given currency is a crypto currency.
func IsCryptoCurrency(currency string) bool {
	_, exists := AllowedCryptoCurrencies[currency]
	return exists
}

// IsAllowedCurrency checks if the given currency is either a fiat or a crypto currency.
func IsAllowedCurrency(currency string) bool {
	return IsFiatCurrency(currency) || IsCryptoCurrency(currency)
}
