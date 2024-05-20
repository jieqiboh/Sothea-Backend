package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jieqiboh/sothea_backend/entities"
	_ "github.com/lib/pq"
)

type postgresPatientRepository struct {
	Conn *sql.DB
}

// NewPostgresPatientRepository will create an object that represent the patient.Repository interface
func NewPostgresPatientRepository(conn *sql.DB) entities.PatientRepository {
	return &postgresPatientRepository{conn}
}

// GetPatientByID returns a Patient struct based on ID. Only guaranteed field is Admin
func (p *postgresPatientRepository) GetPatientByID(ctx context.Context, id int32) (res *entities.Patient, err error) {
	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	rows := tx.QueryRowContext(ctx, "SELECT * FROM admin WHERE id = $1;", id)
	admin := entities.Admin{}
	err = rows.Scan(
		&admin.ID,
		&admin.FamilyGroup,
		&admin.RegDate,
		&admin.Name,
		&admin.KhmerName,
		&admin.Dob,
		&admin.Age,
		&admin.Gender,
		&admin.Village,
		&admin.ContactNo,
		&admin.Pregnant,
		&admin.LastMenstrualPeriod,
		&admin.DrugAllergies,
		&admin.SentToID,
		&admin.Photo,
	)
	if err != nil { // no admin found
		return nil, entities.ErrPatientNotFound
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM pastmedicalhistory WHERE pastmedicalhistory.id = $1;", id)
	pastmedicalhistory := &entities.PastMedicalHistory{}
	err = rows.Scan(
		&pastmedicalhistory.ID,
		&pastmedicalhistory.Tuberculosis,
		&pastmedicalhistory.Diabetes,
		&pastmedicalhistory.Hypertension,
		&pastmedicalhistory.Hyperlipidemia,
		&pastmedicalhistory.ChronicJointPains,
		&pastmedicalhistory.ChronicMuscleAches,
		&pastmedicalhistory.SexuallyTransmittedDisease,
		&pastmedicalhistory.SpecifiedSTDs,
		&pastmedicalhistory.Others,
	)
	if errors.Is(sql.ErrNoRows, err) { // no pastmedicalhistory found
		pastmedicalhistory = nil
	} else if err != nil { // unknown error
		return nil, err
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM socialhistory WHERE socialhistory.id = $1;", id)
	socialhistory := &entities.SocialHistory{}
	err = rows.Scan(
		&socialhistory.ID,
		&socialhistory.PastSmokingHistory,
		&socialhistory.NumberOfYears,
		&socialhistory.CurrentSmokingHistory,
		&socialhistory.CigarettesPerDay,
		&socialhistory.AlcoholHistory,
		&socialhistory.HowRegular,
	)
	if errors.Is(sql.ErrNoRows, err) { // no socialhistory found
		socialhistory = nil
	} else if err != nil { // unknown error
		return nil, err
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM vitalstatistics WHERE vitalstatistics.id = $1;", id)
	vitalstatistics := &entities.VitalStatistics{}
	err = rows.Scan(
		&vitalstatistics.ID,
		&vitalstatistics.Temperature,
		&vitalstatistics.SpO2,
		&vitalstatistics.SystolicBP1,
		&vitalstatistics.DiastolicBP1,
		&vitalstatistics.SystolicBP2,
		&vitalstatistics.DiastolicBP2,
		&vitalstatistics.AverageSystolicBP,
		&vitalstatistics.AverageDiastolicBP,
		&vitalstatistics.HR1,
		&vitalstatistics.HR2,
		&vitalstatistics.AverageHR,
		&vitalstatistics.RandomBloodGlucoseMmolL,
		&vitalstatistics.RandomBloodGlucoseMmolLp,
	)
	if errors.Is(sql.ErrNoRows, err) { // no vitalstatistics found
		vitalstatistics = nil
	} else if err != nil { // unknown error
		return nil, err
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM heightandweight WHERE heightandweight.id = $1;", id)
	heightandweight := &entities.HeightAndWeight{}
	err = rows.Scan(
		&heightandweight.ID,
		&heightandweight.Height,
		&heightandweight.Weight,
		&heightandweight.BMI,
		&heightandweight.BMIAnalysis,
		&heightandweight.PaedsHeight,
		&heightandweight.PaedsWeight,
	)
	if errors.Is(sql.ErrNoRows, err) { // no heightandweight found
		heightandweight = nil
	} else if err != nil { // unknown error
		return nil, err
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM visualacuity WHERE visualacuity.id = $1;", id)
	visualacuity := &entities.VisualAcuity{}
	err = rows.Scan(
		&visualacuity.ID,
		&visualacuity.LEyeVision,
		&visualacuity.REyeVision,
		&visualacuity.AdditionalIntervention,
	)
	if errors.Is(sql.ErrNoRows, err) { // no visualacuity found
		visualacuity = nil
	} else if err != nil { // unknown error
		return nil, err
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM doctorsconsultation WHERE doctorsconsultation.id = $1;", id)
	doctorsconsultation := &entities.DoctorsConsultation{}
	err = rows.Scan(
		&doctorsconsultation.ID,
		&doctorsconsultation.Healthy,
		&doctorsconsultation.Msk,
		&doctorsconsultation.Cvs,
		&doctorsconsultation.Respi,
		&doctorsconsultation.Gu,
		&doctorsconsultation.Git,
		&doctorsconsultation.Eye,
		&doctorsconsultation.Derm,
		&doctorsconsultation.Others,
		&doctorsconsultation.ConsultationNotes,
		&doctorsconsultation.Diagnosis,
		&doctorsconsultation.Treatment,
		&doctorsconsultation.ReferralNeeded,
		&doctorsconsultation.ReferralLoc,
		&doctorsconsultation.Remarks,
	)
	if errors.Is(sql.ErrNoRows, err) { // no doctorsconsultation found
		doctorsconsultation = nil
	} else if err != nil { // unknown error
		return nil, err
	}

	patient := entities.Patient{
		Admin:               &admin,
		PastMedicalHistory:  pastmedicalhistory,
		SocialHistory:       socialhistory,
		VitalStatistics:     vitalstatistics,
		HeightAndWeight:     heightandweight,
		VisualAcuity:        visualacuity,
		DoctorsConsultation: doctorsconsultation,
	}

	if err = tx.Commit(); err != nil { // commit transaction
		return nil, err
	}

	return &patient, nil
}

// InsertPatient inserts a Patient and returns the new id if successful. Only required field is Admin
func (p *postgresPatientRepository) InsertPatient(ctx context.Context, patient *entities.Patient) (int32, error) {
	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil { // error starting transaction
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

	var patientid int32
	if admin == nil { // no admin field
		return -1, entities.ErrMissingAdminCategory
	}
	rows := tx.QueryRowContext(ctx, `INSERT INTO admin (family_group, reg_date, name, khmer_name, dob, age, gender, village, 
	contact_no, pregnant, last_menstrual_period, drug_allergies, sent_to_id, photo) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id`,
		admin.FamilyGroup, admin.RegDate, admin.Name, admin.KhmerName, admin.Dob, admin.Age, admin.Gender, admin.Village, admin.ContactNo,
		admin.Pregnant, admin.LastMenstrualPeriod, admin.DrugAllergies, admin.SentToID, admin.Photo)
	err = rows.Scan(&patientid)
	if err != nil { // error inserting admin
		return -1, err
	}
	if pastmedicalhistory != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO pastmedicalhistory (id, tuberculosis, diabetes, hypertension, 
		hyperlipidemia, chronic_joint_pains, chronic_muscle_aches, sexually_transmitted_disease, specified_stds, others) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
			patientid, pastmedicalhistory.Tuberculosis, pastmedicalhistory.Diabetes, pastmedicalhistory.Hypertension,
			pastmedicalhistory.Hyperlipidemia, pastmedicalhistory.ChronicJointPains, pastmedicalhistory.ChronicMuscleAches,
			pastmedicalhistory.SexuallyTransmittedDisease, pastmedicalhistory.SpecifiedSTDs, pastmedicalhistory.Others)
		if err != nil {
			return -1, err
		}
	}
	if socialhistory != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO socialhistory (id, past_smoking_history, no_of_years, 
		current_smoking_history, cigarettes_per_day, alcohol_history, how_regular) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			patientid, socialhistory.PastSmokingHistory, socialhistory.NumberOfYears, socialhistory.CurrentSmokingHistory,
			socialhistory.CigarettesPerDay, socialhistory.AlcoholHistory, socialhistory.HowRegular)
		if err != nil {
			return -1, err
		}
	}
	if vitalstatistics != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO vitalstatistics (id, temperature, spo2, systolic_bp1, 
diastolic_bp1, systolic_bp2, diastolic_bp2, avg_systolic_bp, avg_diastolic_bp, hr1, hr2, avg_hr, 
rand_blood_glucose_mmolL, rand_blood_glucose_mmolLp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
			patientid, vitalstatistics.Temperature, vitalstatistics.SpO2, vitalstatistics.SystolicBP1,
			vitalstatistics.DiastolicBP1, vitalstatistics.SystolicBP2, vitalstatistics.DiastolicBP2,
			vitalstatistics.AverageSystolicBP, vitalstatistics.AverageDiastolicBP, vitalstatistics.HR1,
			vitalstatistics.HR2, vitalstatistics.AverageHR, vitalstatistics.RandomBloodGlucoseMmolL,
			vitalstatistics.RandomBloodGlucoseMmolLp)
		if err != nil {
			return -1, err
		}
	}
	if heightandweight != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO heightandweight (id, height, weight, bmi, bmi_analysis, 
		paeds_height, paeds_weight) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			patientid, heightandweight.Height, heightandweight.Weight, heightandweight.BMI,
			heightandweight.BMIAnalysis, heightandweight.PaedsHeight, heightandweight.PaedsWeight)
		if err != nil {
			return -1, err
		}
	}
	if visualacuity != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO visualacuity (id, l_eye_vision, r_eye_vision, 
		additional_intervention) VALUES ($1, $2, $3, $4)`,
			patientid, visualacuity.LEyeVision, visualacuity.REyeVision, visualacuity.AdditionalIntervention)
		if err != nil {
			return -1, err
		}
	}
	if doctorsconsultation != nil {
		_, err = tx.ExecContext(ctx, `INSERT INTO doctorsconsultation (id, healthy, msk, cvs, respi, gu, git, eye, 
		derm, others, consultation_notes, diagnosis, treatment, referral_needed, referral_loc, remarks) VALUES ($1, $2, 
		$3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`,
			patientid, doctorsconsultation.Healthy, doctorsconsultation.Msk, doctorsconsultation.Cvs,
			doctorsconsultation.Respi, doctorsconsultation.Gu, doctorsconsultation.Git,
			doctorsconsultation.Eye, doctorsconsultation.Derm, doctorsconsultation.Others,
			doctorsconsultation.ConsultationNotes, doctorsconsultation.Diagnosis, doctorsconsultation.Treatment,
			doctorsconsultation.ReferralNeeded, doctorsconsultation.ReferralLoc, doctorsconsultation.Remarks)
		if err != nil {
			return -1, err
		}
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}
	return patientid, nil
}

func (p *postgresPatientRepository) DeletePatientByID(ctx context.Context, id int32) (int32, error) {
	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM pastmedicalhistory WHERE id = $1", id)
	if err != nil {
		return -1, err
	}
	_, err = tx.Exec("DELETE FROM socialhistory WHERE id = $1", id)
	if err != nil {
		return -1, err
	}
	_, err = tx.Exec("DELETE FROM vitalstatistics WHERE id = $1", id)
	if err != nil {
		return -1, err
	}
	_, err = tx.Exec("DELETE FROM heightandweight WHERE id = $1", id)
	if err != nil {
		return -1, err
	}
	_, err = tx.Exec("DELETE FROM visualacuity WHERE id = $1", id)
	if err != nil {
		return -1, err
	}
	_, err = tx.Exec("DELETE FROM doctorsconsultation WHERE id = $1", id)
	if err != nil {
		return -1, err
	}
	_, err = tx.Exec("DELETE FROM admin WHERE id = $1", id)
	if err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}
	return id, nil
}

// UpdatePatientByID updates an already existing patient, filling out or overriding any of its fields
func (p *postgresPatientRepository) UpdatePatientByID(ctx context.Context, id int32, patient *entities.Patient) (int32, error) {
	// Checks that a patient exists by searching for admin field
	// Then for each non-nil field in patient, updates it

	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	a := patient.Admin
	pmh := patient.PastMedicalHistory
	socialhistory := patient.SocialHistory
	vs := patient.VitalStatistics
	haw := patient.HeightAndWeight
	va := patient.VisualAcuity
	dc := patient.DoctorsConsultation

	// Check that patient exists already
	rows := tx.QueryRowContext(ctx, "SELECT * FROM admin WHERE id = $1;", id)
	prevAdmin := entities.Admin{}
	err = rows.Scan(
		&prevAdmin.ID,
		&prevAdmin.FamilyGroup,
		&prevAdmin.RegDate,
		&prevAdmin.Name,
		&prevAdmin.KhmerName,
		&prevAdmin.Dob,
		&prevAdmin.Age,
		&prevAdmin.Gender,
		&prevAdmin.Village,
		&prevAdmin.ContactNo,
		&prevAdmin.Pregnant,
		&prevAdmin.LastMenstrualPeriod,
		&prevAdmin.DrugAllergies,
		&prevAdmin.SentToID,
		&prevAdmin.Photo,
	)

	if err != nil { // no patient found
		return -1, entities.ErrPatientNotFound
	}
	if a != nil { // Update admin
		_, err = tx.ExecContext(ctx, `UPDATE admin SET family_group = $1, reg_date = $2, name = $3, khmer_name = $4, dob = $5, age = $6, 
		gender = $7, village = $8, contact_no = $9, pregnant = $10, last_menstrual_period = $11, drug_allergies = $12,
		sent_to_id = $13, photo = $14 WHERE id = $15`, a.FamilyGroup, a.RegDate, a.Name, a.KhmerName, a.Dob, a.Age, a.Gender, a.Village, a.ContactNo,
			a.Pregnant, a.LastMenstrualPeriod, a.DrugAllergies, a.SentToID, a.Photo, id)
		if err != nil {
			return -1, err
		}
	}
	if pmh != nil { // Update pastmedicalhistory
		_, err = tx.ExecContext(ctx, `
		INSERT INTO pastmedicalhistory (id, tuberculosis, diabetes, hypertension, hyperlipidemia, chronic_joint_pains,
										 chronic_muscle_aches, sexually_transmitted_disease, specified_stds, others) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		ON CONFLICT (id) DO UPDATE SET
			tuberculosis = $2, 
			diabetes = $3,
			hypertension = $4,
			hyperlipidemia = $5,
			chronic_joint_pains = $6,
			chronic_muscle_aches = $7,
			sexually_transmitted_disease = $8,
			specified_stds = $9,
			others = $10
		`, id, pmh.Tuberculosis, pmh.Diabetes, pmh.Hypertension, pmh.Hyperlipidemia,
			pmh.ChronicJointPains, pmh.ChronicMuscleAches, pmh.SexuallyTransmittedDisease,
			pmh.SpecifiedSTDs, pmh.Others)
		if err != nil {
			return -1, err
		}
	}
	if socialhistory != nil {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO socialhistory (id, past_smoking_history, no_of_years, current_smoking_history, cigarettes_per_day, 
		alcohol_history, how_regular) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		ON CONFLICT (id) DO UPDATE SET
			past_smoking_history = $2,
			no_of_years = $3,
			current_smoking_history = $4,
			cigarettes_per_day = $5,
			alcohol_history = $6,
			how_regular = $7
		`, id, socialhistory.PastSmokingHistory, socialhistory.NumberOfYears, socialhistory.CurrentSmokingHistory,
			socialhistory.CigarettesPerDay, socialhistory.AlcoholHistory, socialhistory.HowRegular)

		if err != nil {
			return -1, err
		}
	}
	if vs != nil {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO vitalstatistics (id, temperature, spO2, systolic_bp1, diastolic_bp1, systolic_bp2, diastolic_bp2, 
		avg_systolic_bp, avg_diastolic_bp, hr1, hr2, avg_hr, rand_blood_glucose_mmolL, rand_blood_glucose_mmolLp) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) 
		ON CONFLICT (id) DO UPDATE SET
			temperature = $2,
			spO2 = $3,
			systolic_bp1 = $4,
			diastolic_bp1 = $5,
			systolic_bp2 = $6,
			diastolic_bp2 = $7,
			avg_systolic_bp = $8,
			avg_diastolic_bp = $9,
			hr1 = $10,
			hr2 = $11,
			avg_hr = $12,
			rand_blood_glucose_mmolL = $13,
			rand_blood_glucose_mmolLp = $14
		`, id, vs.Temperature, vs.SpO2, vs.SystolicBP1, vs.DiastolicBP1, vs.SystolicBP2, vs.DiastolicBP2,
			vs.AverageSystolicBP, vs.AverageDiastolicBP, vs.HR1, vs.HR2, vs.AverageHR, vs.RandomBloodGlucoseMmolL,
			vs.RandomBloodGlucoseMmolLp)

		if err != nil {
			return -1, err
		}
	}
	if haw != nil {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO heightandweight (id, height, weight, bmi, bmi_analysis, paeds_height, paeds_weight) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		ON CONFLICT (id) DO UPDATE SET
			height = $2,
			weight = $3,
			bmi = $4,
			bmi_analysis = $5,
			paeds_height = $6,
			paeds_weight = $7
		`, id, haw.Height, haw.Weight, haw.BMI, haw.BMIAnalysis, haw.PaedsHeight, haw.PaedsWeight)

		if err != nil {
			return -1, err
		}
	}
	if va != nil {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO visualacuity (id, l_eye_vision, r_eye_vision, additional_intervention) 
		VALUES ($1, $2, $3, $4) 
		ON CONFLICT (id) DO UPDATE SET
			l_eye_vision = $2,
			r_eye_vision = $3,
			additional_intervention = $4
		`, id, va.LEyeVision, va.REyeVision, va.AdditionalIntervention)

		if err != nil {
			return -1, err
		}
	}
	if dc != nil {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO doctorsconsultation (id, healthy, msk, cvs, respi, gu, git, eye, derm, others, 
		consultation_notes, diagnosis, treatment, referral_needed, referral_loc, remarks) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) 
		ON CONFLICT(id) 
		DO UPDATE SET
			healthy = $2,
			msk = $3,
			cvs = $4,
			respi = $5,
			gu = $6,
			git = $7,
			eye = $8,
			derm = $9,
			others = $10,
			consultation_notes = $11,
			diagnosis = $12,
			treatment = $13,
			referral_needed = $14,
			referral_loc = $15,
			remarks = $16
		`,
			id, dc.Healthy, dc.Msk, dc.Cvs, dc.Respi, dc.Gu, dc.Git, dc.Eye, dc.Derm, dc.Others, dc.ConsultationNotes,
			dc.Diagnosis, dc.Treatment, dc.ReferralNeeded, dc.ReferralLoc, dc.Remarks)

		if err != nil {
			return -1, err
		}
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}
	return id, nil
}

