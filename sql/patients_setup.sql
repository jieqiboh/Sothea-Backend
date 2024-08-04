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

CREATE TABLE IF NOT EXISTS admin
(
    id                    SERIAL, -- Use SERIAL to auto-increment the ID
    vid                   INTEGER    NOT NULL,
    family_group          TEXT       NOT NULL,
    reg_date              DATE       NOT NULL,
    queue_no              TEXT       NOT NULL,
    name                  TEXT       NOT NULL,
    khmer_name            TEXT       NOT NULL,
    dob                   DATE,
    age                   INTEGER,
    gender                VARCHAR(1) NOT NULL,
    village               TEXT       NOT NULL,
    contact_no            TEXT       NOT NULL,
    pregnant              BOOLEAN    NOT NULL,
    last_menstrual_period Date,
    drug_allergies        TEXT,
    sent_to_id            BOOLEAN    NOT NULL,
    photo                 BYTEA,
    PRIMARY KEY (id, vid)         -- Composite primary key
);

CREATE TABLE IF NOT EXISTS pastmedicalhistory
(
    id                           INTEGER NOT NULL,                       -- Use INTEGER to match the id type from admin
    vid                          INTEGER NOT NULL,                       -- Add vid to match the vid type from admin
    tuberculosis                 BOOLEAN NOT NULL,
    diabetes                     BOOLEAN NOT NULL,
    hypertension                 BOOLEAN NOT NULL,
    hyperlipidemia               BOOLEAN NOT NULL,
    chronic_joint_pains          BOOLEAN NOT NULL,
    chronic_muscle_aches         BOOLEAN NOT NULL,
    sexually_transmitted_disease BOOLEAN NOT NULL,
    specified_stds               TEXT,
    others                       TEXT,
    PRIMARY KEY (id, vid),                                               -- Composite primary key
    CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
);

CREATE TABLE IF NOT EXISTS socialhistory
(
    id                      INTEGER NOT NULL,                            -- Use INTEGER to match the id type from admin
    vid                     INTEGER NOT NULL,                            -- Add vid to match the vid type from admin
    past_smoking_history    BOOLEAN NOT NULL,
    no_of_years             INTEGER,
    current_smoking_history BOOLEAN NOT NULL,
    cigarettes_per_day      INTEGER,
    alcohol_history         BOOLEAN NOT NULL,
    how_regular             VARCHAR(1),
    PRIMARY KEY (id, vid),                                               -- Composite primary key
    CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
);

CREATE TABLE IF NOT EXISTS vitalstatistics
(
    id                        INTEGER       NOT NULL,                    -- Use INTEGER to match the id type from admin
    vid                       INTEGER       NOT NULL,                    -- Add vid to match the vid type from admin
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
    PRIMARY KEY (id, vid),                                               -- Composite primary key
    CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
);

CREATE TABLE IF NOT EXISTS heightandweight
(
    id           INTEGER       NOT NULL,                                 -- Use INTEGER to match the id type from admin
    vid          INTEGER       NOT NULL,                                 -- Add vid to match the vid type from admin
    height       NUMERIC(5, 1) NOT NULL,
    weight       NUMERIC(5, 1) NOT NULL,
    bmi          NUMERIC(5, 1) NOT NULL,
    bmi_analysis TEXT          NOT NULL,
    paeds_height NUMERIC(5, 1) NOT NULL,
    paeds_weight NUMERIC(5, 1) NOT NULL,
    PRIMARY KEY (id, vid),                                               -- Composite primary key
    CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
);

CREATE TABLE IF NOT EXISTS visualacuity
(
    id                      INTEGER NOT NULL,                            -- Use INTEGER to match the id type from admin
    vid                     INTEGER NOT NULL,                            -- Add vid to match the vid type from admin
    l_eye_vision            INTEGER NOT NULL,
    r_eye_vision            INTEGER NOT NULL,
    additional_intervention TEXT,
    PRIMARY KEY (id, vid),                                               -- Composite primary key
    CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
);

