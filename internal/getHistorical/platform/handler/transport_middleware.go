package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/log"
)

type Middleware func(endpoint endpoint.Endpoint) endpoint.Endpoint

func GetHistoricalTransportMiddleware(log kitlog.Logger) Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req := request.(GetHistoricalInternalRequest)
			defer log.Log("Process Finished", "Request", req)
			return e(ctx, request)
		}
	}
}
