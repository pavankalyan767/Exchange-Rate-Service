package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/joho/godotenv"
	service "github.com/pavankalyan767/exchange-rate-service/service"
	"github.com/pavankalyan767/exchange-rate-service/transport"
)

func main() {

	svc := service.ExchangeRateServiceImpl{}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	http.Handle("/fetch_exchange_rate", fetchExchangeRateHandler)
	http.Handle("/convert", convertHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