CREATE TABLE IF NOT EXISTS doctorsconsultation
(
    id                 INTEGER NOT NULL,                                 -- Use INTEGER to match the id type from admin
    vid                INTEGER NOT NULL,                                 -- Add vid to match the vid type from admin
    healthy            BOOLEAN NOT NULL,
    msk                BOOLEAN NOT NULL,
    cvs                BOOLEAN NOT NULL,
    respi              BOOLEAN NOT NULL,
    gu                 BOOLEAN NOT NULL,
    git                BOOLEAN NOT NULL,
    eye                BOOLEAN NOT NULL,
    derm               BOOLEAN NOT NULL,
    others             TEXT    NOT NULL,
    consultation_notes TEXT,
    diagnosis          TEXT,
    treatment          TEXT,
    referral_needed    BOOLEAN NOT NULL,
    referral_loc       TEXT,
    remarks            TEXT,
    PRIMARY KEY (id, vid),                                               -- Composite primary key
    CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
);

/*******************
    Create the trigger function
*******************/

CREATE OR REPLACE FUNCTION set_entry_id() RETURNS TRIGGER AS
$$
DECLARE
    max_entry_id INTEGER;
BEGIN
    -- Check if the ID already exists in the table
    SELECT COALESCE(MAX(VID), 0)
    INTO max_entry_id
    FROM admin
    WHERE ID = NEW.ID;

    -- Increment Entry_ID based on the max_entry_id
    NEW.VID := max_entry_id + 1;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_admin
    BEFORE INSERT
    ON admin
    FOR EACH ROW
EXECUTE FUNCTION set_entry_id();

/*******************
    Create new patients
 */

INSERT INTO admin (family_group, reg_date, queue_no, name, khmer_name, dob, age, gender, village, contact_no, pregnant,
                   last_menstrual_period, drug_allergies, sent_to_id)
VALUES ('S001', '2024-01-10', '1A', 'John Doe', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1994-01-10', 30, 'M', 'SO', '12345678', FALSE, NULL,
        'panadol', FALSE),
       ('S002A', '2024-01-10', '2A', 'Jane Smith', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1999-01-10', 25, 'F', 'SO', '12345679', FALSE,
        NULL, NULL, FALSE),
       ('S002B', '2024-01-10', '2B', 'Bob Smith', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1999-01-10', 25, 'M', 'R1', '99999999', FALSE, NULL,
        'aspirin', FALSE),
       ('S003', '2024-01-10', '3A', 'Bob Johnson', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1989-01-10', 35, 'M', 'R1', '11111111', FALSE,
        NULL, NULL, FALSE),
       ('S004', '2024-01-10', '4B', 'Alice Brown', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1996-01-10', 28, 'F', 'R1', '17283948', FALSE,
        NULL, NULL, FALSE),
       ('S005A', '2024-01-10', '5C', 'Charlie Davis', '១២៣៤ ៥៦៧៨៩០ឥឲ', '1982-01-10', 40, 'M', 'R1', '09876543', FALSE,
        NULL, NULL, FALSE);

INSERT INTO pastmedicalhistory(id, vid, tuberculosis, diabetes, hypertension, hyperlipidemia, chronic_joint_pains,
                               chronic_muscle_aches, sexually_transmitted_disease, specified_stds, others)
VALUES (1, 1, TRUE, FALSE, TRUE, FALSE, FALSE, TRUE, TRUE, 'TRICHOMONAS', 'None'),
       (2, 1, FALSE, TRUE, TRUE, TRUE, FALSE, FALSE, FALSE, '', 'CHILDHOOD LEUKAEMIA'),
       (3, 1, TRUE, FALSE, FALSE, FALSE, TRUE, TRUE, FALSE, '', ''),
       (4, 1, FALSE, FALSE, TRUE, FALSE, TRUE, FALSE, TRUE, 'Syphilis', NULL),
       (5, 1, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE, '', '');

INSERT INTO socialhistory (id, vid, past_smoking_history, no_of_years, current_smoking_history,
                           cigarettes_per_day, alcohol_history, how_regular)
