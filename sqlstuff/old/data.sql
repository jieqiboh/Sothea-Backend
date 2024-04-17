/******************* 
Create Patients
********************/ 

-- CALL add_patient(
--     'Sample Family', '2024-01-10', 'John Doe', 30, 'M',
--     true, false, true, false, true,
--     true, 5, false,
--     98.6, 99.0,
--     65.5, 150.0,
--     20, 20,
--     true, 'Regular check-up, no concerns.', false
-- )

-- Insert entries into the admin table
INSERT INTO admin (family_group, reg_date, name, age, gender) VALUES
  ('S001', '2024-01-10', 'John Doe', 30, 'M'),
  ('S002A', '2024-01-10', 'Jane Smith', 25, 'F'),
  ('S002B', '2024-01-10', 'Bob Smith', 25, 'M'),
  ('S003', '2024-01-10', 'Bob Johnson', 35, 'M'),
  ('S004', '2024-01-10', 'Alice Brown', 28, 'F'),
  ('S005A', '2024-01-10', 'Charlie Davis', 40, 'M'),
  ('S006', '2024-01-10', 'Eva White', 27, 'F'),
  ('S007A', '2024-01-10', 'David Wilson', 32, 'M'),
  ('S007B', '2024-01-10', 'Lily Wilson', 28, 'F'),
  ('S008', '2024-01-10', 'Michael Jackson', 45, 'M'),
  ('S009', '2024-01-10', 'Emily Johnson', 33, 'F'),
  ('S010', '2024-01-10', 'Daniel Brown', 29, 'M'),
  ('S011', '2024-01-10', 'Sophia Davis', 42, 'F'),
  ('S012', '2024-01-10', 'Oliver White', 26, 'M'),
  ('S013', '2024-01-10', 'Emma Smith', 38, 'F'),
  ('S014', '2024-01-10', 'Ava Jones', 31, 'F'),
  ('S015', '2024-01-10', 'Benjamin Green', 27, 'M'),
  ('S016', '2024-01-10', 'Mia Thompson', 35, 'F'),
  ('S017', '2024-01-10', 'Elijah Lee', 29, 'M'),
  ('S018', '2024-01-10', 'Amelia Taylor', 37, 'F'),
  ('S019', '2024-01-10', 'Henry Miller', 33, 'M'),
  ('S020', '2024-01-10', 'Sophie Harris', 28, 'F'),
  ('S021', '2024-01-10', 'Liam Wilson', 44, 'M'),
  ('S022A', '2024-01-10', 'Chloe Brown', 26, 'F'),
  ('S022B', '2024-01-10', 'Jackson Smith', 39, 'M'),
  ('S022C', '2024-01-10', 'Avery Davis', 32, 'F'),
  ('S022D', '2024-01-10', 'Ethan Johnson', 28, 'M'),
  ('S023', '2024-01-10', 'Madison White', 36, 'F'),
  ('S024', '2024-01-10', 'Aiden Jackson', 41, 'M'),
  ('S025', '2024-01-10', 'Zoe Harris', 30, 'F');


-- Insert entries into the pastmedicalhistory table
INSERT INTO pastmedicalhistory (tuberculosis, diabetes, hypertension, hyperlipidemia, chronicjointpains)
VALUES
  (true, false, true, false, true),
  (false, true, false, true, false),
  (true, true, false, false, true),
  (false, false, true, true, false),
  (true, false, false, true, false),
  (false, true, true, false, true),
  (true, false, true, true, false),
  (false, true, false, false, true),
  (true, true, true, false, false),
  (false, false, false, true, true),
  (true, false, true, false, true),
  (false, true, false, true, false),
  (true, true, false, false, true),
  (false, false, true, true, false),
  (true, false, false, true, false);

-- Insert entries into the socialhistory table for id range 5 to 12
INSERT INTO socialhistory (id, past_smoking_history, no_of_years, current_smoking_history)
VALUES
  (5, true, 10, false),
  (6, false, null, true),
  (7, true, 5, false),
  (8, false, null, true),
  (9, true, 8, false),
  (10, false, null, true),
  (11, true, 12, false),
  (12, false, null, true);

-- Insert entries into the vitalstatistics table
INSERT INTO vitalstatistics (id, temperature, spO2)
VALUES
  (10, 98.5, 98.0),
  (11, 99.2, 97.5),
  (12, 98.8, 98.3),
  (13, 98.6, 97.8),
  (14, 99.0, 98.2),
  (15, 98.7, 97.6),
  (16, 98.9, 98.1),
  (17, 98.4, 97.7),
  (18, 99.1, 98.4),
  (19, 98.7, 97.9),
  (20, 98.3, 98.2),
  (21, 98.6, 97.5),
  (22, 99.0, 98.3),
  (23, 98.8, 97.8),
  (24, 98.5, 98.1),
  (25, 99.2, 97.6);

-- Insert entries into the heightandweight table
INSERT INTO heightandweight (id, height, weight)
VALUES
  (10, 170.5, 68.5),
  (11, 165.2, 72.3),
  (12, 180.0, 75.0),
  (13, 175.5, 68.8),
  (14, 168.7, 70.2),
  (15, 172.3, 65.6),
  (16, 178.9, 80.1),
  (17, 162.4, 68.7),
  (18, 175.1, 71.5),
  (19, 168.9, 69.2),
  (20, 172.6, 73.4);

-- Insert entries into the visualacuity table
INSERT INTO visualacuity (id, l_eyevision, r_eyevision)
VALUES
  (2, 20, 18),
  (3, 22, 20),
  (4, 18, 16),
  (5, 24, 22),
  (6, 20, 18),
  (7, 26, 24),
  (8, 18, 16),
  (9, 22, 20),
  (10, 24, 22);

-- Insert entries into the doctorsconsultation table
INSERT INTO doctorsconsultation (id, healthy, consultation_notes, referral_needed)
VALUES
  (10, true, 'Patient is in good health.', false),
  (11, false, 'Patient reported mild cough.', true),
  (12, true, 'No significant health issues reported.', false),
  (13, false, 'Patient complained of joint pain.', true),
  (14, true, 'Routine checkup, no concerns.', false),
  (15, false, 'Patient reported headaches.', true),
  (16, true, 'No health issues reported.', false);

