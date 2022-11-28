package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-data-history-svc/internal/getOneHistorical"
)

func MakeGetOneHistoricalEndpoints(c getOneHistorical.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetOneHistoricalInternalRequest)
		resp, err := c.GetOneHistoricalService(req.ctx, req.Id)
		return GetOneHistoricalInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetOneHistoricalInternalResponse struct {
	Response interface{}
	Err      error
}

type GetOneHistoricalInternalRequest struct {
	Id  string `json:"id" example:"1" validate:"nonzero"`
	ctx context.Context
}
