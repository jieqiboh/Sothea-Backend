-- /* The Actual Setup File to Be Used */
-- Currently commented out to prevent accidental execution
-- /*******************
--     Drop the tables
-- ********************/
-- DROP TABLE IF EXISTS pastmedicalhistory;
-- DROP TABLE IF EXISTS socialhistory;
-- DROP TABLE IF EXISTS vitalstatistics;
-- DROP TABLE IF EXISTS heightandweight;
-- DROP TABLE IF EXISTS visualacuity;
-- DROP TABLE IF EXISTS doctorsconsultation;
-- DROP TABLE IF EXISTS admin;
-- DROP TABLE IF EXISTS users;
--
-- /*******************
-- Create the schema and Load Extensions
-- ********************/
--
-- CREATE TABLE IF NOT EXISTS admin
-- (
--     id                    SERIAL, -- Use SERIAL to auto-increment the ID
--     vid                   INTEGER    NOT NULL,
--     family_group          TEXT       NOT NULL,
--     reg_date              DATE       NOT NULL,
--     queue_no              TEXT       NOT NULL,
--     name                  TEXT       NOT NULL,
--     khmer_name            TEXT       NOT NULL,
--     dob                   DATE,
--     age                   INTEGER,
--     gender                VARCHAR(1) NOT NULL,
--     village               TEXT       NOT NULL,
--     contact_no            TEXT       NOT NULL,
--     pregnant              BOOLEAN    NOT NULL,
--     last_menstrual_period Date,
--     drug_allergies        TEXT,
--     sent_to_id            BOOLEAN    NOT NULL,
--     photo                 BYTEA,
--     PRIMARY KEY (id, vid)         -- Composite primary key
-- );
--
-- CREATE TABLE IF NOT EXISTS pastmedicalhistory
-- (
--     id                           INTEGER NOT NULL,                       -- Use INTEGER to match the id type from admin
--     vid                          INTEGER NOT NULL,                       -- Add vid to match the vid type from admin
--     tuberculosis                 BOOLEAN NOT NULL,
--     diabetes                     BOOLEAN NOT NULL,
--     hypertension                 BOOLEAN NOT NULL,
--     hyperlipidemia               BOOLEAN NOT NULL,
--     chronic_joint_pains          BOOLEAN NOT NULL,
--     chronic_muscle_aches         BOOLEAN NOT NULL,
--     sexually_transmitted_disease BOOLEAN NOT NULL,
--     specified_stds               TEXT,
--     others                       TEXT,
--     PRIMARY KEY (id, vid),                                               -- Composite primary key
--     CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
-- );
--
-- CREATE TABLE IF NOT EXISTS socialhistory
-- (
--     id                      INTEGER NOT NULL,                            -- Use INTEGER to match the id type from admin
--     vid                     INTEGER NOT NULL,                            -- Add vid to match the vid type from admin
--     past_smoking_history    BOOLEAN NOT NULL,
--     no_of_years             INTEGER,
--     current_smoking_history BOOLEAN NOT NULL,
--     cigarettes_per_day      INTEGER,
--     alcohol_history         BOOLEAN NOT NULL,
--     how_regular             VARCHAR(1),
--     PRIMARY KEY (id, vid),                                               -- Composite primary key
--     CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
-- );
--
-- CREATE TABLE IF NOT EXISTS vitalstatistics
-- (
--     id                        INTEGER       NOT NULL,                    -- Use INTEGER to match the id type from admin
--     vid                       INTEGER       NOT NULL,                    -- Add vid to match the vid type from admin
--     temperature               NUMERIC(5, 1) NOT NULL,
--     spo2                      NUMERIC(5, 1) NOT NULL,
--     systolic_bp1              NUMERIC(5, 1) NOT NULL,
--     diastolic_bp1             NUMERIC(5, 1) NOT NULL,
--     systolic_bp2              NUMERIC(5, 1) NOT NULL,
--     diastolic_bp2             NUMERIC(5, 1) NOT NULL,
--     avg_systolic_bp           NUMERIC(5, 1) NOT NULL,
--     avg_diastolic_bp          NUMERIC(5, 1) NOT NULL,
--     hr1                       NUMERIC(5, 1) NOT NULL,
--     hr2                       NUMERIC(5, 1) NOT NULL,
--     avg_hr                    NUMERIC(5, 1) NOT NULL,
--     rand_blood_glucose_mmoll  NUMERIC(5, 1) NOT NULL,
--     PRIMARY KEY (id, vid),                                               -- Composite primary key
--     CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
-- );
--
-- CREATE TABLE IF NOT EXISTS heightandweight
-- (
--     id           INTEGER       NOT NULL,                                 -- Use INTEGER to match the id type from admin
--     vid          INTEGER       NOT NULL,                                 -- Add vid to match the vid type from admin
--     height       NUMERIC(5, 1) NOT NULL,
--     weight       NUMERIC(5, 1) NOT NULL,
--     bmi          NUMERIC(5, 1) NOT NULL,
--     bmi_analysis TEXT          NOT NULL,
--     paeds_height NUMERIC(5, 1) NOT NULL,
--     paeds_weight NUMERIC(5, 1) NOT NULL,
--     PRIMARY KEY (id, vid),                                               -- Composite primary key
--     CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
-- );
--
-- CREATE TABLE IF NOT EXISTS visualacuity
-- (
--     id                      INTEGER NOT NULL,                            -- Use INTEGER to match the id type from admin
--     vid                     INTEGER NOT NULL,                            -- Add vid to match the vid type from admin
--     l_eye_vision            INTEGER NOT NULL,
--     r_eye_vision            INTEGER NOT NULL,
--     additional_intervention TEXT,
--     PRIMARY KEY (id, vid),                                               -- Composite primary key
--     CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
-- );
--
-- CREATE TABLE IF NOT EXISTS doctorsconsultation
-- (
--     id                 INTEGER NOT NULL,                                 -- Use INTEGER to match the id type from admin
--     vid                INTEGER NOT NULL,                                 -- Add vid to match the vid type from admin
--     well            BOOLEAN NOT NULL,
--     msk                BOOLEAN NOT NULL,
--     cvs                BOOLEAN NOT NULL,
--     respi              BOOLEAN NOT NULL,
--     gu                 BOOLEAN NOT NULL,
--     git                BOOLEAN NOT NULL,
--     eye                BOOLEAN NOT NULL,
--     derm               BOOLEAN NOT NULL,
--     others             TEXT    NOT NULL,
--     consultation_notes TEXT,
--     diagnosis          TEXT,
--     treatment          TEXT,
--     referral_needed    BOOLEAN NOT NULL,
--     referral_loc       TEXT,
--     remarks            TEXT,
--     PRIMARY KEY (id, vid),                                               -- Composite primary key
--     CONSTRAINT fk_admin FOREIGN KEY (id, vid) REFERENCES admin (id, vid) -- Foreign key referencing the composite key in admin
-- );
--
-- /*******************
--     Create the trigger function
-- *******************/
--
-- CREATE OR REPLACE FUNCTION set_entry_id() RETURNS TRIGGER AS
-- $$
-- DECLARE
--     max_entry_id INTEGER;
-- BEGIN
--     -- Check if the ID already exists in the table
--     SELECT COALESCE(MAX(VID), 0)
--     INTO max_entry_id
--     FROM admin
--     WHERE ID = NEW.ID;
--
--     -- Increment Entry_ID based on the max_entry_id
--     NEW.VID := max_entry_id + 1;
--
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;
--
-- CREATE TRIGGER before_insert_admin
--     BEFORE INSERT
--     ON admin
--     FOR EACH ROW
-- EXECUTE FUNCTION set_entry_id();
--
-- /*******************
--     Add usernames and passwords
--  */
-- CREATE TABLE users
-- (
--     id SERIAL PRIMARY KEY,              -- Auto-incrementing ID for each user (optional but recommended)
--     username VARCHAR(255) NOT NULL,     -- Username field, with a max length of 255 characters
--     password_hash TEXT NOT NULL       -- Field to store the hashed password
-- );
--
-- -- Insert the users
-- INSERT INTO users (username, password_hash) VALUES ('dr1', '$2a$10$QJVQuKDvgYI7h1bdZ7FeVuo6W/osv1D0QKS80Z49wPYCwl7JM7XYm');
-- INSERT INTO users (username, password_hash) VALUES ('dr2', '$2a$10$YBdJa9dV2GjlbW6To7oWfunZzKm1gUWwa7rRL9oNqai5XPurvXfGu');
-- INSERT INTO users (username, password_hash) VALUES ('dr3', '$2a$10$7AUA6j4Gssu1NeedfH/.nu8qi3mkR/bo3kQ1PQIZgh0CL3J7YqQAW');
-- INSERT INTO users (username, password_hash) VALUES ('dr4', '$2a$10$yVvSf9HcrOCKU20iI3KFMOraUeMeZFz.YAaxy7bC22BiPEUK604C2');
-- INSERT INTO users (username, password_hash) VALUES ('dr5', '$2a$10$KY442JFYtKJQST31XYwJquoD61qhWgpL2ee743oQnLV0IRLY8.sUi');
-- INSERT INTO users (username, password_hash) VALUES ('dr6', '$2a$10$fW.NoRoM6Nxgx/lBLR/s5.Wp4dM2AtecwT.W.QqOj2fu04z2ED4BS');
-- INSERT INTO users (username, password_hash) VALUES ('dr7', '$2a$10$oPJxh3X2nGXqrbuy98tapeBrZ2tIvhvdxSrwo3eZ7BD88Q.gav8MS');
-- INSERT INTO users (username, password_hash) VALUES ('dr8', '$2a$10$eivqzqXzgEdu42kl3w9mveqW03UhEtCdJewRUDncFEfBPDxdfNeGi');
-- INSERT INTO users (username, password_hash) VALUES ('dr9', '$2a$10$iSD4RiAOHsNM90DeA2mkcejotEoQtIfFx5tDfIIwIobgCzncsE0ka');
-- INSERT INTO users (username, password_hash) VALUES ('dr10', '$2a$10$810uwCZxg3PMLmfF2hLSeeYIYTKM5ND3pqd/ui1ZPYO2SfgMTFSdK');
-- INSERT INTO users (username, password_hash) VALUES ('dr11', '$2a$10$WDe8w9bbz7RyBftb.ybamOvRUmpyON79mb6tXM/g7sN77QZY8VLMC');
-- INSERT INTO users (username, password_hash) VALUES ('dr12', '$2a$10$SJroXorL7Bj98/qRE6/FYORtdChWveQIyMjUhOO29xi/z8MYOnGba');
-- INSERT INTO users (username, password_hash) VALUES ('dr13', '$2a$10$UkbEtCJdeGuX/FHLTTlCE.OsrQ7xjVl0rY3Bo8NWY3SOkJK/KT0G.');
-- INSERT INTO users (username, password_hash) VALUES ('dr14', '$2a$10$I0f4b7QtrS4eH9SNETGUsOcwoqQqlQKRbUD2DITs6y5lB/cdlqOzC');
-- INSERT INTO users (username, password_hash) VALUES ('dr15', '$2a$10$UhtnHUmTU4DJTUXVaBR.5OFaS2fiimN6YWhOYTzaofZwMtKF/PKDC');
-- INSERT INTO users (username, password_hash) VALUES ('dr16', '$2a$10$gYTT5hStXUKrpK0Tv9sTluvsK3dwIa8vEJMzPeIuaZ6.MV1YCuJ9O');
-- INSERT INTO users (username, password_hash) VALUES ('dr17', '$2a$10$aIXgj2VmQuyih/KOMxvxeudRI6DRgiYUy1IMvZ3e.xFh2xZcyWpzq');
-- INSERT INTO users (username, password_hash) VALUES ('dr18', '$2a$10$xJQIZNVos4y5LkYekDRHV.KeSgKnEzV71uYN3tYRZ3vaJoDPNGSBK');
-- INSERT INTO users (username, password_hash) VALUES ('dr19', '$2a$10$pvo3eKjP4ZFNx.bE1l9mwuFT87Wnm.FqDHvlJSFbmokP7DkYwy9Oq');
-- INSERT INTO users (username, password_hash) VALUES ('dr20', '$2a$10$7ZKQ56ckZ2Z9Jye1em62iOrKeMLggM5oI2V1Jpjw52Cmg4nKTmNOG');
