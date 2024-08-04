package postgres

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"github.com/jieqiboh/sothea_backend/entities"
	"github.com/jieqiboh/sothea_backend/util"
	"github.com/joho/sqltocsv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type postgresPatientRepository struct {
	Conn *sql.DB
}

// NewPostgresPatientRepository will create an object that represent the patient.Repository interface
func NewPostgresPatientRepository(conn *sql.DB) entities.PatientRepository {
	return &postgresPatientRepository{conn}
}

// GetPatientVisit returns a Patient struct representing a single visit based on ID, and Entry ID. Only guaranteed field is Admin
func (p *postgresPatientRepository) GetPatientVisit(ctx context.Context, id int32, vid int32) (res *entities.Patient, err error) {
	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	rows := tx.QueryRowContext(ctx, "SELECT * FROM admin WHERE id = $1 AND vid = $2;", id, vid)
	admin := entities.Admin{}
	var photoBytes []byte
	err = rows.Scan(
		&admin.ID,
		&admin.VID,
		&admin.FamilyGroup,
		&admin.RegDate,
		&admin.QueueNo,
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
		&photoBytes,
	)
	photoBase64 := base64.StdEncoding.EncodeToString(photoBytes)
	admin.Photo = &photoBase64
	if err != nil { // no admin found
		return nil, entities.ErrPatientVisitNotFound
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM pastmedicalhistory WHERE pastmedicalhistory.id = $1 AND pastmedicalhistory.vid = $2;", id, vid)
	pastmedicalhistory := &entities.PastMedicalHistory{}
	err = rows.Scan(
		&pastmedicalhistory.ID,
		&pastmedicalhistory.VID,
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

	rows = tx.QueryRowContext(ctx, "SELECT * FROM socialhistory WHERE socialhistory.id = $1 AND socialhistory.vid = $2;", id, vid)
	socialhistory := &entities.SocialHistory{}
	err = rows.Scan(
		&socialhistory.ID,
		&socialhistory.VID,
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

	rows = tx.QueryRowContext(ctx, "SELECT * FROM vitalstatistics WHERE vitalstatistics.id = $1 AND vitalstatistics.vid = $2;", id, vid)
	vitalstatistics := &entities.VitalStatistics{}
	err = rows.Scan(
		&vitalstatistics.ID,
		&vitalstatistics.VID,
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

	rows = tx.QueryRowContext(ctx, "SELECT * FROM heightandweight WHERE heightandweight.id = $1 AND heightandweight.vid = $2;", id, vid)
	heightandweight := &entities.HeightAndWeight{}
	err = rows.Scan(
		&heightandweight.ID,
		&heightandweight.VID,
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

	rows = tx.QueryRowContext(ctx, "SELECT * FROM visualacuity WHERE visualacuity.id = $1 AND visualacuity.vid = $2;", id, vid)
	visualacuity := &entities.VisualAcuity{}
	err = rows.Scan(
		&visualacuity.ID,
		&visualacuity.VID,
		&visualacuity.LEyeVision,
		&visualacuity.REyeVision,
		&visualacuity.AdditionalIntervention,
	)
	if errors.Is(sql.ErrNoRows, err) { // no visualacuity found
		visualacuity = nil
	} else if err != nil { // unknown error
		return nil, err
	}

	rows = tx.QueryRowContext(ctx, "SELECT * FROM doctorsconsultation WHERE doctorsconsultation.id = $1 AND doctorsconsultation.vid = $2;", id, vid)
	doctorsconsultation := &entities.DoctorsConsultation{}
	err = rows.Scan(
		&doctorsconsultation.ID,
		&doctorsconsultation.VID,
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

// CreatePatient inserts a new Admin category for a new patient and returns the new id if successful.
func (p *postgresPatientRepository) CreatePatient(ctx context.Context, admin *entities.Admin) (int32, error) {
	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil { // error starting transaction
		return -1, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var patientid int32
	if admin == nil { // no admin field
		return -1, entities.ErrMissingAdminCategory
	}
	rows := tx.QueryRowContext(ctx, `INSERT INTO admin (family_group, reg_date, queue_no, name, khmer_name, dob, age, gender, village, 
	contact_no, pregnant, last_menstrual_period, drug_allergies, sent_to_id, photo) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id`,
		admin.FamilyGroup, admin.RegDate, admin.QueueNo, admin.Name, admin.KhmerName, admin.Dob, admin.Age, admin.Gender, admin.Village, admin.ContactNo,
		admin.Pregnant, admin.LastMenstrualPeriod, admin.DrugAllergies, admin.SentToID, admin.Photo)
	err = rows.Scan(&patientid)
	if err != nil { // error inserting admin
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}
	return patientid, nil
}

// CreatePatientVisit inserts a new Admin category for an existing patient and returns the new vid if successful. Only required field is Admin
func (p *postgresPatientRepository) CreatePatientVisit(ctx context.Context, id int32, admin *entities.Admin) (int32, error) {
	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil { // error starting transaction
		return -1, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var patientid int32
	if admin == nil { // no admin field
		return -1, entities.ErrMissingAdminCategory
	}

	// Check that patient exists
	doesPatientExist, err := p.checkPatientExists(ctx, id)
	if err != nil { // query error
		return -1, err
	} else if doesPatientExist == false { // no query error, and patient doesn't exist
		return -1, entities.ErrPatientNotFound
	}

	rows := tx.QueryRowContext(ctx, `INSERT INTO admin (id, family_group, reg_date, queue_no, name, khmer_name, dob, age, gender, village, 
	contact_no, pregnant, last_menstrual_period, drug_allergies, sent_to_id, photo) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING vid`,
		id, admin.FamilyGroup, admin.RegDate, admin.QueueNo, admin.Name, admin.KhmerName, admin.Dob, admin.Age, admin.Gender, admin.Village, admin.ContactNo,
		admin.Pregnant, admin.LastMenstrualPeriod, admin.DrugAllergies, admin.SentToID, admin.Photo)
	err = rows.Scan(&patientid)
	if err != nil { // error inserting admin
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}
	return patientid, nil
}

// Deletes all patient entries where id and vid match
func (p *postgresPatientRepository) DeletePatientVisit(ctx context.Context, id int32, vid int32) error {
	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM pastmedicalhistory WHERE pastmedicalhistory.id = $1 AND pastmedicalhistory.vid = $2;", id, vid)
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM socialhistory WHERE socialhistory.id = $1 AND socialhistory.vid = $2;", id, vid)
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM vitalstatistics WHERE vitalstatistics.id = $1 AND vitalstatistics.vid = $2;", id, vid)
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM heightandweight WHERE heightandweight.id = $1 AND heightandweight.vid = $2;", id, vid)
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM visualacuity WHERE visualacuity.id = $1 AND visualacuity.vid = $2;", id, vid)
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM doctorsconsultation WHERE doctorsconsultation.id = $1 AND doctorsconsultation.vid = $2;", id, vid)
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM admin WHERE id = $1 AND vid = $2", id, vid)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

// UpdatePatientVisit updates a visit for an existing patient, filling out or overriding any of its fields
func (p *postgresPatientRepository) UpdatePatientVisit(ctx context.Context, id int32, vid int32, patient *entities.Patient) error {
	// Checks that a patient exists by searching for admin field
	// Then for each non-nil field in patient, updates it
	// Start a new transaction
	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Check that patient visit exists
	doesPatientVisitExist, err := p.checkPatientVisitExists(ctx, id, vid)
	if err != nil {
		return err
	} else if doesPatientVisitExist == false {
		return entities.ErrPatientVisitNotFound
	}

	a := patient.Admin
	pmh := patient.PastMedicalHistory
	socialhistory := patient.SocialHistory
	vs := patient.VitalStatistics
	haw := patient.HeightAndWeight
	va := patient.VisualAcuity
	dc := patient.DoctorsConsultation
	if a != nil { // Update admin
		_, err = tx.ExecContext(ctx, `UPDATE admin SET family_group = $1, reg_date = $2, queue_no = $3, name = $4, khmer_name = $5, dob = $6, age = $7, 
		gender = $8, village = $9, contact_no = $10, pregnant = $11, last_menstrual_period = $12, drug_allergies = $13,
		sent_to_id = $14, photo = $15 WHERE id = $16 AND vid = $17`, a.FamilyGroup, a.RegDate, a.QueueNo, a.Name, a.KhmerName, a.Dob, a.Age, a.Gender, a.Village, a.ContactNo,
			a.Pregnant, a.LastMenstrualPeriod, a.DrugAllergies, a.SentToID, a.Photo, id, vid)
		if err != nil {
			return err
		}
	}
	if pmh != nil { // Update pastmedicalhistory, use insert into on conflict update because not it isn't guaranteed to exist
		_, err = tx.ExecContext(ctx, `
		INSERT INTO pastmedicalhistory (id, vid, tuberculosis, diabetes, hypertension, hyperlipidemia, chronic_joint_pains,
										 chronic_muscle_aches, sexually_transmitted_disease, specified_stds, others) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
		ON CONFLICT (id, vid) DO UPDATE SET
			tuberculosis = $3, 
			diabetes = $4,
			hypertension = $5,
			hyperlipidemia = $6,
			chronic_joint_pains = $7,
			chronic_muscle_aches = $8,
			sexually_transmitted_disease = $9,
			specified_stds = $10,
			others = $11
		`, id, vid, pmh.Tuberculosis, pmh.Diabetes, pmh.Hypertension, pmh.Hyperlipidemia,
			pmh.ChronicJointPains, pmh.ChronicMuscleAches, pmh.SexuallyTransmittedDisease,
			pmh.SpecifiedSTDs, pmh.Others)
		if err != nil {
			return err
		}
	}
	if socialhistory != nil {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO socialhistory (id, vid, past_smoking_history, no_of_years, current_smoking_history, cigarettes_per_day, 
		alcohol_history, how_regular) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
		ON CONFLICT (id, vid) DO UPDATE SET
			past_smoking_history = $3,
			no_of_years = $4,
			current_smoking_history = $5,
			cigarettes_per_day = $6,
			alcohol_history = $7,
			how_regular = $8
		`, id, vid, socialhistory.PastSmokingHistory, socialhistory.NumberOfYears, socialhistory.CurrentSmokingHistory,
			socialhistory.CigarettesPerDay, socialhistory.AlcoholHistory, socialhistory.HowRegular)

		if err != nil {
			return err
		}
	}
	if vs != nil {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO vitalstatistics (id, vid, temperature, spO2, systolic_bp1, diastolic_bp1, systolic_bp2, diastolic_bp2, 
		avg_systolic_bp, avg_diastolic_bp, hr1, hr2, avg_hr, rand_blood_glucose_mmolL, rand_blood_glucose_mmolLp) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) 
		ON CONFLICT (id, vid) DO UPDATE SET
			temperature = $3,
			spO2 = $4,
			systolic_bp1 = $4,
			diastolic_bp1 = $6,
			systolic_bp2 = $7,
			diastolic_bp2 = $8,
			avg_systolic_bp = $9,
			avg_diastolic_bp = $10,
			hr1 = $11,
			hr2 = $12,
			avg_hr = $13,
			rand_blood_glucose_mmolL = $14,
			rand_blood_glucose_mmolLp = $15
		`, id, vid, vs.Temperature, vs.SpO2, vs.SystolicBP1, vs.DiastolicBP1, vs.SystolicBP2, vs.DiastolicBP2,
			vs.AverageSystolicBP, vs.AverageDiastolicBP, vs.HR1, vs.HR2, vs.AverageHR, vs.RandomBloodGlucoseMmolL,
			vs.RandomBloodGlucoseMmolLp)

		if err != nil {
			return err
		}
	}
	if haw != nil {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO heightandweight (id, vid, height, weight, bmi, bmi_analysis, paeds_height, paeds_weight) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
		ON CONFLICT (id, vid) DO UPDATE SET
			height = $3,
			weight = $4,
			bmi = $5,
			bmi_analysis = $6,
			paeds_height = $7,
			paeds_weight = $8
		`, id, vid, haw.Height, haw.Weight, haw.BMI, haw.BMIAnalysis, haw.PaedsHeight, haw.PaedsWeight)

		if err != nil {
			return err
		}
	}
	if va != nil {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO visualacuity (id, vid, l_eye_vision, r_eye_vision, additional_intervention) 
		VALUES ($1, $2, $3, $4, $5) 
		ON CONFLICT (id, vid) DO UPDATE SET
			l_eye_vision = $3,
			r_eye_vision = $4,
			additional_intervention = $5
		`, id, vid, va.LEyeVision, va.REyeVision, va.AdditionalIntervention)

		if err != nil {
			return err
		}
	}
	if dc != nil {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO doctorsconsultation (id, vid, healthy, msk, cvs, respi, gu, git, eye, derm, others, 
		consultation_notes, diagnosis, treatment, referral_needed, referral_loc, remarks) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) 
		ON CONFLICT(id, vid) 
		DO UPDATE SET
			healthy = $3,
			msk = $4,
			cvs = $5,
			respi = $6,
			gu = $7,
			git = $8,
			eye = $9,
			derm = $10,
			others = $11,
			consultation_notes = $12,
			diagnosis = $13,
			treatment = $14,
			referral_needed = $15,
			referral_loc = $16,
			remarks = $17
		`,
			id, vid, dc.Healthy, dc.Msk, dc.Cvs, dc.Respi, dc.Gu, dc.Git, dc.Eye, dc.Derm, dc.Others, dc.ConsultationNotes,
			dc.Diagnosis, dc.Treatment, dc.ReferralNeeded, dc.ReferralLoc, dc.Remarks)

		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *postgresPatientRepository) GetAllAdmin(ctx context.Context) ([]entities.PartAdmin, error) {
	var rows *sql.Rows
	result := make([]entities.PartAdmin, 0)
	query := "SELECT id, queue_no, name, khmer_name, dob, gender, contact_no FROM ADMIN"
	rows, err := p.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		partadmin := entities.PartAdmin{}
		err = rows.Scan(&partadmin.ID, &partadmin.QueueNo, &partadmin.Name, &partadmin.KhmerName, &partadmin.Dob, &partadmin.Gender, &partadmin.ContactNo)

		if err != nil {
			return nil, err
		}
		result = append(result, partadmin)
	}

	return result, nil
}

func (p *postgresPatientRepository) GetPatientMeta(ctx context.Context, id int32) (*entities.Patient, error) {
	//
	//TODO implement me
	panic("implement me")
}

func (p *postgresPatientRepository) GetPatientVisitMeta(ctx context.Context, id int32, vid int32) (*entities.Patient, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postgresPatientRepository) GetAllPatientVisitMeta(ctx context.Context) ([]entities.PartAdmin, error) {
	//TODO implement me
	panic("implement me")
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

func (p *postgresPatientRepository) ExportDatabaseToCSV(ctx context.Context) error {
	query := `SELECT
        a.id,
        a.family_group,
        a.reg_date,
        a.queue_no,
        a.name,
        a.khmer_name,
        a.dob,
        a.age,
        a.gender,
        a.village,
        a.contact_no,
        a.pregnant,
        a.last_menstrual_period,
        a.drug_allergies,
        a.sent_to_id,
        -- Past Medical History
        p.tuberculosis,
        p.diabetes,
        p.hypertension,
        p.hyperlipidemia,
        p.chronic_joint_pains,
        p.chronic_muscle_aches,
        p.sexually_transmitted_disease,
        p.specified_stds,
        p.others AS pmh_others,
        -- Social History
        s.past_smoking_history,
        s.no_of_years,
        s.current_smoking_history,
        s.cigarettes_per_day,
        s.alcohol_history,
        s.how_regular,
        -- Vital Statistics
        v.temperature,
        v.spo2,
        v.systolic_bp1,
        v.diastolic_bp1,
        v.systolic_bp2,
        v.diastolic_bp2,
        v.avg_systolic_bp,
        v.avg_diastolic_bp,
        v.hr1,
        v.hr2,
        v.avg_hr,
        v.rand_blood_glucose_mmolL,
        v.rand_blood_glucose_mmolLp,
        -- Height and Weight
        h.height,
        h.weight,
        h.bmi,
        h.bmi_analysis,
        h.paeds_height,
        h.paeds_weight,
        -- Visual Acuity
        va.l_eye_vision,
        va.r_eye_vision,
        va.additional_intervention,
        -- Doctors Consultation
        d.healthy,
        d.msk,
        d.cvs,
        d.respi,
        d.gu,
        d.git,
        d.eye,
        d.derm,
        d.others AS dc_others,
        d.consultation_notes,
        d.diagnosis,
        d.treatment,
        d.referral_needed,
        d.referral_loc,
        d.remarks
    FROM
        admin a
    LEFT JOIN
        pastmedicalhistory p ON a.id = p.id
    LEFT JOIN
        socialhistory s ON a.id = s.id
    LEFT JOIN
        vitalstatistics v ON a.id = v.id
    LEFT JOIN
        heightandweight h ON a.id = h.id
    LEFT JOIN
        visualacuity va ON a.id = va.id
    LEFT JOIN
        doctorsconsultation d ON a.id = d.id`
	// Execute the query
	rows, err := p.Conn.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	filePath := util.MustGitPath("repository/tmp/output.csv")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	err = sqltocsv.WriteFile(filePath, rows)
	if err != nil {
		panic(err)
	}

	return nil
}

func (p *postgresPatientRepository) checkPatientExists(ctx context.Context, id int32) (bool, error) {
	prevAdmin := entities.Admin{}
	err := p.Conn.QueryRowContext(ctx, "SELECT * FROM admin WHERE id = $1;", id).Scan(
		&prevAdmin.ID,
		&prevAdmin.VID,
		&prevAdmin.FamilyGroup,
		&prevAdmin.RegDate,
		&prevAdmin.QueueNo,
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
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		log.Fatalf("query error: %v\n", err)
		return false, err
	}

	return true, nil
}

func (p *postgresPatientRepository) checkPatientVisitExists(ctx context.Context, id int32, vid int32) (bool, error) {
	prevAdmin := entities.Admin{}
	err := p.Conn.QueryRowContext(ctx, "SELECT * FROM admin WHERE id = $1 AND vid = $2;", id, vid).Scan(
		&prevAdmin.ID,
		&prevAdmin.VID,
		&prevAdmin.FamilyGroup,
		&prevAdmin.RegDate,
		&prevAdmin.QueueNo,
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
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		log.Fatalf("query error: %v\n", err)
		return false, err
	}

	return true, nil
}
