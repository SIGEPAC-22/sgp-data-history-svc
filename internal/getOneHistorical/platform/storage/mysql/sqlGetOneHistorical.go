package mysql

import "time"

type SqlGetOneHistorical struct {
	idPatient     int        `db:"his_id_patient"`
	idPatientFile int        `db:"his_id_file_patient"`
	fullName      string     `db:"his_first_name"`
	fullLastName  string     `db:"his_second_name"`
	admissionDate time.Time  `db:"his_admission_date"`
	highDate      *time.Time `db:"his_high_date"`
	lowDate       *time.Time `db:"his_low_date"`
}
