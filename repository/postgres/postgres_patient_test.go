package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jieqiboh/sothea_backend/entities"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var db *sql.DB

var admin = entities.Admin{
	FamilyGroup:   entities.PtrTo("S001"),
	RegDate:       entities.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Name:          entities.PtrTo("John Doe"),
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

func initDb() {
	// Initialize global variables
	viper.SetConfigFile(`../../config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbName := viper.GetString(`database.name`)
	dbPassword := viper.GetString(`database.password`)
	dbSslMode := viper.GetString(`database.sslmode`)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, dbSslMode)

	// Open a database connection
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
}

// Gets Patient with highest ID from db, deletes it, then inserts another Patient
func TestRun(t *testing.T) {
	t.Run("TestDB", TestDB)
	t.Run("TestGetPatientByID", TestGetPatientByID)
	t.Run("TestDeletePatientByID", TestDeletePatientByID)
	t.Run("TestInsertPatient", TestInsertPatient)
	t.Run("TestUpdatePatientByID", TestUpdatePatientByID)
	//t.Run("TestInsertAdmin", TestInsertAdmin)
}

func TestDB(t *testing.T) {
	initDb()
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
			&admin.Dob,
			&admin.Age,
			&admin.Gender,
			&admin.Village,
			&admin.ContactNo,
			&admin.Pregnant,
			&admin.LastMenstrualPeriod,
			&admin.DrugAllergies,
			&admin.SentToID,
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

	log.Println(result[0])
}

func TestGetPatientByID(t *testing.T) {
	initDb()
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
	initDb()
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
	initDb()
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
	initDb()
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
	initDb()
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

	// Update Patient6 admin id, should fail
	//updatedAdmin = models.Admin{
	//	ID: int64(10000),
	//}
	//patient7 := domain.Patient{
	//	Admin: &updatedAdmin,
	//}
	//id7, err := patient_repo.UpdatePatientByID(context.Background(), id6, &patient7)
	//assert.Equal(t, int64(-1), id7)
	//assert.NotNil(t, err)

	// Update a non-existing patient, should fail

	// Delete Patient4

}
