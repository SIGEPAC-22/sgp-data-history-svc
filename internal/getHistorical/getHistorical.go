package getHistorical

import "context"

type Repository interface {
	GetHistoricalRepository(ctx context.Context) ([]GetHistoricalResponse, error)
}

type Service interface {
	GetHistoricalService(ctx context.Context) ([]GetHistoricalResponse, error)
}

type GetHistoricalResponse struct {
	IdPatient     int    `json:"idPatient"`
	IdPatientFile int    `json:"idPatientFile"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	AdmissionDate string `json:"admissionDate"`
	HighDate      string `json:"highDate"`
	LowDate       string `json:"lowDate"`
}
