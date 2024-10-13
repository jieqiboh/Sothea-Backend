# Backend Developer Documentation
### Last Updated: August 29, 2024

For future developers, this document will serve as a guide to understanding the backend and how to work with it, as well as the numerous design choices made.

## Overview
Sothea-Backend provides a REST API for Sothea-Frontend to interact with, to do CRUD operations on patients. 
It is written in Go, due to its speed, reliability and ease of use. The backend uses a PostgreSQL database to store patient data.  
It also draws inspiration from [Bob's Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). 
Some key concepts used include Dependency Rules and use cases, which make testing and maintenance easier.
For basic deployment instructions, refer to the README.md file. 

## Entity Design 
The center of the backend is the patients entity, a representation of a patient's data. A patient comprises the following categories, each with their own fields:  
Patient SQL schema: `/sql/patients_setup.sql`
Patient Golang struct schema: `/entities`
- Admin
- Past Medical History
- Social History
- Vital Statistics
- Height and Weight
- Visual Acuity
- Doctor's Consultation

Patients go through the physical health screening stations with the admin station first, and the rest having no guaranteed order.
Hence, if a patient exists, they will have an admin row, but may not have the other categories present yet.  
Additionally, patients may have multiple visits, with multiple rows for each category, representing the previous years of visits.  
Every row in the patient database will have the following structure:  
```
+------------+----------+-----------------------+  
| patient_id | visit_id | rest of categories    |  
+------------+----------+-----------------------+  
```
The patient id is used to uniquely identify a patient, while the visit id is used to narrow down which visit the visit is associated with.
The choice to use an visit id to identify the visit an visit is associated with, instead of the date was made to allow for easier querying, since the Date data type might be tricker to work with.

#### Admin Category:  

| Attribute               | Type    | Description                                                                                                         |
|-------------------------|---------|---------------------------------------------------------------------------------------------------------------------|
| `id`                    | integer | Auto-incrementing with each new patient created                                                                     |
| `vid`                   | integer | Auto-incrementing with each new visit created, for each respective patient                                          |
| `family_group`          | string  | Generally a string of the form "S001", "S002A" identifying the village and family number, not enforced in the db.   |
| `reg_date`              | date    | Date of registration. Set by the doctors.                                                                           |
| `queue_no`              | string  | String that can be of the form "1A", "2A", etc, not enforced                                                        |
| `name`                  | string  | Patient Name                                                                                                        |
| `khmer_name`            | string  | Patient Khmer Name                                                                                                  |
| `dob`                   | date    | Patient date of birth                                                                                               |
| `age`                   | integer | Patient age                                                                                                         |
| `gender`                | char(1) | Patient gender. Not enforced M or F.                                                                                |
| `village`               | string  | Patient village of origin.                                                                                          |
| `contact_no`            | string  | Patient contact number                                                                                              |
| `pregnant`              | boolean | Boolean indicating if patient is pregnant.                                                                          |
| `last_menstrual_period` | date    | Date of last menstrual period.                                                                                      |
| `drug_allergies`        | string  | Patient drug allergies                                                                                              |
| `sent_to_id`            | boolean | Boolean indicating if patient was sent to infectious diseases station.                                              |
| `photo`                 | bytea   | Postgres binary string data type, used to store photo                                                               |
Optional Fields: `dob`, `age`, `contact_no`, `photo`

#### Past Medical History Category:

| Attribute                      | Type    | Description                                                                |
|--------------------------------|---------|----------------------------------------------------------------------------|
| `id`                           | integer | Auto-incrementing with each new patient created                            |
| `vid`                          | integer | Auto-incrementing with each new visit created, for each respective patient |
| `diabetes`                     | boolean | Boolean indicating if patient has diabetes                                 |
| `hypertension`                 | boolean | Boolean indicating if patient has hypertension                             |
| `hyperlipidemia`               | boolean | Boolean indicating if patient has hyperlipidemia                           |
| `chronic_joint_pains`          | boolean | Boolean indicating if patient has chronic joint pains                      |
| `chronic_muscle_aches`         | boolean | Boolean indicating if patient has chronic muscle aches                     |
| `sexually_transmitted_disease` | boolean | Boolean indicating if patient has a sexually transmitted disease           |
| `specified_stds`               | string  | Details of specified sexually transmitted diseases, if any                 |
| `others`                       | string  | Any other relevant medical conditions or notes                             |
Optional Fields: `specified_stds`, `others`

#### Social History Category:

| Attribute                 | Type    | Description                                                                  |
|---------------------------|---------|------------------------------------------------------------------------------|
| `id`                      | integer | Auto-incrementing primary key                                                |
| `vid`                     | integer | Auto-incrementing with each new visit created, for each respective patient   |
| `past_smoking_history`    | boolean | Boolean indicating if the patient has a history of smoking                   |
| `no_of_years`             | integer | Number of years the patient has smoked                                       |
| `current_smoking_history` | boolean | Boolean indicating if the patient currently smokes                           |
| `cigarettes_per_day`      | integer | Number of cigarettes the patient smokes per day                              |
| `alcohol_history`         | boolean | Boolean indicating if the patient has a history of alcohol consumption       |
| `how_regular`             | string  | Char 'A', 'B', 'C, 'D' indicating patient's frequency of alcohol consumption |
Optional Fields: `no_of_years`, `cigarettes_per_day`, `how_regular`'

#### Vital Statistics Category:

| Attribute                   | Type    | Description                                                                |
|-----------------------------|---------|----------------------------------------------------------------------------|
| `id`                        | integer | Auto-incrementing primary key                                              |
| `vid`                       | integer | Auto-incrementing with each new visit created, for each respective patient |
| `temperature`               | float   | Patient's body temperature in degrees Celsius, 5 digits with 1 dp          |
| `spo2`                      | float   | Oxygen saturation level in percentage                                      |
| `systolic_bp1`              | float   | Systolic blood pressure measurement 1                                      |
| `diastolic_bp1`             | float   | Diastolic blood pressure measurement 1                                     |
| `systolic_bp2`              | float   | Systolic blood pressure measurement 2                                      |
| `diastolic_bp2`             | float   | Diastolic blood pressure measurement 2                                     |
| `avg_systolic_bp`           | float   | Average systolic blood pressure                                            |
| `avg_diastolic_bp`          | float   | Average diastolic blood pressure                                           |
| `hr1`                       | float   | Heart rate measurement 1                                                   |
| `hr2`                       | float   | Heart rate measurement 2                                                   |
| `avg_hr`                    | float   | Average heart rate                                                         |
| `rand_blood_glucose_mmolL`  | float   | Random blood glucose level in mmol/L                                       |
| `rand_blood_glucose_mmolLp` | float   | Random blood glucose level in mmol/Lp                                      |
Optional Fields: -

#### Height and Weight Category:

| Attribute      | Type    | Description                                                                |
|----------------|---------|----------------------------------------------------------------------------|
| `id`           | integer | Auto-incrementing primary key                                              |
| `vid`          | integer | Auto-incrementing with each new visit created, for each respective patient |
| `height`       | float   | Patient's height in meters (with 1 decimal place)                          |
| `weight`       | float   | Patient's weight in kilograms (with 1 decimal place)                       |
| `bmi`          | float   | Body Mass Index (BMI) value (with 1 decimal place)                         |
| `bmi_analysis` | string  | Analysis or interpretation of the BMI value                                |
| `paeds_height` | float   | Pediatric height measurement in meters (with 1 decimal place)              |
| `paeds_weight` | float   | Pediatric weight measurement in kilograms (with 1 decimal place)           |
Optional Fields: `paeds_height`, `paeds_weight`

#### Visual Acuity Category:

| Attribute                 | Type    | Description                                                                |
|---------------------------|---------|----------------------------------------------------------------------------|
| `id`                      | integer | Auto-incrementing primary key                                              |
| `vid`                     | integer | Auto-incrementing with each new visit created, for each respective patient |
| `l_eye_vision`            | integer | Vision measurement for the left eye                                        |
| `r_eye_vision`            | integer | Vision measurement for the right eye                                       |
| `additional_intervention` | string  | Details of any additional intervention or notes                            |
Optional Fields: `additional_intervention`

#### Doctors Consultation Category:

| Attribute            | Type    | Description                                                                |
|----------------------|---------|----------------------------------------------------------------------------|
| `id`                 | integer | Auto-incrementing primary key                                              |
| `vid`                | integer | Auto-incrementing with each new visit created, for each respective patient |
| `msk`                | boolean | Boolean indicating if musculoskeletal system is healthy                    |
| `cvs`                | boolean | Boolean indicating if cardiovascular system is healthy                     |
| `respi`              | boolean | Boolean indicating if respiratory system is healthy                        |
| `gu`                 | boolean | Boolean indicating if                                                      |
| `git`                | boolean | Boolean indicating if                                                      |
| `eye`                | boolean | Boolean indicating if eye health is satisfactory                           |
| `derm`               | boolean | Boolean indicating if dermatological health is satisfactory                |
| `others`             | string  | Details of other relevant health assessments                               |
| `consultation_notes` | string  | General notes from the consultation                                        |
| `diagnosis`          | string  | Diagnosis provided during consultation                                     |
| `treatment`          | string  | Treatment plan or recommendations                                          |
| `referral_needed`    | boolean | Boolean indicating if a referral is needed                                 |
| `referral_loc`       | string  | Location or details for the referral, if needed                            |
| `remarks`            | string  | Additional remarks or comments                                             |
Optional Fields: `others`, `consultation_notes`, `diagnosis`, `treatment`, `referral_loc`, `remarks`

## Database and Database Schema
The database used is PostgreSQL. To interact with the db, the Golang database driver used is [lib/pq](https://github.com/lib/pq),
with raw SQL statements. We have decided against using an ORM due to the interesting nature of the patient entity, which may or may not have all categories filled out, save for the admin category.
Safer alternatives such as [sqlx](https://github.com/jmoiron/sqlx) or [sqlcl](https://github.com/sqlc-dev/sqlc) could be considered in the future, but we decided to stick with raw SQL for now.

Additionally, we've tried to keep the schema simple, choosing not to compute derived fields at the database or backend level.

## Directory Structure
```
.
├── README.md - Contains basic instructions for setting up the backend.
├── DOCS.md - Hello future developer! This is the backend documentation file.
├── go.mod 
├── go.sum 
├── .gitignore - Contains files to be ignored by git.
├── config.json - Contains configuration for backend in development mode.
├── prod.json - Contains configuration for backend in production mode.
├── Dockerfile - Contains the Dockerfile for the database.
├── main.go - Entry point for the backend.
├── .github - Folder containing github actions.
├── controllers
│   ├── middleware 
│   │   └── auth.go - Middleware for authentication.
│   ├── patient_handler.go - Handles patient requests.
│   └── login_handler.go - Handles login requests.
├── entities - Contains struct definitions for the patient entity and errors.
│   ├── admin.go 
│   ├── pastmedicalhistory.go
│   ├── socialhistory.go
│   ├── vitalstatistics.go
│   ├── heightandweight.go
│   ├── visualacuity.go
│   ├── doctorsconsultation.go
│   ├── patient.go - Contains the patient struct definition encompassing all categories, as well as PatientUseCase and PatientRepository interfaces.
│   ├── patientmeta.go - Contains the patientmeta struct definition.
│   ├── patientvisitmeta.go - Contains the patientvisitmeta struct definition.
│   ├── user.go - Contains the user struct definition.
│   ├── errors.go - Contains custom error definitions.
│   └── helper.go - Contains helper functions for the entities.
├── mocks - Folder containing autogenerated mocks and dummy patient data for testing.
├── repository 
│   ├── postgres
│   │   └── postgres_patient.go - Contains the postgres implementation of the PatientRepository interface. 
│   └── tmp - Contains output.csv files, a temporary file generated when exporting patient data.
├── sql 
│   └── patients_setup.sql - Contains the SQL schema for the patients table.
├── usecases
│   ├── patient_ucase.go - Contains the PatientUseCase interface and its implementation.
│   └── login_ucase.go - Contains the LoginUseCase interface and its implementation.
├── util 
│   └── helper.go - Contains general helper functions for the backend.
```
Test files have been excluded from the directory structure for brevity.

## Error Handling
We have defined some custom errors in `entities/errors.go`.  
They serve to make passing errors around easier, and to provide more context to the error, such as whether a Patient or PatientVisit was not found.

HTTP Error Codes Used:
- 400: Bad Request
- 401: Unauthorized
- 404: Not Found
- 500: Internal Server Error

## Middleware
The backend uses a simple authentication middleware to check if the user is authenticated in the `controllers/middleware/auth.go` file.
The middleware checks for the presence of a JWT token in the Authorization header, and verifies it using the secret key in the config file.

## Configuration
The backend uses a `config.json` or `prod.json` file to store configuration settings, and extracts the values on startup using Viper.
The configuration file helps to keep sensitive information such as the database URL and JWT secret key secure.

## Testing
The backend uses the standard Go testing package for testing.
Testing is done at the controller and repository levels, with the use of mocks to simulate the usecases.

For testing at the controller level, we used [vektra/mockery](https://github.com/vektra/mockery) to generate mocks for the usecases and repositories.
The mocks are stored in the `mocks` folder, and are used in the controller tests.  
For testing at the repository level, we opted for using [Dockertest](https://github.com/ory/dockertest) to spin up a temporary PostgreSQL container for testing, and to run the tests against it.
This is to ensure that the data access layer, which is far more complex to mock, performs exactly as expected.
See why you probably shouldn't mock the database [here](https://dominikbraun.io/blog/you-probably-shouldnt-mock-the-database/).

## Naming and Code Conventions
SQL Fields: snake_case  
e.g. `consultation_notes`  

Golang Struct Fields: CamelCase with first letter capitalised  
e.g. `RegDate`  

JSON Fields: camelCase  
e.g. `pastSmokingHistory`  

## Miscellaneous Design Choices
### Using Pointers instead of Defined Null Types
We have chosen to use pointers for fields that can be null in the database, such as `dob`, `age`, `sent_to_id`, `photo`, etc.  
While null types exist in the [lib/pq](https://github.com/lib/pq) library, we have chosen to use pointers instead.
For example, instead of using `sql.NullString`, we use `*string` for fields that can be null.
This works at the data access layer, but fails at the controller layer where JSON marshalling happens.  

For primitive types like bool, they can only take on 2 values: true, false. When I marshal JSON into a struct, null fields get converted into the false value for booleans. For `binding: required` gotags, this gets treated as the field being null.
For both optional and not null fields, this means the validator cannot tell if a value is false because it is explicitly assigned to be false, or because it was null, and got marshalled into the false value.

The only workaround is to use a pointer, which allows it to take on a third value, nil.

Additionally, I updated String() methods to use SafeDeref, a helper method, to dereference pointers to their respective null types if needed instead of just nil, which cannot be printed.


### Export to DB Feature
An important feature of the backend is for users to be able to easily export the patient data to a CSV file.
Due to the lack of support for CSV exports in the lib/pq database drivers, we had to implement the feature manually.  
~~One of the approaches used was to execute the `COPY` command in the PostgresQL server to generate a csv file, leveraging the in-built feature. However, due to storage being isolated in the Docker container, an additional volume mount is needed to access the generated file. 
While it is a little messy, this method doesn't require us to use an external client such as psql, and doesn't us to handle the messy typecasting of data to strings.~~
We have since moved to using a simple existing golang library, [sqltocsv](https://github.com/joho/sqltocsv), which writes the rows to a csv file easily.

### Import to DB Feature
This feature allows importing existing CSV files to the database. The CSV file must be in the correct format, with the correct headers, and the correct order of columns.  
Note: This feature is not accessible to users since it should be rarely invoked, only to reset the database to a known state in the event of system failures :O  

Some desired properties of the import to DB feature include:
- Ability to use exported csv files from /export-db as a backup
- Easy to use and debug (If I ever have to use this feature, I need it to be quick to minimise downtime)