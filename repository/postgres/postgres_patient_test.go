package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jieqiboh/sothea_backend/entities"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var admin = entities.Admin{
	FamilyGroup:   entities.PtrTo("S001"),
	RegDate:       entities.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Name:          entities.PtrTo("John Doe"),
	KhmerName:     entities.PtrTo("១២៣៤ ៥៦៧៨៩០ឥឲ"),
	Dob:           entities.PtrTo(time.Date(1994, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Age:           entities.PtrTo(30),
	Gender:        entities.PtrTo("M"),
	Village:       entities.PtrTo("SO"),
	ContactNo:     entities.PtrTo("12345678"),
	Pregnant:      entities.PtrTo(false),
	DrugAllergies: entities.PtrTo("panadol"),
	SentToID:      entities.PtrTo(false),
}
var pastmedicalhistory = entities.PastMedicalHistory{
	Tuberculosis:               entities.PtrTo(true),
	Diabetes:                   entities.PtrTo(false),
	Hypertension:               entities.PtrTo(true),
	Hyperlipidemia:             entities.PtrTo(false),
	ChronicJointPains:          entities.PtrTo(false),
	ChronicMuscleAches:         entities.PtrTo(true),
	SexuallyTransmittedDisease: entities.PtrTo(true),
	SpecifiedSTDs:              entities.PtrTo("TRICHOMONAS"),
	Others:                     nil,
}
var socialhistory = entities.SocialHistory{
	PastSmokingHistory:    entities.PtrTo(true),
	NumberOfYears:         entities.PtrTo(int32(15)),
	CurrentSmokingHistory: entities.PtrTo(false),
	CigarettesPerDay:      nil,
	AlcoholHistory:        entities.PtrTo(true),
	HowRegular:            entities.PtrTo("A"),
}
var vitalstatistics = entities.VitalStatistics{
	Temperature:              entities.PtrTo(36.5),
	SpO2:                     entities.PtrTo(98.0),
	SystolicBP1:              entities.PtrTo(120.0),
	DiastolicBP1:             entities.PtrTo(80.0),
	SystolicBP2:              entities.PtrTo(122.0),
	DiastolicBP2:             entities.PtrTo(78.0),
	AverageSystolicBP:        entities.PtrTo(121.0),
	AverageDiastolicBP:       entities.PtrTo(79.0),
	HR1:                      entities.PtrTo(72.0),
	HR2:                      entities.PtrTo(71.0),
	AverageHR:                entities.PtrTo(71.5),
	RandomBloodGlucoseMmolL:  entities.PtrTo(5.4),
	RandomBloodGlucoseMmolLp: entities.PtrTo(5.3),
}
var heightandweight = entities.HeightAndWeight{
	Height:      entities.PtrTo(170.0),
	Weight:      entities.PtrTo(70.0),
	BMI:         entities.PtrTo(24.2),
	BMIAnalysis: entities.PtrTo("normal weight"),
	PaedsHeight: entities.PtrTo(90.0),
	PaedsWeight: entities.PtrTo(80.0),
}
var visualacuity = entities.VisualAcuity{
	LEyeVision:             entities.PtrTo(int32(20)),
	REyeVision:             entities.PtrTo(int32(20)),
	AdditionalIntervention: entities.PtrTo("VISUAL FIELD TEST REQUIRED"),
}
var doctorsconsultation = entities.DoctorsConsultation{
	Healthy:           entities.PtrTo(true),
	Msk:               entities.PtrTo(false),
	Cvs:               entities.PtrTo(false),
	Respi:             entities.PtrTo(true),
	Gu:                entities.PtrTo(true),
	Git:               entities.PtrTo(false),
	Eye:               entities.PtrTo(true),
	Derm:              entities.PtrTo(false),
	Others:            entities.PtrTo(false),
	ConsultationNotes: entities.PtrTo("CHEST PAIN, SHORTNESS OF BREATH, COUGH"),
	Diagnosis:         entities.PtrTo("ACUTE BRONCHITIS"),
	Treatment:         entities.PtrTo("REST, HYDRATION, COUGH SYRUP"),
	ReferralNeeded:    entities.PtrTo(false),
	ReferralLoc:       nil,
	Remarks:           entities.PtrTo("MONITOR FOR RESOLUTION"),
}

var db *sql.DB

const (
	PostgresPassword = "postgres"
	PostgresUser     = "jieqiboh"
	PostgresDB       = "patients"
)

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// Define the run options
	runOptions := &dockertest.RunOptions{
		Name:       "sothea-db", // Container name
		Repository: "sothea-db", // Image name
		Tag:        "latest",    // Image tag
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432/tcp": {{HostIP: "", HostPort: "5432"}},
		},
	}

	// Build and run the Docker container
	resource, err := pool.BuildAndRunWithOptions("../../Dockerfile", runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", PostgresUser, PostgresPassword,
		hostAndPort, PostgresDB)

	log.Println("Connecting to database on url: ", databaseUrl)

	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func runMigrations(migrationsPath string, db *sql.DB) error {
	if migrationsPath == "" {
		return errors.New("missing sql path")
	}
	driver, err := pg.WithInstance(db, &pg.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+migrationsPath, "postgres", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func TestDB(t *testing.T) {
	id := 1
	rows, err := db.Query("SELECT * FROM admin WHERE id = $1", id)
	result := make([]entities.Admin, 0)

	for rows.Next() {
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
		if err != nil {
			panic(err)
		}
		result = append(result, admin)
	}

	assert.NotNil(t, *result[0].Dob)
	assert.NotNil(t, result[0].Gender)
	assert.NotNil(t, result[0].Village)
	assert.NotNil(t, result[0].ContactNo)
	assert.NotNil(t, result[0].Pregnant)
	assert.Nil(t, result[0].LastMenstrualPeriod)
	assert.NotNil(t, result[0].DrugAllergies)
	assert.NotNil(t, result[0].SentToID)
	assert.NotNil(t, result[0].Photo)

	log.Println(result[0])
}

func TestGetPatientByID(t *testing.T) {
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	id := 1
	p, err := patient_repo.GetPatientByID(context.Background(), int32(id))
	assert.Nil(t, err)

	log.Println(p.Admin)
	log.Println(p.PastMedicalHistory)
	log.Println(p.SocialHistory)
	log.Println(p.VitalStatistics)
	log.Println(p.HeightAndWeight)
	log.Println(p.VisualAcuity)
	log.Println(p.DoctorsConsultation)
}

func TestDeletePatientByID(t *testing.T) {
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	var latestId int32
	err := db.QueryRow("SELECT ID FROM admin ORDER BY ID DESC LIMIT 1").Scan(&latestId)
	if err != nil {
		log.Fatal("Getting latest id failed:", err)
	}

	res, err := patient_repo.DeletePatientByID(context.Background(), latestId)

	assert.Nil(t, err)
	log.Println("Deleted Patient of ID: ", res)
}

func TestInsertPatient(t *testing.T) {
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	patient := entities.Patient{
		Admin:               &admin,
		PastMedicalHistory:  &pastmedicalhistory,
		SocialHistory:       &socialhistory,
		VitalStatistics:     &vitalstatistics,
		HeightAndWeight:     &heightandweight,
		VisualAcuity:        &visualacuity,
		DoctorsConsultation: &doctorsconsultation,
	}

	id, err := patient_repo.InsertPatient(context.Background(), &patient)
	assert.Nil(t, err)
	log.Println("Patient Successfully inserted with id: ", id)

}

func TestUpdatePatientByID(t *testing.T) {
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	patient := entities.Patient{
		Admin:               &admin,
		PastMedicalHistory:  nil,
		SocialHistory:       nil,
		VitalStatistics:     nil,
		HeightAndWeight:     &heightandweight,
		VisualAcuity:        &visualacuity,
		DoctorsConsultation: &doctorsconsultation,
	}

	id, err := patient_repo.InsertPatient(context.Background(), &patient)
	assert.Nil(t, err)
	log.Println("Patient Successfully inserted with id: ", id)
	admin.ID = id
	pastmedicalhistory.ID = id
	socialhistory.ID = id
	vitalstatistics.ID = id
	heightandweight.ID = id
	visualacuity.ID = id
	doctorsconsultation.ID = id

	updatePatient := entities.Patient{
		Admin:               &admin,
		PastMedicalHistory:  &pastmedicalhistory,
		SocialHistory:       &socialhistory,
		VitalStatistics:     &vitalstatistics,
		HeightAndWeight:     &heightandweight,
		VisualAcuity:        &visualacuity,
		DoctorsConsultation: nil,
	}
	id, err = patient_repo.UpdatePatientByID(context.Background(), id, &updatePatient)

	assert.Nil(t, err)
	log.Println("Patient Successfully Updated with id: ", id)

	p, err := patient_repo.GetPatientByID(context.Background(), id)
	assert.Nil(t, err)

	expectedPatient := entities.Patient{
		Admin:               &admin,
		PastMedicalHistory:  &pastmedicalhistory,
		SocialHistory:       &socialhistory,
		VitalStatistics:     &vitalstatistics,
		HeightAndWeight:     &heightandweight,
		VisualAcuity:        &visualacuity,
		DoctorsConsultation: &doctorsconsultation,
	}
	// Time is not equal
	assert.Equal(t, p.PastMedicalHistory, expectedPatient.PastMedicalHistory)
	assert.Equal(t, p.SocialHistory, expectedPatient.SocialHistory)
	assert.Equal(t, p.VitalStatistics, expectedPatient.VitalStatistics)
	assert.Equal(t, p.HeightAndWeight, expectedPatient.HeightAndWeight)
	assert.Equal(t, p.VisualAcuity, expectedPatient.VisualAcuity)
	assert.Equal(t, p.DoctorsConsultation, expectedPatient.DoctorsConsultation)
	log.Println(p.Admin)
	log.Println(p.PastMedicalHistory)
	log.Println(p.SocialHistory)
	log.Println(p.VitalStatistics)
	log.Println(p.HeightAndWeight)
	log.Println(p.VisualAcuity)
	log.Println(p.DoctorsConsultation)
}

func TestFull(t *testing.T) {
	// Tests some edge cases and ensures desired behaviour is maintained
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	// InsertPatient with only admin and docconsult field - should successfully insert
	patient1 := entities.Patient{
		Admin:               &admin,
		PastMedicalHistory:  nil,
		SocialHistory:       nil,
		VitalStatistics:     nil,
		HeightAndWeight:     nil,
		VisualAcuity:        nil,
		DoctorsConsultation: &doctorsconsultation,
	}
	id1, err := patient_repo.InsertPatient(context.Background(), &patient1)
	assert.NotNil(t, id1)
	assert.Nil(t, err)

	// InsertPatient with no admin field - should return -1 and error
	patient2 := entities.Patient{
		Admin:               nil,
		PastMedicalHistory:  nil,
		SocialHistory:       nil,
		VitalStatistics:     &vitalstatistics,
		HeightAndWeight:     nil,
		VisualAcuity:        nil,
		DoctorsConsultation: &doctorsconsultation,
	}
	id2, err := patient_repo.InsertPatient(context.Background(), &patient2)
	assert.Equal(t, id2, int32(-1))
	assert.NotNil(t, err)

	// GetPatient with id that doesn't exist
	patient3, err := patient_repo.GetPatientByID(context.Background(), -1)
	assert.Nil(t, patient3)
	assert.NotNil(t, err)

	// GetPatient with only admin and socialhistory field filled in
	patient4 := entities.Patient{
		Admin:         &admin,
		SocialHistory: &socialhistory,
	}
	id4, err := patient_repo.InsertPatient(context.Background(), &patient4)
	var latestId int32
	err = db.QueryRow("SELECT ID FROM admin ORDER BY ID DESC LIMIT 1").Scan(&latestId)
	if err != nil {
		log.Fatal("Getting latest id failed:", err)
	}
	patient5, err := patient_repo.GetPatientByID(context.Background(), latestId)
	assert.Nil(t, err)
	assert.Nil(t, patient5.PastMedicalHistory)
	assert.Nil(t, patient5.VitalStatistics)
	assert.Nil(t, patient5.HeightAndWeight)
	assert.Nil(t, patient5.VisualAcuity)
	assert.Nil(t, patient5.DoctorsConsultation)
	assert.NotNil(t, patient5.Admin)
	assert.NotNil(t, patient5.SocialHistory)

	// Update Patient4 and update all admin fields except id, add vitalstatistics and visualacuity
	updatedAdmin := entities.Admin{
		FamilyGroup:         entities.PtrTo("S001"),
		RegDate:             entities.PtrTo(time.Now()),
		Name:                entities.PtrTo("Updated Name Here"),
		KhmerName:           entities.PtrTo("ចវបនមឦ។៊"),
		Dob:                 entities.PtrTo(time.Date(1994, time.January, 10, 0, 0, 0, 0, time.UTC)),
		Age:                 entities.PtrTo(5),
		Gender:              entities.PtrTo("F"),
		Village:             entities.PtrTo("SO"),
		ContactNo:           entities.PtrTo("12345678"),
		Pregnant:            entities.PtrTo(false),
		LastMenstrualPeriod: entities.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
		DrugAllergies:       entities.PtrTo("panadol"),
		SentToID:            entities.PtrTo(false),
	}
	patient6 := entities.Patient{
		Admin:           &updatedAdmin,
		VitalStatistics: &vitalstatistics,
		VisualAcuity:    &visualacuity,
	}
	id6, err := patient_repo.UpdatePatientByID(context.Background(), id4, &patient6)
	assert.NotNil(t, id6)
	assert.Nil(t, err)

	updatedPatient6, err := patient_repo.GetPatientByID(context.Background(), id6)
	assert.Nil(t, err)
	assert.NotNil(t, updatedPatient6.Admin)
	assert.NotNil(t, updatedPatient6.SocialHistory)
	assert.NotNil(t, updatedPatient6.VitalStatistics)
	assert.NotNil(t, updatedPatient6.VisualAcuity)
	assert.Nil(t, updatedPatient6.PastMedicalHistory)
	assert.Nil(t, updatedPatient6.HeightAndWeight)
	assert.Nil(t, updatedPatient6.DoctorsConsultation)
	assert.Equal(t, updatedAdmin.FamilyGroup, updatedPatient6.Admin.FamilyGroup)
	assert.Equal(t, updatedAdmin.Name, updatedPatient6.Admin.Name)
	assert.Equal(t, updatedAdmin.Age, updatedPatient6.Admin.Age)
	assert.Equal(t, updatedAdmin.Gender, updatedPatient6.Admin.Gender)
}

func TestGetAllAdmin(t *testing.T) {
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	admins, err := patient_repo.GetAllAdmin(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, admins)
}

func TestPostgresPatientRepository_SearchPatients(t *testing.T) {
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	admins, err := patient_repo.SearchPatients(context.Background(), "១២៣៤ ៥៦៧៨៩០ឥឲ")
	assert.Nil(t, err)
	assert.NotNil(t, admins)
}
