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
	RegDate:     time.Now(),
	Name:        "Charlie Taylor",
	Age:         20,
	Gender:      "M",
}
var pastmedicalhistory = models.PastMedicalHistory{
	Tuberculosis:      models.BoolPtr(true),
	Diabetes:          models.BoolPtr(false),
	Hypertension:      models.BoolPtr(false),
	Hyperlipidemia:    models.BoolPtr(true),
	ChronicJointPains: models.BoolPtr(false),
}
var socialhistory = models.SocialHistory{
	PastSmokingHistory: models.BoolPtr(false),
	NumberOfYears: sql.NullInt32{
		Int32: 0,
		Valid: false,
	},
	CurrentSmokingHistory: models.BoolPtr(true),
}
var vitalstatistics = models.VitalStatistics{
	Temperature: 98.5,
	SpO2:        98.0,
}
var heightandweight = models.HeightAndWeight{
	Height: 170.6,
	Weight: 55.5,
}
var visualacuity = models.VisualAcuity{
	LEyeVision: 20,
	REyeVision: 18,
}
var doctorsconsultation = models.DoctorsConsultation{
	Healthy:           models.BoolPtr(false),
	ConsultationNotes: sql.NullString{},
	ReferralNeeded:    models.BoolPtr(false),
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
	//t.Run("TestInsertPastMedicalHistory", TestInsertPastMedicalHistory)
	//t.Run("TestInsertSocialHistory", TestInsertSocialHistory)
	//t.Run("TestInsertVitalStatistics", TestInsertVitalStatistics)
	//t.Run("TestInsertHeightAndWeight", TestInsertHeightAndWeight)
	//t.Run("TestInsertVisualAcuity", TestInsertVisualAcuity)
	//t.Run("TestInsertDoctorsConsultation", TestInsertDoctorsConsultation)
	//t.Run("TestFetchAdminFromID", TestFetchAdminFromID)
	//t.Run("TestFetchPastMedicalHistoryFromID", TestFetchPastMedicalHistoryFromID)
	//t.Run("TestFetchSocialHistoryFromID", TestFetchSocialHistoryFromID)
	//t.Run("TestFetchVitalStatisticsFromID", TestFetchVitalStatisticsFromID)
	//t.Run("TestFetchHeightAndWeightFromID", TestFetchHeightAndWeightFromID)
	//t.Run("TestFetchVisualAcuityFromID", TestFetchVisualAcuityFromID)
	//t.Run("TestFetchDoctorsConsultationFromID", TestFetchDoctorsConsultationFromID)
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
			&admin.Age,
			&admin.Gender,
		)
		if err != nil {
			panic(err)
		}
		result = append(result, admin)
	}

	assert.NotNil(t, result[0].Name)
	assert.NotNil(t, result[0].Age)
	assert.NotNil(t, result[0].RegDate)
	assert.NotNil(t, result[0].FamilyGroup)
	assert.NotNil(t, result[0].ID)
	log.Println(result[0].ToString())
}

func TestGetPatientByID(t *testing.T) {
	initDb()
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	id := 10
	p, err := patient_repo.GetPatientByID(context.Background(), int64(id))
	assert.Nil(t, err)

	log.Println(p.Admin)
}

func TestDeletePatientByID(t *testing.T) {
	initDb()
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	var latestId int64
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
	log.Println(p.Admin.ToString())
	log.Println(p.PastMedicalHistory.ToString())
	log.Println(p.SocialHistory.ToString())
	log.Println(p.VitalStatistics.ToString())
	log.Println(p.HeightAndWeight.ToString())
	log.Println(p.VisualAcuity.ToString())
	log.Println(p.DoctorsConsultation.ToString())
}

func TestFull(t *testing.T) {
	// Tests some edge cases and ensures desired behaviour is maintained
	initDb()
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	// InsertPatient with only admin and docconsult field - should return id and nil error
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
	assert.Equal(t, id2, int64(-1))
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
	var latestId int64
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
		FamilyGroup: "S001",
		RegDate:     time.Now(),
		Name:        "Updated Name Here",
		Age:         5,
		Gender:      "F",
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

func TestGetAllAdmin(t *testing.T) {
	initDb()
	repo := NewPostgresPatientRepository(db)

	patient_repo, ok := repo.(*postgresPatientRepository)
	if !ok {
		log.Fatal("Failed to assert repo")
	}

	adminArray, err := patient_repo.GetAllFromAdmin(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, adminArray)
}
