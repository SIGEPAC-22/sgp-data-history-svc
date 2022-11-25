package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-data-history-svc/internal/getOneHistorical"
	"sgp-data-history-svc/kit/constants"
	"time"
)

type GetOneHistoricalRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetOneHistoricalRepo(db *sql.DB, log log.Logger) *GetOneHistoricalRepo {
	return &GetOneHistoricalRepo{db, log}
}

func (g *GetOneHistoricalRepo) GetOneHistoricalRepository(ctx context.Context, id string) (getOneHistorical.GetOneHistoricalResponse, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)

	rows := g.db.QueryRowContext(ctx, "SELECT his_id_patient, his_id_file_patient, his_first_name, his_second_name, his_admission_date, his_high_date, his_low_date FROM his_historical WHERE his_id_patient = ?", id)
	var respDB SqlGetOneHistorical
	if err := rows.Scan(&respDB.idPatient, &respDB.idPatientFile, &respDB.firstName, &respDB.lastName, &respDB.admissionDate, &respDB.highDate, &respDB.lowDate); err != nil {
		g.log.Log("Data not found", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return getOneHistorical.GetOneHistoricalResponse{}, errors.New("Data not found")
	}
	resp := getOneHistorical.GetOneHistoricalResponse{
		IdPatient:     respDB.idPatient,
		IdPatientFile: respDB.idPatientFile,
		FirstName:     respDB.firstName,
		LastName:      respDB.lastName,
		AdmissionDate: respDB.admissionDate.Format(config.GetString("app-properties.getHistorical.dateAdmission-Format")),
		HighDate:      transformerPointer(respDB.highDate),
		LowDate:       transformerPointer(respDB.lowDate),
	}
	return resp, nil
}

func transformerPointer(date *time.Time) string {

	if date != nil {
		var dateConverter string

		config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
		dateConverter = date.Format(config.GetString("app-properties.getHistorical.dateAdmission-Format"))
		return dateConverter
	} else {
		return "not available"
	}

}