VALUES (1, 1, TRUE, 15, FALSE, NULL, TRUE, 'A'),
       (2, 1, FALSE, NULL, TRUE, 10, TRUE, 'D'),
       (3, 1, TRUE, 20, TRUE, 5, FALSE, NULL),
       (4, 1, TRUE, 10, FALSE, NULL, TRUE, 'B'),
       (5, 1, FALSE, NULL, FALSE, NULL, FALSE, NULL);

INSERT INTO vitalstatistics (id, vid, temperature, spo2, systolic_bp1, diastolic_bp1, systolic_bp2, diastolic_bp2,
                             avg_systolic_bp, avg_diastolic_bp, hr1, hr2, avg_hr, rand_blood_glucose_mmolL,
                             rand_blood_glucose_mmolLp)
VALUES (1, 1, 36.5, 98, 120, 80, 122, 78, 121, 79, 72, 71, 71.5, 5.4, 5.3),
       (2, 1, 37.0, 97, 130, 85, 128, 82, 129, 83, 68, 70, 69, 5.7, 5.6),
       (3, 1, 36.8, 99, 118, 78, 120, 76, 119, 77, 75, 76, 75.5, 5.6, 5.5),
       (4, 1, 36.7, 98, 125, 82, 124, 80, 124.5, 81, 70, 72, 71, 5.3, 5.2);

INSERT INTO heightandweight (id, vid, height, weight, bmi, bmi_analysis, paeds_height, paeds_weight)
VALUES (1, 1, 170, 70, 24.2, 'normal weight', 90, 80),
       (2, 1, 165, 55, 20.2, 'normal weight', 95, 90),
       (3, 1, 180, 85, 26.2, 'overweight', 80, 95);

INSERT INTO visualacuity (id, vid, l_eye_vision, r_eye_vision, additional_intervention)
VALUES (1, 1, 20, 20, 'VISUAL FIELD TEST REQUIRED'),
       (2, 1, 15, 20, 'REFERRED TO BOC');

INSERT INTO doctorsconsultation (id, vid, healthy, msk, cvs, respi, gu, git, eye, derm, others,
                                 consultation_notes, diagnosis, treatment, referral_needed,
                                 referral_loc, remarks)
VALUES (1, 1, TRUE, FALSE, FALSE, TRUE, TRUE, FALSE, TRUE, FALSE, 'LEUKAEMIA',
        'CHEST PAIN, SHORTNESS OF BREATH, COUGH', 'ACUTE BRONCHITIS',
        'REST, HYDRATION, COUGH SYRUP', FALSE, NULL, 'MONITOR FOR RESOLUTION');

/*******************
    Add additional entries for patient 1 and 2
 */
INSERT INTO admin (id, family_group, reg_date, queue_no, name, khmer_name, dob, age, gender, village, contact_no,
                   pregnant, last_menstrual_period, drug_allergies, sent_to_id, photo)
VALUES (1, 'Family 1', '2023-07-01', 'Q123', 'John Doe', 'ខេមរ', '1990-01-01', 33, 'M', 'Village 1', '123456789', false,
        '2023-06-01', 'None', false, NULL);

INSERT INTO admin (id, family_group, reg_date, queue_no, name, khmer_name, dob, age, gender, village, contact_no,
                   pregnant, last_menstrual_period, drug_allergies, sent_to_id, photo)
VALUES (1, 'Family 2', '2023-07-02', 'Q124', 'Jane Doe', 'ចន ឌូ', '1990-01-011', 34, 'F', 'Village 2', '987654321',
        true, '2023-06-15', 'Penicillin', true, NULL);

INSERT INTO admin (id, family_group, reg_date, queue_no, name, khmer_name, dob, age, gender, village, contact_no,
                   pregnant, last_menstrual_period, drug_allergies, sent_to_id, photo)
VALUES (1, 'Family 1', '2023-07-03', 'Q125', 'Alice Doe', 'អាលីស ស្ម៊ីត', '1990-01-010', 35, 'F', 'Village 1',
        '555666777', false, '2023-05-01', 'None', false, NULL);

