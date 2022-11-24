package getOneHistoricalService

import (
	"context"
	"github.com/go-kit/log"
	"sgp-data-history-svc/internal/getOneHistorical"
	"sgp-data-history-svc/kit/constants"
)

type Service struct {
	RepoDB getOneHistorical.Repository
	logger log.Logger
}

func NewService(repoBD getOneHistorical.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoBD, logger: logger}
}

func (s *Service) GetOneHistoricalService(ctx context.Context, id string) ([]getOneHistorical.GetOneHistoricalResponse, error) {
	s.logger.Log("Start Endpoint GetHistorical", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetOneHistoricalRepository(ctx, id)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []getOneHistorical.GetOneHistoricalResponse{}, err
	}
	return resp, nil
}
