package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/joho/godotenv"
	"github.com/pavankalyan767/exchange-rate-service/cache"
	"github.com/pavankalyan767/exchange-rate-service/client"
	service "github.com/pavankalyan767/exchange-rate-service/service"
	"github.com/pavankalyan767/exchange-rate-service/transport"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a logger with a timestamp and caller information.
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "exchange-rate-service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "exchange-rate-service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "exchange-rate-service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	// Load environment variables from .env file.
	err := godotenv.Load(".env")
	if err != nil {
		logger.Log("Error", "failed to load .env file", "err", err)
		// Exit if environment variables cannot be loaded, as the service won't function.
		os.Exit(1)
	}

	// Read API keys and URLs from environment variables.
	fiatapikey := os.Getenv("FIAT_API_KEY")
	if fiatapikey == "" {
		logger.Log("Error", "FIAT_API_KEY environment variable is not set. Exiting.")
		os.Exit(1)
	}
	cryptoapikey := os.Getenv("CRYPTO_API_KEY")
	if cryptoapikey == "" {
		logger.Log("Warning", "CRYPTO_API_KEY environment variable is not set. Crypto-related requests will not be available.")
	}
	fiatUrl := os.Getenv("FIAT_API_URL")
	if fiatUrl == "" {
		logger.Log("Error", "FIAT_API_URL environment variable is not set. Exiting.")
		os.Exit(1)
	}
	cryptoUrl := os.Getenv("CRYPTO_API_URL")
	if cryptoUrl == "" {
		logger.Log("Warning", "CRYPTO_API_URL environment variable is not set. Crypto-related requests will not be available.")
	}

	// Initialize the API client and caches.
	apiClient := client.NewAPIClient(fiatUrl, cryptoUrl, fiatapikey, cryptoapikey, logger)
	fiatCache := cache.NewCache(5*time.Minute, 10*time.Minute,logger)
	cryptoCache := cache.NewCache(5*time.Minute, 10*time.Minute,logger)

	// Initialize the core service.
	var svc service.ExchangeRateService
	svc = service.NewExchangeRateServiceImpl(fiatCache, cryptoCache)

	svc = service.NewLoggingMiddleware(logger, svc)
	svc = service.NewInstrumentingMiddleware(requestCount, requestLatency, countResult, svc)

	// Initialize the rate fetcher.
	rate_fetcher := service.NewRateFetcher(apiClient, fiatCache, cryptoCache,logger)

	// --- Polling Logic ---
	// Create a single context to manage all background goroutines.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // This will cancel the context when main() returns.

	// Perform an initial fetch for fiat and crypto rates.
	if err := rate_fetcher.LiveRate(ctx); err != nil {
		logger.Log("Error", fmt.Sprintf("failed to fetch initial fiat rates: %v", err))
	}
	if err := rate_fetcher.CryptoRate(ctx); err != nil {
		logger.Log("Error", fmt.Sprintf("failed to fetch initial crypto rates: %v", err))
	}
	logger.Log("message", "Initial fetch of live rates complete.")

	// Start a single background goroutine for hourly polling of both rate types.
	if err := rate_fetcher.HistoricalRate(ctx); err != nil {
		logger.Log("Failed to fetch historical rates on startup: %v", err)
	}
	if err := rate_fetcher.CryptoRate(ctx); err != nil {
		logger.Log("Error", fmt.Sprintf("error during hourly crypto rate polling: %v", err))
	}

	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			logger.Log("message", "Hourly poll initiated...")
			// Use the main context for the API calls.
			if err := rate_fetcher.LiveRate(ctx); err != nil {
				logger.Log("Error", fmt.Sprintf("error during hourly fiat rate polling: %v", err))
			}
			if err := rate_fetcher.CryptoRate(ctx); err != nil {
				logger.Log("Error", fmt.Sprintf("error during hourly crypto rate polling: %v", err))
			}
		}
	}()

	// --- Endpoint and Middleware Setup ---

	// Create all endpoints from the service.
	endpoints := transport.MakeEndpoints(svc)

	// Apply logging middleware to each endpoint.
	endpoints.HistoryEndpoint = transport.LoggingMiddleware(
		log.With(logger, "method", "history"),
	)(endpoints.HistoryEndpoint)

	endpoints.FetchEndpoint = transport.LoggingMiddleware(
		log.With(logger, "method", "fetch"),
	)(endpoints.FetchEndpoint)

	endpoints.ConvertEndpoint = transport.LoggingMiddleware(
		log.With(logger, "method", "convert"),
	)(endpoints.ConvertEndpoint)

	// --- HTTP Handlers and Server ---

	// Create HTTP handlers for each endpoint.
	fetchHandler := httptransport.NewServer(
		endpoints.FetchEndpoint,
		transport.DecodeFetchRateRequest,
		transport.EncodeResponse,
	)

	convertHandler := httptransport.NewServer(
		endpoints.ConvertEndpoint,
		transport.DecodeConvertRequest,
		transport.EncodeResponse,
	)

	historyHandler := httptransport.NewServer(
		endpoints.HistoryEndpoint,
		transport.DecodeHistoryRequest,
		transport.EncodeResponse,
	)

	// Register handlers with their paths.
	http.Handle("/fetch", fetchHandler)
	http.Handle("/convert", convertHandler)
	http.Handle("/history", historyHandler)
	http.Handle("/metrics", promhttp.Handler())

	// Start the HTTP server.
	logger.Log("message", "HTTP server listening", "port", "8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Log("Error", "server failed to start", "err", err)
		os.Exit(1)
	}
}
