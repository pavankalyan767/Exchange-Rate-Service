package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type URLBuilder struct{}

func (b *URLBuilder) BuildLiveURL(baseURL, apiKey, endpoint, baseCurrency string) string {
	return fmt.Sprintf("%s%s?access_key=%s&source=%s", baseURL, endpoint, apiKey, baseCurrency)
}

func (b *URLBuilder) BuildHistoryURL(baseURL, apiKey, startDate, endDate string, currencies map[string]struct{}) string {
	var currency_slice []string
	for currency := range currencies{
		currency_slice = append(currency_slice,currency)
	}
	currenciesStr := strings.Join(currency_slice, ",")

	return fmt.Sprintf("%s/timeframe?access_key=%s&currencies=%s&start_date=%s&end_date=%s",
		baseURL, apiKey, currenciesStr, startDate, endDate)
}

func APIClient(request_url string) ([]byte, error) {

	client := &http.Client{Timeout: 10 * time.Second}
	ctx := context.Background()
	fmt.Println("Request URL:", request_url)
	req, err := http.NewRequestWithContext(ctx, "GET", request_url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("response body", resp)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch exchange rate from API")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
