package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jieqiboh/sothea_backend/domain"
	"github.com/jieqiboh/sothea_backend/models"
	_ "github.com/lib/pq"
)

type postgresPatientRepository struct {
	Conn *sql.DB
}

// NewPostgresPatientRepository will create an object that represent the patient.Repository interface
func NewPostgresPatientRepository(conn *sql.DB) domain.PatientRepository {
	return &postgresPatientRepository{conn}
}

// GetPatientByID returns a Patient struct based on ID. Only guaranteed field is Admin
func (p *postgresPatientRepository) GetPatientByID(ctx context.Context, id int64) (res *domain.Patient, err error) {
	// Create a helper function for preparing failure results.
	fail := func(err error) (*domain.Patient, error) {
		return nil, fmt.Errorf("GetPatientByID: %v", err)
	}

	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	rows := tx.QueryRowContext(ctx, "SELECT * FROM admin WHERE id = $1;", id)
	admin := models.Admin{}
	err = rows.Scan(&admin.ID, &admin.FamilyGroup, &admin.RegDate, &admin.Name, &admin.Age, &admin.Gender)
	if err != nil {
		return fail(err)
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM pastmedicalhistory WHERE pastmedicalhistory.id = $1;", id)
	pastmedicalhistory := &models.PastMedicalHistory{}
	err = rows.Scan(
		&pastmedicalhistory.ID,
		&pastmedicalhistory.Tuberculosis,
		&pastmedicalhistory.Diabetes,
		&pastmedicalhistory.Hypertension,
		&pastmedicalhistory.Hyperlipidemia,
		&pastmedicalhistory.ChronicJointPains,
	)
	if errors.Is(sql.ErrNoRows, err) {
		pastmedicalhistory = nil
	} else if err != nil {
		return fail(err)
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM socialhistory WHERE socialhistory.id = $1;", id)
	socialhistory := &models.SocialHistory{}
	err = rows.Scan(
		&socialhistory.ID,
		&socialhistory.PastSmokingHistory,
		&socialhistory.NumberOfYears,
		&socialhistory.CurrentSmokingHistory,
	)
	if errors.Is(sql.ErrNoRows, err) {
		socialhistory = nil
	} else if err != nil {
		return fail(err)
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM vitalstatistics WHERE vitalstatistics.id = $1;", id)
	vitalstatistics := &models.VitalStatistics{}
	err = rows.Scan(
		&vitalstatistics.ID,
		&vitalstatistics.Temperature,
		&vitalstatistics.SpO2,
	)
	if errors.Is(sql.ErrNoRows, err) {
		vitalstatistics = nil
	} else if err != nil {
		return fail(err)
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM heightandweight WHERE heightandweight.id = $1;", id)
	heightandweight := &models.HeightAndWeight{}
	err = rows.Scan(
		&heightandweight.ID,
		&heightandweight.Height,
		&heightandweight.Weight,
	)
	if errors.Is(sql.ErrNoRows, err) {
		heightandweight = nil
	} else if err != nil {
		return fail(err)
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM visualacuity WHERE visualacuity.id = $1;", id)
	visualacuity := &models.VisualAcuity{}
	err = rows.Scan(
		&visualacuity.ID,
		&visualacuity.LEyeVision,
		&visualacuity.REyeVision,
	)
	if errors.Is(sql.ErrNoRows, err) {
		visualacuity = nil
	} else if err != nil {
		return fail(err)
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM doctorsconsultation WHERE doctorsconsultation.id = $1;", id)
	doctorsconsultation := &models.DoctorsConsultation{}
	err = rows.Scan(
		&doctorsconsultation.ID,
		&doctorsconsultation.Healthy,
		&doctorsconsultation.ConsultationNotes,
		&doctorsconsultation.ReferralNeeded,
	)
	if errors.Is(sql.ErrNoRows, err) {
		doctorsconsultation = nil
	} else if err != nil {
		return fail(err)
	}

	patient := domain.Patient{
		Admin:               &admin,
		PastMedicalHistory:  pastmedicalhistory,
		SocialHistory:       socialhistory,
		VitalStatistics:     vitalstatistics,
		HeightAndWeight:     heightandweight,
		VisualAcuity:        visualacuity,
		DoctorsConsultation: doctorsconsultation,
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return &patient, nil
}

// InsertPatient inserts a Patient and returns the new id if successful. Only required field is Admin
func (p *postgresPatientRepository) InsertPatient(ctx context.Context, patient *domain.Patient) (int64, error) {
	// Create a helper function for preparing failure results.
	fail := func(err error) (int64, error) {
		return -1, fmt.Errorf("InsertPatient: %v", err)
	}

	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	admin := patient.Admin
	pastmedicalhistory := patient.PastMedicalHistory
	socialhistory := patient.SocialHistory
	vitalstatistics := patient.VitalStatistics
	heightandweight := patient.HeightAndWeight
	visualacuity := patient.VisualAcuity
	doctorsconsultation := patient.DoctorsConsultation

	var patientid int64
	if admin == nil {
		return fail(errors.New("Admin field cannot be nil"))
	}
	rows := tx.QueryRowContext(ctx, `INSERT INTO admin (family_group, reg_date, name, age, gender) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		admin.FamilyGroup, admin.RegDate.Format("2006-01-02"), admin.Name, admin.Age, admin.Gender)
	err = rows.Scan(&patientid)
	if err != nil {
		return fail(err)
	}
	if pastmedicalhistory != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO pastmedicalhistory (id, tuberculosis, diabetes, hypertension, hyperlipidemia, chronicjointpains) VALUES ($1, $2, $3, $4, $5, $6)`,
			patientid, pastmedicalhistory.Tuberculosis, pastmedicalhistory.Diabetes, pastmedicalhistory.Hypertension, pastmedicalhistory.Hyperlipidemia, pastmedicalhistory.ChronicJointPains)
		if err != nil {
			return fail(err)
		}
	}
	if socialhistory != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO socialhistory (id, past_smoking_history, no_of_years, current_smoking_history) VALUES ($1, $2, $3, $4)`,
			patientid, socialhistory.PastSmokingHistory, socialhistory.NumberOfYears, socialhistory.CurrentSmokingHistory)
		if err != nil {
			return fail(err)
		}
	}
	if vitalstatistics != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO vitalstatistics (id, temperature, SpO2) VALUES ($1, $2, $3)`,
			patientid, vitalstatistics.Temperature, vitalstatistics.SpO2)
		if err != nil {
			return fail(err)
		}
	}
	if heightandweight != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO heightandweight (id, height, weight) VALUES ($1, $2, $3) RETURNING id`,
			patientid, heightandweight.Height, heightandweight.Weight)
		if err != nil {
			return fail(err)
		}
	}
	if visualacuity != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO visualacuity (id, l_eyevision, r_eyevision) VALUES ($1, $2, $3)`,
			patientid, visualacuity.LEyeVision, visualacuity.REyeVision)
		if err != nil {
			return fail(err)
		}
	}
	if doctorsconsultation != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO doctorsconsultation (id, healthy, consultation_notes, referral_needed) VALUES ($1, $2, $3, $4) RETURNING id`,
			patientid, doctorsconsultation.Healthy, doctorsconsultation.ConsultationNotes, doctorsconsultation.ReferralNeeded)
		if err != nil {
			return fail(err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	return patientid, nil
}

func (p *postgresPatientRepository) DeletePatientByID(ctx context.Context, id int64) (int64, error) {
	// Create a helper function for preparing failure results.
	fail := func(err error) (int64, error) {
		return -1, fmt.Errorf("DeletePatientByID: %v", err)
	}

	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM pastmedicalhistory WHERE id = $1", id)
	if err != nil {
		return fail(err)
	}
	_, err = tx.Exec("DELETE FROM socialhistory WHERE id = $1", id)
	if err != nil {
		return fail(err)
	}
	_, err = tx.Exec("DELETE FROM vitalstatistics WHERE id = $1", id)
	if err != nil {
		return fail(err)
	}
	_, err = tx.Exec("DELETE FROM heightandweight WHERE id = $1", id)
	if err != nil {
		return fail(err)
	}
	_, err = tx.Exec("DELETE FROM visualacuity WHERE id = $1", id)
	if err != nil {
		return fail(err)
	}
	_, err = tx.Exec("DELETE FROM doctorsconsultation WHERE id = $1", id)
	if err != nil {
		return fail(err)
	}
	_, err = tx.Exec("DELETE FROM admin WHERE id = $1", id)
	if err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	return id, nil
}

// UpdatePatientByID updates an already existing patient, filling out or overriding any of its fields
func (p *postgresPatientRepository) UpdatePatientByID(ctx context.Context, id int64, patient *domain.Patient) (int64, error) {
	// Checks that a patient exists by searching for admin field
	// Then for each non-nil field in patient, updates it
	// Create a helper function for preparing failure results.
	fail := func(err error) (int64, error) {
		return -1, fmt.Errorf("InsertPatient: %v", err)
	}

	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	admin := patient.Admin
	pastmedicalhistory := patient.PastMedicalHistory
	socialhistory := patient.SocialHistory
	vitalstatistics := patient.VitalStatistics
	heightandweight := patient.HeightAndWeight
	visualacuity := patient.VisualAcuity
	doctorsconsultation := patient.DoctorsConsultation

	// Check that patient exists
	rows := tx.QueryRowContext(ctx, "SELECT * FROM admin WHERE id = $1;", id)
	prevAdmin := models.Admin{}
	err = rows.Scan(&prevAdmin.ID, &prevAdmin.FamilyGroup, &prevAdmin.RegDate, &prevAdmin.Name, &prevAdmin.Age, &prevAdmin.Gender)
	if err != nil {
		return fail(err)
	}
	if admin != nil { // Update admin
		_, err = tx.ExecContext(ctx, `UPDATE admin SET family_group = $1, reg_date = $2, name = $3, age = $4, gender = $5 WHERE id = $6`,
			admin.FamilyGroup, admin.RegDate.Format("2006-01-02"), admin.Name, admin.Age, admin.Gender, id)
		if err != nil {
			return fail(err)
		}
	}
	if pastmedicalhistory != nil { // Update pastmedicalhistory
		_, err = tx.ExecContext(ctx, `INSERT INTO pastmedicalhistory (id, tuberculosis, diabetes, hypertension, hyperlipidemia, chronicjointpains) 
											  VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT(id)  DO UPDATE SET
												tuberculosis = $2, 
												diabetes = $3,
												hypertension = $4,
												hyperlipidemia = $5,
												chronicjointpains = $6`,
			id, pastmedicalhistory.Tuberculosis, pastmedicalhistory.Diabetes, pastmedicalhistory.Hypertension, pastmedicalhistory.Hyperlipidemia, pastmedicalhistory.ChronicJointPains)
		if err != nil {
			return fail(err)
		}
	}
	if socialhistory != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO socialhistory (id, past_smoking_history, no_of_years, current_smoking_history) 
		VALUES ($1, $2, $3, $4) 
		ON CONFLICT(id) 
		DO UPDATE SET
			past_smoking_history = $2,
			no_of_years = $3,
			current_smoking_history = $4`,
			id, socialhistory.PastSmokingHistory, socialhistory.NumberOfYears, socialhistory.CurrentSmokingHistory)

		if err != nil {
			return fail(err)
		}
	}
	if vitalstatistics != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO vitalstatistics (id, temperature, spO2) 
		VALUES ($1, $2, $3) 
		ON CONFLICT(id) 
		DO UPDATE SET
			temperature = $2,
			spO2 = $3`,
			id, vitalstatistics.Temperature, vitalstatistics.SpO2)

		if err != nil {
			return fail(err)
		}
	}
	if heightandweight != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO heightandweight (id, height, weight) 
		VALUES ($1, $2, $3) 
		ON CONFLICT(id) 
		DO UPDATE SET
			height = $2,
			weight = $3`,
			id, heightandweight.Height, heightandweight.Weight)

		if err != nil {
			return fail(err)
		}
	}
	if visualacuity != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO visualacuity (id, l_eyevision, r_eyevision) 
		VALUES ($1, $2, $3) 
		ON CONFLICT(id) 
		DO UPDATE SET
			l_eyevision = $2,
			r_eyevision = $3`,
			id, visualacuity.LEyeVision, visualacuity.REyeVision)

		if err != nil {
			return fail(err)
		}
	}
	if doctorsconsultation != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO doctorsconsultation (id, healthy, consultation_notes, referral_needed) 
		VALUES ($1, $2, $3, $4) 
		ON CONFLICT(id) 
		DO UPDATE SET
			healthy = $2,
			consultation_notes = $3,
			referral_needed = $4`,
			id, doctorsconsultation.Healthy, doctorsconsultation.ConsultationNotes, doctorsconsultation.ReferralNeeded)

		if err != nil {
			return fail(err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	return id, nil
}

func (p *postgresPatientRepository) GetAllFromAdmin(ctx context.Context) ([]models.Admin, error) {
	var rows *sql.Rows
	result := make([]models.Admin, 0)
	query := "SELECT * FROM admin"
	rows, err := p.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		admin := models.Admin{}
		err = rows.Scan(&admin.ID, &admin.FamilyGroup, &admin.RegDate, &admin.Name, &admin.Age, &admin.Gender)

		if err != nil {
			return nil, err
		}
		result = append(result, admin)
	}

	return result, nil
}
