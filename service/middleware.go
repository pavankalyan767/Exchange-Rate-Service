package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/log"
	"github.com/pavankalyan767/exchange-rate-service/types"
)

type loggingMiddleware struct {
	logger log.Logger
	next   ExchangeRateService
}

func NewLoggingMiddleware(logger log.Logger, next ExchangeRateService) ExchangeRateService {
	// We return a pointer to the middleware struct, which satisfies the interface.
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

// FetchFiatRate implements the ExchangeRateService interface.
// It logs the call and delegates to the next service.
func (mw *loggingMiddleware) FetchRate(ctx context.Context, req *types.FetchRateRequest) (output float64, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "fetch_fiat_rate",
			"input_base", req.BaseCurrency,
			"input_target", req.TargetCurrency,
			"output_rate", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	// Call the next service in the chain.
	output, err = mw.next.FetchRate(ctx, req)
	return
}

// Convert implements the ExchangeRateService interface.
// It logs the call and delegates to the next service.
func (mw *loggingMiddleware) Convert(ctx context.Context, req *types.ConvertRequest) (output float64, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "convert",
			"BaseCurrency", req.BaseCurrency,
			"TargetCurrency", req.TargetCurrency,
			"input_amount", req.Amount,
			"output_rate", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Convert(ctx, req)
	return
}

// History implements the ExchangeRateService interface.
// The return type is now float64, so we've updated the logic to match.
func (mw *loggingMiddleware) History(ctx context.Context, req *types.HistoryRequest) (output map[string]float64, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "history",
			"input_base", req.BaseCurrency,
			"input_target", req.TargetCurrency,
			"input_from", req.From,
			"input_to", req.To,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.History(ctx, req)
	return
}

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           ExchangeRateService
}

func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram,countResult metrics.Histogram, next ExchangeRateService) ExchangeRateService {
	// We return a pointer to the middleware struct, which satisfies the interface.
	return &instrumentingMiddleware{
		requestCount: requestCount,
		requestLatency: requestLatency,
		countResult: countResult,
		next: next,

	}
}
func (mw *instrumentingMiddleware) FetchRate(ctx context.Context, req *types.FetchRateRequest) (output float64, err error) {

	defer func(begin time.Time) {
		lvs := []string{"method", "FetchRate", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.FetchRate(ctx,req)
	return
}

// Convert implements the ExchangeRateService interface.
// It logs the call and delegates to the next service.
func (mw *instrumentingMiddleware) Convert(ctx context.Context, req *types.ConvertRequest) (output float64, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Convert", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Convert(ctx,req)
	return
}

// History implements the ExchangeRateService interface.
// The return type is now float64, so we've updated the logic to match.
func (mw *instrumentingMiddleware) History(ctx context.Context, req *types.HistoryRequest) (output map[string]float64, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "History", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.History(ctx,req)
	return
}