func (p *postgresPatientRepository) GetAllAdmin(ctx context.Context) ([]entities.PartAdmin, error) {
	var rows *sql.Rows
	result := make([]entities.PartAdmin, 0)
	query := "SELECT id, name, khmer_name, dob, gender, contact_no FROM ADMIN"
	rows, err := p.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		partadmin := entities.PartAdmin{}
		err = rows.Scan(&partadmin.ID, &partadmin.Name, &partadmin.KhmerName, &partadmin.Dob, &partadmin.Gender, &partadmin.ContactNo)

		if err != nil {
			return nil, err
		}
		result = append(result, partadmin)
	}

	return result, nil
}

func (p *postgresPatientRepository) SearchPatients(ctx context.Context, search string) ([]entities.PartAdmin, error) {
	var rows *sql.Rows
	result := make([]entities.PartAdmin, 0)

	rows, err := p.Conn.QueryContext(ctx, `
				SELECT id, name, khmer_name, dob, gender, contact_no 
				FROM ADMIN 
				WHERE LOWER(name) = LOWER($1) OR LOWER(khmer_name) = LOWER($2)`, search, search)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		partadmin := entities.PartAdmin{}
		err = rows.Scan(&partadmin.ID, &partadmin.Name, &partadmin.KhmerName, &partadmin.Dob, &partadmin.Gender, &partadmin.ContactNo)

		if err != nil {
			return nil, err
		}
		result = append(result, partadmin)
	}

	return result, nil
}
