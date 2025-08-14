package types

import "time"

// FetchExchangeRate types
type FetchExchangeRateRequest struct {
	BaseCurrency   string `json:"baseCurrency"`
	TargetCurrency string `json:"targetCurrency"`
}

type FetchExchangeRateResponse struct {
	Rate  float64 `json:"rate"`
	Error string  `json:"err,omitempty"`
}

// Convert types
type ConvertRequest struct {
	From   string    `json:"from"`
	To     string    `json:"to"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}

type ConvertResponse struct {
	Amount float64 `json:"amount"`
	Error  string  `json:"err,omitempty"`
}
