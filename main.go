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

	fiatCache := cache.NewCache(5*time.Minute, 10*time.Minute)
	cryptoCache := cache.NewCache(5*time.Minute, 10*time.Minute)

	fiatapikey := os.Getenv("FIAT_API_KEY")
	if fiatapikey == "" {
		log.Fatalf("Error: EXCHANGE_RATE_API_KEY environment variable is not set.")
	}
	cryptoapikey := os.Getenv("CRYPTO_API_KEY")
	if cryptoapikey == "" {
		log.Printf("Warning: CRYPTO_API_KEY environment variable is not set. Crypto-related requests will not be available.")
	}
	fiatUrl := os.Getenv("FIAT_API_URL")
	if fiatUrl == "" {
		log.Fatalf("Error: BASE_API_URL environment variable is not set.")
	}
	cryptoUrl := os.Getenv("CRYPTO_API_URL")
	if cryptoUrl == "" {
		log.Printf("Warning: CRYPTO_API_URL environment variable is not set. Crypto-related requests will not be available.")
	}
	apiClient := client.NewAPIClient(fiatUrl, cryptoUrl, fiatapikey, cryptoapikey)

	svc := service.NewExchangeRateServiceImpl(fiatCache, cryptoCache)
	rate_fetcher := service.NewRateFetcher(apiClient, fiatCache, cryptoCache)

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
	log.Println("Starting hourly polling for live fiat rates...")
	ctx, cancel := context.WithCancel(context.Background())
	if err := rate_fetcher.LiveRate(ctx); err != nil {
		log.Fatalf("Failed to fetch live rates on startup: %v", err)
	}
	defer cancel() // Ensure the context is cancelled when the main function exits
	ticker := time.NewTicker(1 * time.Hour)
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

	log.Println("Starting hourly polling for live crypto rates...")
	ctx1, cancel := context.WithCancel(context.Background())
	if err := rate_fetcher.CryptoRate(ctx1); err != nil {
		log.Fatalf("Failed to fetch live rates on startup: %v", err)
	}
	defer cancel() // Ensure the context is cancelled when the main function exits
	ticker1 := time.NewTicker(1 * time.Hour)
	defer ticker1.Stop()

	go func() {
		// This loop waits for the ticker to "tick"
		for range ticker.C {
			log.Println("Hourly poll initiated...")
			// Create a new, short-lived context for each recurring API call.
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			if err := rate_fetcher.CryptoRate(ctx); err != nil {
				log.Printf("Error during hourly rate polling: %v", err)
			}
			cancel() // Clean up context after the call is done
		}
	}()

	FetchFiatRateHandler := httptransport.NewServer(
		transport.FetchFiatRateEndpoint(svc),
		transport.DecodeFetchFiatRateRequest,
		transport.EncodeFetchFiatRateResponse,
	)
	convertHandler := httptransport.NewServer(
		transport.ConvertFiatEndPoint(svc),
		transport.DecodeConvertFiatRequest,
		transport.EncodeConvertFiatResponse,
	)
	historyHandler := httptransport.NewServer(
		transport.HistoryEndpoint(svc),
		transport.DecodeHistoryRequest,
		transport.EncodeHistoryResponse,
	)

	http.Handle("/fetch", FetchFiatRateHandler)
	http.Handle("/convert", convertHandler)
	http.Handle("/history", historyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
