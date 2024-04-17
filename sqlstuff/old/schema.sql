/******************* 
Create the schema and Load Extensions
********************/ 

CREATE TABLE IF NOT EXISTS admin (
    ID SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    family_group VARCHAR(32) NOT NULL,
    reg_date DATE NOT NULL,
    name VARCHAR(32),
    age INTEGER NOT NULL,
    gender VARCHAR(3) NOT NULL
);

CREATE TABLE IF NOT EXISTS pastmedicalhistory (
    id SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    tuberculosis BOOLEAN NOT NULL,
    diabetes BOOLEAN NOT NULL,
    hypertension BOOLEAN NOT NULL,
    hyperlipidemia BOOLEAN NOT NULL,
    chronicjointpains BOOLEAN NOT NULL,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);

CREATE TABLE IF NOT EXISTS socialhistory (
    id SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    past_smoking_history BOOLEAN NOT NULL,
    no_of_years INTEGER,
    current_smoking_history BOOLEAN NOT NULL,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);

CREATE TABLE IF NOT EXISTS vitalstatistics (
    id SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    temperature NUMERIC(5, 1) NOT NULL,
    spO2 NUMERIC(5, 1) NOT NULL,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);

CREATE TABLE IF NOT EXISTS heightandweight (
    id SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    height NUMERIC(5, 1) NOT NULL,
    weight NUMERIC(5, 1) NOT NULL,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);

CREATE TABLE IF NOT EXISTS visualacuity (
    id SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    l_eyevision INTEGER NOT NULL,
    r_eyevision INTEGER NOT NULL,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);

CREATE TABLE IF NOT EXISTS doctorsconsultation (
    id SERIAL PRIMARY KEY, -- Use SERIAL to auto-increment the ID
    healthy BOOLEAN NOT NULL,
    consultation_notes TEXT,
    referral_needed BOOLEAN NOT NULL,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES admin (id)
);
