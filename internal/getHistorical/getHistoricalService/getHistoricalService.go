package getHistoricalService

import (
	"context"
	"github.com/go-kit/log"
	"sgp-data-history-svc/internal/getHistorical"
	"sgp-data-history-svc/kit/constants"
)

type Service struct {
	RepoDB getHistorical.Repository
	logger log.Logger
}

func NewService(repoBD getHistorical.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoBD, logger: logger}
}

func (s *Service) GetHistoricalService(ctx context.Context) ([]getHistorical.GetHistoricalResponse, error) {
	s.logger.Log("Start Endpoint GetHistorical", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetHistoricalRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []getHistorical.GetHistoricalResponse{}, err
	}
	return resp, nil

}
