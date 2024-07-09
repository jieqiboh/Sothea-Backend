/*******************
    Drop the tables
********************/
DROP TABLE IF EXISTS pastmedicalhistory;
DROP TABLE IF EXISTS socialhistory;
DROP TABLE IF EXISTS vitalstatistics;
DROP TABLE IF EXISTS heightandweight;
DROP TABLE IF EXISTS visualacuity;
DROP TABLE IF EXISTS doctorsconsultation;
DROP TABLE IF EXISTS admin;

/*******************
Create the schema and Load Extensions
********************/

CREATE TABLE IF NOT EXISTS admin (
    ID                    SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    family_group          TEXT NOT NULL,
    reg_date              DATE NOT NULL,
    name                  TEXT NOT NULL,
    khmer_name            TEXT NOT NULL,
    dob                   DATE NOT NULL,
    age                   INTEGER NOT NULL,
    gender                VARCHAR(1) NOT NULL,
    village               TEXT NOT NULL,
    contact_no            TEXT NOT NULL,
    pregnant              BOOLEAN NOT NULL,
    last_menstrual_period Date,
    drug_allergies        TEXT,
    sent_to_id            BOOLEAN NOT NULL,
    photo                 BYTEA
);

CREATE TABLE IF NOT EXISTS pastmedicalhistory (
    id                           SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    tuberculosis                 BOOLEAN NOT NULL,
    diabetes                     BOOLEAN NOT NULL,
    hypertension                 BOOLEAN NOT NULL,
    hyperlipidemia               BOOLEAN NOT NULL,
    chronic_joint_pains          BOOLEAN NOT NULL,
    chronic_muscle_aches         BOOLEAN NOT NULL,
    sexually_transmitted_disease BOOLEAN NOT NULL,
    specified_stds               TEXT,
    others                       TEXT,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);

CREATE TABLE IF NOT EXISTS socialhistory (
    id                      SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    past_smoking_history    BOOLEAN NOT NULL,
    no_of_years             INTEGER,
    current_smoking_history BOOLEAN NOT NULL,
    cigarettes_per_day      INTEGER,
    alcohol_history BOOLEAN NOT NULL,
    how_regular             VARCHAR(1),
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);

CREATE TABLE IF NOT EXISTS vitalstatistics (
    id                        SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    temperature               NUMERIC(5, 1) NOT NULL,
    spo2                      NUMERIC(5, 1) NOT NULL,
    systolic_bp1              NUMERIC(5, 1) NOT NULL,
    diastolic_bp1             NUMERIC(5, 1) NOT NULL,
    systolic_bp2              NUMERIC(5, 1) NOT NULL,
    diastolic_bp2             NUMERIC(5, 1) NOT NULL,
    avg_systolic_bp           NUMERIC(5, 1) NOT NULL,
    avg_diastolic_bp          NUMERIC(5, 1) NOT NULL,
    hr1                       NUMERIC(5, 1) NOT NULL,
    hr2                       NUMERIC(5, 1) NOT NULL,
    avg_hr                    NUMERIC(5, 1) NOT NULL,
    rand_blood_glucose_mmolL  NUMERIC(5, 1) NOT NULL,
    rand_blood_glucose_mmolLp NUMERIC(5, 1) NOT NULL,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);

CREATE TABLE IF NOT EXISTS heightandweight (
    id            SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    height        NUMERIC(5, 1) NOT NULL,
    weight        NUMERIC(5, 1) NOT NULL,
    bmi           NUMERIC(5, 1) NOT NULL,
    bmi_analysis  TEXT NOT NULL,
    paeds_height NUMERIC(5, 1) NOT NULL,
    paeds_weight  NUMERIC(5, 1) NOT NULL,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);

CREATE TABLE IF NOT EXISTS visualacuity (
    id                      SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    l_eye_vision            INTEGER NOT NULL,
    r_eye_vision            INTEGER NOT NULL,
    additional_intervention TEXT,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);

CREATE TABLE IF NOT EXISTS doctorsconsultation (
    id                 SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    healthy            BOOLEAN NOT NULL,
    msk                BOOLEAN NOT NULL,
    cvs                BOOLEAN NOT NULL,
    respi              BOOLEAN NOT NULL,
    gu                 BOOLEAN NOT NULL,
    git                BOOLEAN NOT NULL,
    eye                BOOLEAN NOT NULL,
    derm               BOOLEAN NOT NULL,
    others             TEXT NOT NULL,
    consultation_notes TEXT,
    diagnosis          TEXT,
    treatment          TEXT,
    referral_needed    BOOLEAN NOT NULL,
    referral_loc       TEXT,
    remarks            TEXT,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);


