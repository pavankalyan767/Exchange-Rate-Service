package types

// FetchExchangeRate types
type FetchRateRequest struct {
	BaseCurrency   string `json:"base_currency" schema:"base_currency"`
	TargetCurrency string `json:"target_currency"schema:"target_currency"`
	Date           string `json:"date"`
}

type FetchRateResponse struct {
	Rate  float64 `json:"rate"`
	Error string  `json:"err,omitempty"`
}

// Convert types
type ConvertRequest struct {
	BaseCurrency   string  `json:"base_currency" schema:"base_currency"`
	TargetCurrency string  `json:"target_currency" schema:"target_currency"`
	Date           string  `json:"date" schema:"date"`
	Amount         float64 `json:"amount" schema:"amount"`
}

// ConvertResponse defines the structure for a currency conversion response.
type ConvertResponse struct {
	ConvertedAmount float64 `json:"convertedAmount"`
	Error           string  `json:"error,omitempty"`
}

// History types
type HistoryRequest struct {
	BaseCurrency   string `json:"base_currency" schema:"base_currency"`
	TargetCurrency string `json:"target_currency" schema:"target_currency"`
	From           string `json:"from" schema:"from"`
	To             string `json:"to"schema:"to"`
}

type HistoryResponse struct {
	Rates map[string]float64 `json:"rates"`
	Error string             `json:"err,omitempty"`
}
