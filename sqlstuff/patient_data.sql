INSERT INTO admin (family_group, reg_date, name, dob, age, gender, village, contact_no, pregnant,
                   last_menstrual_period, drug_allergies, sent_to_id) VALUES
('S001', '2024-01-10', 'John Doe', '1994-01-10', 30, 'M', 'SO', '12345678', FALSE, NULL, 'panadol', FALSE),
('S002A', '2024-01-10', 'Jane Smith', '1999-01-10', 25, 'F', 'SO', '12345679', FALSE, NULL, NULL, FALSE),
('S002B', '2024-01-10', 'Bob Smith', '1999-01-10', 25, 'M', 'R1', '99999999', FALSE, NULL, 'aspirin', FALSE),
('S003', '2024-01-10', 'Bob Johnson', '1989-01-10', 35, 'M', 'R1', '11111111', FALSE, NULL, NULL, FALSE),
('S004', '2024-01-10', 'Alice Brown', '1996-01-10', 28, 'F', 'R1', '17283948', FALSE, NULL, NULL, FALSE),
('S005A', '2024-01-10', 'Charlie Davis', '1982-01-10', 40, 'M', 'R1', '09876543', FALSE, NULL, NULL, FALSE);

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
  (36.7, 98, 125, 82, 124, 80, 124.5, 81, 70, 72, 71, 5.3, 5.2),
--   (36.6, 96, 122, 80, 124, 78, 123, 79, 74, 75, 74.5, 5.2, 5.1);

INSERT INTO heightandweight (
    height, weight, bmi, bmi_analysis, paeds_height, paeds_weight
) VALUES
(170, 70, 24.2, 'normal weight', 90, 80),
(165, 55, 20.2, 'normal weight', 95, 90),
(180, 85, 26.2, 'overweight', 80, 95),
-- (160, 50, 19.5, 'normal weight', 100, 85),
-- (175, 75, 24.5, 'underweight', 85, 92);

INSERT INTO visualacuity (
    l_eye_vision, r_eye_vision, additional_intervention
) VALUES
(20, 20, 'VISUAL FIELD TEST REQUIRED'),
(15, 20, 'REFERRED TO BOC'),
-- (20, 15, NULL),
-- (18, 18, 'RECOMMEND GLASSES'),
-- (20, 20, NULL);

INSERT INTO doctorsconsultation (
    healthy, msk, cvs, respi, gu, git, eye, derm, others,
    consultation_notes, diagnosis, treatment, referral_needed,
    referral_loc, remarks
) VALUES
(TRUE, FALSE, FALSE, TRUE, TRUE, FALSE, TRUE, FALSE, FALSE,
'CHEST PAIN, SHORTNESS OF BREATH, COUGH', 'ACUTE BRONCHITIS',
'REST, HYDRATION, COUGH SYRUP', FALSE, NULL, 'MONITOR FOR RESOLUTION'),
-- (TRUE, TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE,
-- 'KNEE PAIN, SWELLING', 'OSTEOARTHRITIS OF THE KNEE',
-- 'NSAIDs, PHYSICAL THERAPY', FALSE, NULL, 'FOLLOW UP IN 2 WEEKS'),
-- (FALSE, FALSE, TRUE, TRUE, TRUE, FALSE, FALSE, FALSE, FALSE,
-- 'PALPITATIONS, CHEST TIGHTNESS, SHORTNESS OF BREATH', 'ANXIETY-RELATED CARDIAC SYMPTOMS',
-- 'ANXIOLYTICS, BREATHING EXERCISES', TRUE, 'CARDIO CLINIC', 'URGENT REFERRAL FOR CARDIAC ASSESSMENT'),
-- (TRUE, FALSE, TRUE, FALSE, FALSE, FALSE, FALSE, FALSE, FALSE,
-- 'BLOOD IN URINE, PAIN DURING URINATION', 'URINARY TRACT INFECTION',
-- 'ANTIBIOTICS, FLUIDS', FALSE, NULL, NULL);
-- (FALSE, TRUE, FALSE, FALSE, TRUE, FALSE, TRUE, FALSE, FALSE,
-- 'JOINT PAIN, LIMITED MOBILITY', 'OSTEOARTHRITIS OF THE HIP',
-- 'PAIN MANAGEMENT, PHYSIOTHERAPY', TRUE, 'ORTHO CLINIC', 'REFERRAL FOR FURTHER EVALUATION AND MANAGEMENT');