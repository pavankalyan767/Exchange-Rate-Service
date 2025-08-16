package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/joho/godotenv"
	"github.com/pavankalyan767/exchange-rate-service/cache"
	"github.com/pavankalyan767/exchange-rate-service/client"
	service "github.com/pavankalyan767/exchange-rate-service/service"
	"github.com/pavankalyan767/exchange-rate-service/transport"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rateCache := cache.NewCache(5*time.Minute, 10*time.Minute)

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalf("Error: EXCHANGE_RATE_API_KEY environment variable is not set.")
	}
	baseUrl := os.Getenv("BASE_API_URL")
	if baseUrl == "" {
		log.Fatalf("Error: BASE_API_URL environment variable is not set.")
	}
	apiClient := client.NewAPIClient(baseUrl, apiKey)

	svc := service.NewExchangeRateServiceImpl(rateCache)
	rate_fetcher := service.NewRateFetcher(apiClient, rateCache)

	// --- Core Polling Logic ---

	// 1. Perform an initial fetch of historical rates with a dedicated context.
	log.Println("Performing initial fetch of historical rates...")
	initialCtx, initialCancel := context.WithTimeout(context.Background(), 30*time.Second)
	if err := rate_fetcher.HistoricalRate(initialCtx); err != nil {
		log.Fatalf("Failed to fetch historical rates on startup: %v", err)
	}
	initialCancel() // Call cancel for the initial fetch's context

	// 2. Start a background goroutine for hourly polling.
	//    The ticker will tick every hour, triggering a new API call.
	log.Println("Starting hourly polling for live rates...")
	ctx,cancel := context.WithCancel(context.Background())
	if err:= rate_fetcher.LiveRate(ctx); err != nil {
		log.Fatalf("Failed to fetch live rates on startup: %v", err)
	}
	defer cancel() // Ensure the context is cancelled when the main function exits
	ticker := time.NewTicker( 1* time.Hour)
	defer ticker.Stop()

	go func() {
		// This loop waits for the ticker to "tick"
		for range ticker.C {
			log.Println("Hourly poll initiated...")
			// Create a new, short-lived context for each recurring API call.
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			if err := rate_fetcher.LiveRate(ctx); err != nil {
				log.Printf("Error during hourly rate polling: %v", err)
			}
			cancel() // Clean up context after the call is done
		}
	}()

	fetchExchangeRateHandler := httptransport.NewServer(
		transport.FetchExchangeRateEndPoint(svc),
		transport.DecodeFetchExchangeRateRequest,
		transport.EncodeFetchExchangeRateResponse,
	)
	convertHandler := httptransport.NewServer(
		transport.ConvertEndPoint(svc),
		transport.DecodeConvertRequest,
		transport.EncodeConvertResponse,
	)
	historyHandler := httptransport.NewServer(
		transport.HistoryEndpoint(svc),
		transport.DecodeHistoryRequest,
		transport.EncodeHistoryResponse,
	)

	http.Handle("/fetch", fetchExchangeRateHandler)
	http.Handle("/convert", convertHandler)
	http.Handle("/history", historyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
