package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jieqiboh/sothea_backend/domain"
	"github.com/jieqiboh/sothea_backend/models"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var db *sql.DB

var admin = models.Admin{
	FamilyGroup:   models.PtrTo("S001"),
	RegDate:       models.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Name:          models.PtrTo("John Doe"),
	Dob:           models.PtrTo(time.Date(1994, time.January, 10, 0, 0, 0, 0, time.UTC)),
	Age:           models.PtrTo(30),
	Gender:        models.PtrTo("M"),
	Village:       models.PtrTo("SO"),
	ContactNo:     models.PtrTo("12345678"),
	Pregnant:      models.PtrTo(false),
	DrugAllergies: models.PtrTo("panadol"),
	SentToID:      models.PtrTo(false),
}
var pastmedicalhistory = models.PastMedicalHistory{
	Tuberculosis:               models.PtrTo(true),
	Diabetes:                   models.PtrTo(false),
	Hypertension:               models.PtrTo(true),
	Hyperlipidemia:             models.PtrTo(false),
	ChronicJointPains:          models.PtrTo(false),
	ChronicMuscleAches:         models.PtrTo(true),
	SexuallyTransmittedDisease: models.PtrTo(true),
	SpecifiedSTDs:              models.PtrTo("TRICHOMONAS"),
	Others:                     nil,
}
var socialhistory = models.SocialHistory{
	PastSmokingHistory:    models.PtrTo(true),
	NumberOfYears:         models.PtrTo(int32(15)),
	CurrentSmokingHistory: models.PtrTo(false),
	CigarettesPerDay:      nil,
	AlcoholHistory:        models.PtrTo(true),
	HowRegular:            models.PtrTo("A"),
}
var vitalstatistics = models.VitalStatistics{
	Temperature:              models.PtrTo(36.5),
	SpO2:                     models.PtrTo(98.0),
	SystolicBP1:              models.PtrTo(120.0),
	DiastolicBP1:             models.PtrTo(80.0),
	SystolicBP2:              models.PtrTo(122.0),
	DiastolicBP2:             models.PtrTo(78.0),
	AverageSystolicBP:        models.PtrTo(121.0),
	AverageDiastolicBP:       models.PtrTo(79.0),
	HR1:                      models.PtrTo(72.0),
	HR2:                      models.PtrTo(71.0),
	AverageHR:                models.PtrTo(71.5),
	RandomBloodGlucoseMmolL:  models.PtrTo(5.4),
	RandomBloodGlucoseMmolLp: models.PtrTo(5.3),
}
var heightandweight = models.HeightAndWeight{
	Height:      models.PtrTo(170.0),
	Weight:      models.PtrTo(70.0),
	BMI:         models.PtrTo(24.2),
	BMIAnalysis: models.PtrTo("normal weight"),
	PaedsHeight: models.PtrTo(90.0),
	PaedsWeight: models.PtrTo(80.0),
}
var visualacuity = models.VisualAcuity{
	LEyeVision:             models.PtrTo(int32(20)),
	REyeVision:             models.PtrTo(int32(20)),
	AdditionalIntervention: models.PtrTo("VISUAL FIELD TEST REQUIRED"),
}
var doctorsconsultation = models.DoctorsConsultation{
	Healthy:           models.PtrTo(true),
	Msk:               models.PtrTo(false),
	Cvs:               models.PtrTo(false),
	Respi:             models.PtrTo(true),
	Gu:                models.PtrTo(true),
	Git:               models.PtrTo(false),
	Eye:               models.PtrTo(true),
	Derm:              models.PtrTo(false),
	Others:            models.PtrTo(false),
	ConsultationNotes: models.PtrTo("CHEST PAIN, SHORTNESS OF BREATH, COUGH"),
	Diagnosis:         models.PtrTo("ACUTE BRONCHITIS"),
	Treatment:         models.PtrTo("REST, HYDRATION, COUGH SYRUP"),
	ReferralNeeded:    models.PtrTo(false),
	ReferralLoc:       nil,
	Remarks:           models.PtrTo("MONITOR FOR RESOLUTION"),
}

func initDb() {
	// Initialize global variables
	viper.SetConfigFile(`../../../config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbName := viper.GetString(`database.name`)
	dbSslMode := viper.GetString(`database.sslmode`)

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbName, dbSslMode)

	// Open a database connection
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// You might want to check the connection here to handle errors
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
	result := make([]models.Admin, 0)

	for rows.Next() {
		admin := models.Admin{}
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

	patient := domain.Patient{
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

	patient := domain.Patient{
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

	updatePatient := domain.Patient{
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

	expectedPatient := domain.Patient{
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
	initDb()
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	// InsertPatient with only admin and docconsult field - should successfully insert
	patient1 := domain.Patient{
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
	patient2 := domain.Patient{
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
	patient4 := domain.Patient{
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
	updatedAdmin := models.Admin{
		FamilyGroup:         models.PtrTo("S001"),
		RegDate:             models.PtrTo(time.Now()),
		Name:                models.PtrTo("Updated Name Here"),
		Dob:                 models.PtrTo(time.Date(1994, time.January, 10, 0, 0, 0, 0, time.UTC)),
		Age:                 models.PtrTo(5),
		Gender:              models.PtrTo("F"),
		Village:             models.PtrTo("SO"),
		ContactNo:           models.PtrTo("12345678"),
		Pregnant:            models.PtrTo(false),
		LastMenstrualPeriod: models.PtrTo(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
		DrugAllergies:       models.PtrTo("panadol"),
		SentToID:            models.PtrTo(false),
	}
	patient6 := domain.Patient{
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

//func TestGetAllAdmin(t *testing.T) {
//	initDb()
//	repo := NewPostgresPatientRepository(db)
//
//	patient_repo, ok := repo.(*postgresPatientRepository)
//	if !ok {
//		log.Fatal("Failed to assert repo")
//	}
//
//	adminArray, err := patient_repo.GetAllFromAdmin(context.Background())
//	assert.Nil(t, err)
//	assert.NotNil(t, adminArray)
//}