INSERT INTO admin (id, family_group, reg_date, queue_no, name, khmer_name, dob, age, gender, village, contact_no,
                   pregnant, last_menstrual_period, drug_allergies, sent_to_id, photo)
VALUES (2, 'B009', '2023-07-03', 'Q125', 'Walter White', 'អាលីស ស្ម៊ីត', '1990-01-010', 52, 'M', 'ABQ',
        '555666777', false, '2023-05-01', 'None', false, NULL);

/*******************
    Add remaining categories for second entry for patient 1 and 2
 */

INSERT INTO pastmedicalhistory(id, vid, tuberculosis, diabetes, hypertension, hyperlipidemia, chronic_joint_pains,
                               chronic_muscle_aches, sexually_transmitted_disease, specified_stds, others)
VALUES (1, 2, TRUE, FALSE, TRUE, FALSE, FALSE, TRUE, TRUE, 'TRICHOMONAS', 'None'),
       (2, 2, FALSE, TRUE, TRUE, TRUE, FALSE, FALSE, FALSE, '', 'CHILDHOOD LEUKAEMIA');

INSERT INTO socialhistory (id, vid, past_smoking_history, no_of_years, current_smoking_history,
                           cigarettes_per_day, alcohol_history, how_regular)
VALUES (1, 2, TRUE, 15, FALSE, NULL, TRUE, 'A'),
       (2, 2, FALSE, NULL, TRUE, 10, TRUE, 'D');

INSERT INTO vitalstatistics (id, vid, temperature, spo2, systolic_bp1, diastolic_bp1, systolic_bp2, diastolic_bp2,
                             avg_systolic_bp, avg_diastolic_bp, hr1, hr2, avg_hr, rand_blood_glucose_mmolL,
                             rand_blood_glucose_mmolLp)
VALUES (1, 2, 36.5, 98, 120, 80, 122, 78, 121, 79, 72, 71, 71.5, 5.4, 5.3),
       (2, 2, 37.0, 97, 130, 85, 128, 82, 129, 83, 68, 70, 69, 5.7, 5.6);

INSERT INTO heightandweight (id, vid, height, weight, bmi, bmi_analysis, paeds_height, paeds_weight)
VALUES (1, 2, 170, 70, 24.2, 'normal weight', 90, 80),
       (2, 2, 165, 55, 20.2, 'normal weight', 95, 90);

INSERT INTO visualacuity (id, vid, l_eye_vision, r_eye_vision, additional_intervention)
VALUES (1, 2, 20, 20, 'VISUAL FIELD TEST REQUIRED'),
       (2, 2, 15, 20, 'REFERRED TO BOC');

INSERT INTO doctorsconsultation (id, vid, healthy, msk, cvs, respi, gu, git, eye, derm, others,
                                 consultation_notes, diagnosis, treatment, referral_needed,
                                 referral_loc, remarks)
VALUES (1, 2, TRUE, FALSE, FALSE, TRUE, TRUE, FALSE, TRUE, FALSE, 'LEUKAEMIA',
        'CHEST PAIN, SHORTNESS OF BREATH, COUGH', 'ACUTE BRONCHITIS',
        'REST, HYDRATION, COUGH SYRUP', FALSE, NULL, 'MONITOR FOR RESOLUTION');
INSERT INTO doctorsconsultation (id, vid, healthy, msk, cvs, respi, gu, git, eye, derm, others,
                                 consultation_notes, diagnosis, treatment, referral_needed,
                                 referral_loc, remarks)
VALUES (2, 2, TRUE, FALSE, FALSE, TRUE, TRUE, FALSE, TRUE, FALSE, 'LEUKAEMIA',
        'CHEST PAIN, SHORTNESS OF BREATH, COUGH', 'ACUTE BRONCHITIS',
        'REST, HYDRATION, COUGH SYRUP', FALSE, NULL, 'MONITOR FOR RESOLUTION');