package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-data-history-svc/internal/getHistorical"
	"sgp-data-history-svc/kit/constants"
	"time"
)

type GetHistoricalRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetHistoricalRepo(db *sql.DB, log log.Logger) *GetHistoricalRepo {
	return &GetHistoricalRepo{db, log}
}

func (g *GetHistoricalRepo) GetHistoricalRepository(ctx context.Context) ([]getHistorical.GetHistoricalResponse, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	//id := config.GetInt("app-properties.getPatient.idStatusActive")

	rows, errDB := g.db.QueryContext(ctx, "SELECT his_id_patient, his_id_file_patient, concat(his_first_name,' ', his_second_name) as full_name, concat(his_first_last_name,' ', his_second_last_name) as full_last_name, his_admission_date, his_high_date, his_low_date FROM his_historical;")
	if errDB != nil {
		g.log.Log("Error while trying to get information for historical", constants.UUID, ctx.Value(constants.UUID))
		return []getHistorical.GetHistoricalResponse{}, errDB
	}
	defer rows.Close()
	var resp []getHistorical.GetHistoricalResponse
	for rows.Next() {
		var respDB SqlGetHistorical
		if err := rows.Scan(&respDB.idPatient, &respDB.idPatientFile, &respDB.fullName, &respDB.fullLastName, &respDB.admissionDate, &respDB.highDate, &respDB.lowDate); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getHistorical.GetHistoricalResponse{}, err
		}
		resp = append(resp, getHistorical.GetHistoricalResponse{
			IdPatient:     respDB.idPatient,
			IdPatientFile: respDB.idPatientFile,
			FirstName:     respDB.fullName,
			LastName:      respDB.fullLastName,
			AdmissionDate: respDB.admissionDate.Format(config.GetString("app-properties.getHistorical.dateAdmission-Format")),
			HighDate:      transformerPointer(respDB.highDate),
			LowDate:       transformerPointer(respDB.lowDate),
		})
	}
	if len(resp) == 0 {
		g.log.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
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
