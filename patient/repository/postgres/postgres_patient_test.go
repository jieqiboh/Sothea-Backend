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
	FamilyGroup: "S001",
	RegDate:     time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
	Name:        "John Doe",
	Dob:         time.Date(1994, time.January, 10, 0, 0, 0, 0, time.UTC),
	Age:         30,
	Gender:      "M",
	Village:     "SO",
	ContactNo:   "12345678",
	Pregnant:    false,
	DrugAllergies: sql.NullString{
		String: "panadol",
		Valid:  false,
	},
	SentToID: false,
}
var pastmedicalhistory = models.PastMedicalHistory{
	Tuberculosis:               true,
	Diabetes:                   false,
	Hypertension:               true,
	Hyperlipidemia:             false,
	ChronicJointPains:          false,
	ChronicMuscleAches:         true,
	SexuallyTransmittedDisease: true,
	SpecifiedSTDs: sql.NullString{
		String: "TRICHOMONAS",
		Valid:  true,
	},
	Others: sql.NullString{
		String: "",
		Valid:  false,
	},
}
var socialhistory = models.SocialHistory{
	PastSmokingHistory: true,
	NumberOfYears: sql.NullInt32{
		Int32: 15,
		Valid: true,
	},
	CurrentSmokingHistory: false,
	CigarettesPerDay: sql.NullInt32{
		Int32: 0,
		Valid: false,
	},
	AlcoholHistory: true,
	HowRegular: sql.NullString{
		String: "A",
		Valid:  true,
	},
}
var vitalstatistics = models.VitalStatistics{
	Temperature:              36.5,
	SpO2:                     98,
	SystolicBP1:              120,
	DiastolicBP1:             80,
	SystolicBP2:              122,
	DiastolicBP2:             78,
	AverageSystolicBP:        121,
	AverageDiastolicBP:       79,
	HR1:                      72,
	HR2:                      71,
	AverageHR:                71.5,
	RandomBloodGlucoseMmolL:  5.4,
	RandomBloodGlucoseMmolLp: 5.3,
}
var heightandweight = models.HeightAndWeight{
	Height:      170,
	Weight:      70,
	BMI:         24.2,
	BMIAnalysis: "normal weight",
	PaedsHeight: 90,
	PaedsWeight: 80,
}
var visualacuity = models.VisualAcuity{
	LEyeVision: 20,
	REyeVision: 20,
	AdditionalIntervention: sql.NullString{
		String: "VISUAL FIELD TEST REQUIRED",
		Valid:  true,
	},
}
var doctorsconsultation = models.DoctorsConsultation{
	Healthy: true,
	Msk:     false,
	Cvs:     false,
	Respi:   true,
	Gu:      true,
	Git:     false,
	Eye:     true,
	Derm:    false,
	Others:  false,
	ConsultationNotes: sql.NullString{
		String: "CHEST PAIN, SHORTNESS OF BREATH, COUGH",
		Valid:  true,
	},
	Diagnosis: sql.NullString{
		String: "ACUTE BRONCHITIS",
		Valid:  true,
	},
	Treatment: sql.NullString{
		String: "REST, HYDRATION, COUGH SYRUP",
		Valid:  true,
	},
	ReferralNeeded: false,
	ReferralLoc: sql.NullString{
		String: "",
		Valid:  false,
	},
	Remarks: sql.NullString{
		String: "MONITOR FOR RESOLUTION",
		Valid:  true,
	},
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
	t.Run("TestDeletePatientByID", TestDeletePatientByID)
	t.Run("TestInsertPatient", TestInsertPatient)
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

	assert.NotNil(t, result[0].Dob)
	assert.NotNil(t, result[0].Gender)
	assert.NotNil(t, result[0].Village)
	assert.NotNil(t, result[0].ContactNo)
	assert.NotNil(t, result[0].Pregnant)
	lastMenstrualPeriod, _ := result[0].LastMenstrualPeriod.Value()
	assert.Nil(t, lastMenstrualPeriod)
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

//func TestUpdatePatientByID(t *testing.T) {
//	initDb()
//	repo := NewPostgresPatientRepository(db)
//
//	patient_repo, ok := repo.(*postgresPatientRepository)
//	if !ok {
//		log.Fatal("Failed to assert repo")
//	}
//
//	patient := domain.Patient{
//		Admin:               &admin,
//		PastMedicalHistory:  nil,
//		SocialHistory:       nil,
//		VitalStatistics:     nil,
//		HeightAndWeight:     &heightandweight,
//		VisualAcuity:        &visualacuity,
//		DoctorsConsultation: &doctorsconsultation,
//	}
//
//	id, err := patient_repo.InsertPatient(context.Background(), &patient)
//	assert.Nil(t, err)
//	log.Println("Patient Successfully inserted with id: ", id)
//	admin.ID = id
//	pastmedicalhistory.ID = id
//	socialhistory.ID = id
//	vitalstatistics.ID = id
//	heightandweight.ID = id
//	visualacuity.ID = id
//	doctorsconsultation.ID = id
//
//	updatePatient := domain.Patient{
//		Admin:               &admin,
//		PastMedicalHistory:  &pastmedicalhistory,
//		SocialHistory:       &socialhistory,
//		VitalStatistics:     &vitalstatistics,
//		HeightAndWeight:     &heightandweight,
//		VisualAcuity:        &visualacuity,
//		DoctorsConsultation: nil,
//	}
//	id, err = patient_repo.UpdatePatientByID(context.Background(), id, &updatePatient)
//
//	assert.Nil(t, err)
//	log.Println("Patient Successfully Updated with id: ", id)
//
//	p, err := patient_repo.GetPatientByID(context.Background(), id)
//	assert.Nil(t, err)
//
//	expectedPatient := domain.Patient{
//		Admin:               &admin,
//		PastMedicalHistory:  &pastmedicalhistory,
//		SocialHistory:       &socialhistory,
//		VitalStatistics:     &vitalstatistics,
//		HeightAndWeight:     &heightandweight,
//		VisualAcuity:        &visualacuity,
//		DoctorsConsultation: &doctorsconsultation,
//	}
//	// Time is not equal
//	assert.Equal(t, p.PastMedicalHistory, expectedPatient.PastMedicalHistory)
//	assert.Equal(t, p.SocialHistory, expectedPatient.SocialHistory)
//	assert.Equal(t, p.VitalStatistics, expectedPatient.VitalStatistics)
//	assert.Equal(t, p.HeightAndWeight, expectedPatient.HeightAndWeight)
//	assert.Equal(t, p.VisualAcuity, expectedPatient.VisualAcuity)
//	assert.Equal(t, p.DoctorsConsultation, expectedPatient.DoctorsConsultation)
//	log.Println(p.Admin.ToString())
//	log.Println(p.PastMedicalHistory.ToString())
//	log.Println(p.SocialHistory.ToString())
//	log.Println(p.VitalStatistics.ToString())
//	log.Println(p.HeightAndWeight.ToString())
//	log.Println(p.VisualAcuity.ToString())
//	log.Println(p.DoctorsConsultation.ToString())
//}

//func TestFull(t *testing.T) {
//	// Tests some edge cases and ensures desired behaviour is maintained
//	initDb()
//	repo := NewPostgresPatientRepository(db)
//
//	patient_repo, ok := repo.(*postgresPatientRepository)
//	if !ok {
//		log.Fatal("Failed to assert repo")
//	}
//
//	// InsertPatient with only admin and docconsult field - should return id and nil error
//	patient1 := domain.Patient{
//		Admin:               &admin,
//		PastMedicalHistory:  nil,
//		SocialHistory:       nil,
//		VitalStatistics:     nil,
//		HeightAndWeight:     nil,
//		VisualAcuity:        nil,
//		DoctorsConsultation: &doctorsconsultation,
//	}
//	id1, err := patient_repo.InsertPatient(context.Background(), &patient1)
//	assert.NotNil(t, id1)
//	assert.Nil(t, err)
//
//	// InsertPatient with no admin field - should return -1 and error
//	patient2 := domain.Patient{
//		Admin:               nil,
//		PastMedicalHistory:  nil,
//		SocialHistory:       nil,
//		VitalStatistics:     &vitalstatistics,
//		HeightAndWeight:     nil,
//		VisualAcuity:        nil,
//		DoctorsConsultation: &doctorsconsultation,
//	}
//	id2, err := patient_repo.InsertPatient(context.Background(), &patient2)
//	assert.Equal(t, id2, int64(-1))
//	assert.NotNil(t, err)
//
//	// GetPatient with id that doesn't exist
//	patient3, err := patient_repo.GetPatientByID(context.Background(), -1)
//	assert.Nil(t, patient3)
//	assert.NotNil(t, err)
//
//	// GetPatient with only admin and socialhistory field filled in
//	patient4 := domain.Patient{
//		Admin:         &admin,
//		SocialHistory: &socialhistory,
//	}
//	id4, err := patient_repo.InsertPatient(context.Background(), &patient4)
//	var latestId int64
//	err = db.QueryRow("SELECT ID FROM admin ORDER BY ID DESC LIMIT 1").Scan(&latestId)
//	if err != nil {
//		log.Fatal("Getting latest id failed:", err)
//	}
//	patient5, err := patient_repo.GetPatientByID(context.Background(), latestId)
//	assert.Nil(t, err)
//	assert.Nil(t, patient5.PastMedicalHistory)
//	assert.Nil(t, patient5.VitalStatistics)
//	assert.Nil(t, patient5.HeightAndWeight)
//	assert.Nil(t, patient5.VisualAcuity)
//	assert.Nil(t, patient5.DoctorsConsultation)
//	assert.NotNil(t, patient5.Admin)
//	assert.NotNil(t, patient5.SocialHistory)
//
//	// Update Patient4 and update all admin fields except id, add vitalstatistics and visualacuity
//	updatedAdmin := models.Admin{
//		FamilyGroup: "S001",
//		RegDate:     time.Now(),
//		Name:        "Updated Name Here",
//		Age:         5,
//		Gender:      "F",
//	}
//	patient6 := domain.Patient{
//		Admin:           &updatedAdmin,
//		VitalStatistics: &vitalstatistics,
//		VisualAcuity:    &visualacuity,
//	}
//	id6, err := patient_repo.UpdatePatientByID(context.Background(), id4, &patient6)
//	assert.NotNil(t, id6)
//	assert.Nil(t, err)
//
//	updatedPatient6, err := patient_repo.GetPatientByID(context.Background(), id6)
//	assert.Nil(t, err)
//	assert.NotNil(t, updatedPatient6.Admin)
//	assert.NotNil(t, updatedPatient6.SocialHistory)
//	assert.NotNil(t, updatedPatient6.VitalStatistics)
//	assert.NotNil(t, updatedPatient6.VisualAcuity)
//	assert.Nil(t, updatedPatient6.PastMedicalHistory)
//	assert.Nil(t, updatedPatient6.HeightAndWeight)
//	assert.Nil(t, updatedPatient6.DoctorsConsultation)
//	assert.Equal(t, updatedAdmin.FamilyGroup, updatedPatient6.Admin.FamilyGroup)
//	assert.Equal(t, updatedAdmin.Name, updatedPatient6.Admin.Name)
//	assert.Equal(t, updatedAdmin.Age, updatedPatient6.Admin.Age)
//	assert.Equal(t, updatedAdmin.Gender, updatedPatient6.Admin.Gender)
//
//	// Update Patient6 admin id, should fail
//	//updatedAdmin = models.Admin{
//	//	ID: int64(10000),
//	//}
//	//patient7 := domain.Patient{
//	//	Admin: &updatedAdmin,
//	//}
//	//id7, err := patient_repo.UpdatePatientByID(context.Background(), id6, &patient7)
//	//assert.Equal(t, int64(-1), id7)
//	//assert.NotNil(t, err)
//
//	// Update a non-existing patient, should fail
//
//	// Delete Patient4
//
//}

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