/*******************
    Load the data
 */

INSERT INTO admin (family_group, reg_date, name, khmer_name, dob, age, gender, village, contact_no, pregnant,
                   last_menstrual_period, drug_allergies, sent_to_id) VALUES
('S001', '2024-01-10', 'John Doe', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1994-01-10', 30, 'M', 'SO', '12345678', FALSE, NULL, 'panadol', FALSE),
('S002A', '2024-01-10', 'Jane Smith', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1999-01-10', 25, 'F', 'SO', '12345679', FALSE, NULL, NULL, FALSE),
('S002B', '2024-01-10', 'Bob Smith', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1999-01-10', 25, 'M', 'R1', '99999999', FALSE, NULL, 'aspirin', FALSE),
('S003', '2024-01-10', 'Bob Johnson', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1989-01-10', 35, 'M', 'R1', '11111111', FALSE, NULL, NULL, FALSE),
('S004', '2024-01-10', 'Alice Brown', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1996-01-10', 28, 'F', 'R1', '17283948', FALSE, NULL, NULL, FALSE),
('S005A', '2024-01-10', 'Charlie Davis', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1982-01-10', 40, 'M', 'R1', '09876543', FALSE, NULL, NULL, FALSE);

INSERT INTO pastmedicalhistory(tuberculosis, diabetes, hypertension, hyperlipidemia, chronic_joint_pains,
                                chronic_muscle_aches, sexually_transmitted_disease, specified_stds, others) VALUES
(TRUE, FALSE, TRUE, FALSE, FALSE, TRUE, TRUE, 'TRICHOMONAS', 'None'),
(FALSE, TRUE, TRUE, TRUE, FALSE, FALSE, FALSE, '', 'CHILDHOOD LEUKAEMIA'),
(TRUE, FALSE, FALSE, FALSE, TRUE, TRUE, FALSE, '', ''),
(FALSE, FALSE, TRUE, FALSE, TRUE, FALSE, TRUE, 'Syphilis', NULL),
(FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, '', '');

INSERT INTO socialhistory (past_smoking_history, no_of_years, current_smoking_history,
                           cigarettes_per_day, alcohol_history, how_regular) VALUES
(TRUE, 15, FALSE, NULL, TRUE, 'A'),
(FALSE, NULL, TRUE, 10, TRUE, 'D'),
(TRUE, 20, TRUE, 5, FALSE, NULL),
(TRUE, 10, FALSE, NULL, TRUE, 'B'),
(FALSE, NULL, FALSE, NULL, FALSE, NULL);

INSERT INTO vitalstatistics (
    temperature, spo2, systolic_bp1, diastolic_bp1, systolic_bp2, diastolic_bp2,
    avg_systolic_bp, avg_diastolic_bp, hr1, hr2, avg_hr, rand_blood_glucose_mmolL,
    rand_blood_glucose_mmolLp
) VALUES
  (36.5, 98, 120, 80, 122, 78, 121, 79, 72, 71, 71.5, 5.4, 5.3),
  (37.0, 97, 130, 85, 128, 82, 129, 83, 68, 70, 69, 5.7, 5.6),
  (36.8, 99, 118, 78, 120, 76, 119, 77, 75, 76, 75.5, 5.6, 5.5),
  (36.7, 98, 125, 82, 124, 80, 124.5, 81, 70, 72, 71, 5.3, 5.2);

INSERT INTO heightandweight (
    height, weight, bmi, bmi_analysis, paeds_height, paeds_weight
) VALUES
(170, 70, 24.2, 'normal weight', 90, 80),
(165, 55, 20.2, 'normal weight', 95, 90),
(180, 85, 26.2, 'overweight', 80, 95);

INSERT INTO visualacuity (
    l_eye_vision, r_eye_vision, additional_intervention
) VALUES
(20, 20, 'VISUAL FIELD TEST REQUIRED'),
(15, 20, 'REFERRED TO BOC');

INSERT INTO doctorsconsultation (
    healthy, msk, cvs, respi, gu, git, eye, derm, others,
    consultation_notes, diagnosis, treatment, referral_needed,
    referral_loc, remarks
) VALUES
(TRUE, FALSE, FALSE, TRUE, TRUE, FALSE, TRUE, FALSE, 'LEUKAEMIA',
'CHEST PAIN, SHORTNESS OF BREATH, COUGH', 'ACUTE BRONCHITIS',
'REST, HYDRATION, COUGH SYRUP', FALSE, NULL, 'MONITOR FOR RESOLUTION');