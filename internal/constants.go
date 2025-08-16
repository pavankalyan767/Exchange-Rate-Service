package internal
var AllowedCurrencies = map[string]struct{}{
	"USD": {},
	"INR": {},
	"EUR": {},
	"JPY": {},
	"GBP": {},
}

const (
	// LookbackDays is the maximum number of days for historical data.
	LookbackDays = 90

	// DateFormat is the required date format for historical requests.
	DateFormat = "2006-01-02"
)

