package getOneHistorical

import "context"

type Repository interface {
	GetOneHistoricalRepository(ctx context.Context, id string) ([]GetOneHistoricalResponse, error)
}

type Service interface {
	GetOneHistoricalService(ctx context.Context, id string) ([]GetOneHistoricalResponse, error)
}

type GetOneComorbidityRequest struct {
	Id string `json:"id"`
}

type GetOneHistoricalResponse struct {
	IdPatient     int    `json:"idPatient"`
	IdPatientFile int    `json:"idPatientFile"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	AdmissionDate string `json:"admissionDate"`
	HighDate      string `json:"highDate"`
	LowDate       string `json:"lowDate"`
}
