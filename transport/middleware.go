package transport

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {

				logger.Log(
					"transport_level_log", "true",
					"duration_ms", time.Since(begin).Milliseconds(),
					"err", err,
				)
			}(time.Now()) // Pass the current time to the deferred function.

			response, err = next(ctx, request)
			return response, err
		}
	}
}
