package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-data-history-svc/internal/getHistorical"
)

func MakeGetHistoricalEndpoints(c getHistorical.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetHistoricalInternalRequest)
		resp, err := c.GetHistoricalService(req.ctx)
		return GetHistoricalInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetHistoricalInternalResponse struct {
	Response interface{}
	Err      error
}

type GetHistoricalInternalRequest struct {
	ctx context.Context
}
