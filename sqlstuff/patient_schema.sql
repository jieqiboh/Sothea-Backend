/******************* 
Create the schema and Load Extensions
********************/ 

CREATE TABLE IF NOT EXISTS admin (
    ID                    SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    family_group          TEXT NOT NULL,
    reg_date              DATE NOT NULL,
    name                  TEXT NOT NULL,
    dob                   DATE NOT NULL,
    age                   INTEGER NOT NULL,
    gender                VARCHAR(1) NOT NULL,
    village               TEXT NOT NULL,
    contact_no            TEXT NOT NULL,
    pregnant              BOOLEAN NOT NULL,
    last_menstrual_period Date,
    drug_allergies        TEXT,
    sent_to_id            BOOLEAN NOT NULL
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
    others             BOOLEAN NOT NULL,
    consultation_notes TEXT,
    diagnosis          TEXT,
    treatment          TEXT,
    referral_needed    BOOLEAN NOT NULL,
    referral_loc       TEXT,
    remarks            TEXT
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);
