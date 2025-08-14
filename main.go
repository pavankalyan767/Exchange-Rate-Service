package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/joho/godotenv"
	service"github.com/pavankalyan767/exchange-rate-service/service"
	"github.com/pavankalyan767/exchange-rate-service/transport"
)

func main() {

	svc := service.ExchangeRateServiceImpl{}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fetchExchangeRateHandler := httptransport.NewServer(
		FetchExchangeRateEndPoint(svc),
		decodeFetchExchangeRateRequest,
		encodeResponse,
	)

	http.Handle("/fetch_exchange_rate", fetchExchangeRateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
