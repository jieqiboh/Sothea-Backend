package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	entities "github.com/jieqiboh/sothea_backend/entities"
	"github.com/jieqiboh/sothea_backend/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	Others:            entities.PtrTo("TRICHOMONAS VAGINALIS"),
	ConsultationNotes: entities.PtrTo("CHEST PAIN, SHORTNESS OF BREATH, COUGH"),
	Diagnosis:         entities.PtrTo("ACUTE BRONCHITIS"),
	Treatment:         entities.PtrTo("REST, HYDRATION, COUGH SYRUP"),
	ReferralNeeded:    entities.PtrTo(false),
	ReferralLoc:       nil,
	Remarks:           entities.PtrTo("MONITOR FOR RESOLUTION"),
}

var patient = entities.Patient{
	Admin:               &admin,
	PastMedicalHistory:  &pastmedicalhistory,
	SocialHistory:       &socialhistory,
	VitalStatistics:     &vitalstatistics,
	HeightAndWeight:     &heightandweight,
	VisualAcuity:        &visualacuity,
	DoctorsConsultation: &doctorsconsultation,
}

// NewPatientHandler will initialize the patients/ resources endpoint, excluding the auth middleware
func newTestPatientHandler(e *gin.Engine, us entities.PatientUseCase) {
	handler := &PatientHandler{
		Usecase: us,
	}

	e.GET("/patient/:id", handler.GetPatientByID)
	e.POST("/patient", handler.InsertPatient)
	e.DELETE("/patient/:id", handler.DeletePatientByID)
	e.PATCH("/patient/:id", handler.UpdatePatientByID)
	e.GET("/get-all-admin", handler.GetAllAdmin)
	e.GET("/search-patients/:search-name", handler.SearchPatients)
}

// Success - 200 OK
// Patient Not Found - 404 Not Found
func TestGetPatientByID_Success(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	num := 1
	mockUsecase.On("GetPatientByID", context.Background(), int32(num)).Return(&patient, nil)
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/patient/1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetPatientByID_NotFound(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	num := -1
	mockUsecase.On("GetPatientByID", context.Background(), int32(num)).Return(nil, entities.ErrPatientNotFound)
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/patient/-1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

// Success - 200 OK
// Missing Admin Field - 400 Bad Request
// Invalid Parameters - 400 Bad Request
// Json Marshalling Error - 400 Bad Request
// Empty Request Body - 400 Bad Request
func TestInsertPatient_Success(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	id := int32(7)
	mockUsecase.On("InsertPatient", context.Background(), &mocks.ValidPatient).Return(id, nil)
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/patient", strings.NewReader(mocks.ValidPatientJson))

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestInsertPatient_MissingAdmin_Failure(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	mockUsecase.On("InsertPatient", context.Background(), &mocks.MissingAdminPatient).Return(int32(-1), entities.ErrMissingAdminCategory)
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/patient", strings.NewReader(mocks.MissingAdminPatientJson))

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestInsertPatient_InvalidParameters_Failure(t *testing.T) {
	// When ShouldBindJSON fails - e.g. A required field is not present
	var mockUsecase mocks.PatientUseCase
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/patient", strings.NewReader(mocks.InvalidParametersPatientJson))

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestInsertPatient_JSONMarshallingError_Failure(t *testing.T) {
	// When marshalling into a JSON struct fails - e.g. data type mismatch between the JSON and the expected golang struct field type
	var mockUsecase mocks.PatientUseCase
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/patient", strings.NewReader(mocks.JSONMarshallingErrorPatientJson))

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestInsertPatient_EmptyRequestBody_Failure(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/patient", strings.NewReader(""))

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

// Success - 200 OK
// Patient Not Found - 404 Not Found
func TestDeletePatientByID_Success(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	num := 1
	mockUsecase.On("DeletePatientByID", context.Background(), int32(num)).Return(int32(num), nil)
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/patient/1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeletePatientByID_NotFound_Failure(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	num := -1
	mockUsecase.On("DeletePatientByID", context.Background(), int32(num)).Return(int32(-1), entities.ErrPatientNotFound)
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/patient/-1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

// Success - 200 OK
// Patient Not Found - 404 Not Found
// Invalid Parameters - 400 Bad Request
// Json Marshalling Error - 400 Bad Request
// Empty Request Body - 400 Bad Request
func TestUpdatePatientByID_Success(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	id := int32(1)
	mockUsecase.On("UpdatePatientByID", context.Background(), id, &mocks.MissingAdminPatient).Return(id, nil)
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/patient/1", strings.NewReader(mocks.MissingAdminPatientJson))

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdatePatientByID_NotFound_Failure(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	id := int32(-1)
	mockUsecase.On("UpdatePatientByID", context.Background(), id, &mocks.ValidPatient).Return(int32(-1), entities.ErrPatientNotFound)
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/patient/-1", strings.NewReader(mocks.ValidPatientJson))

	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestUpdatePatientByID_InvalidParameters_Failure(t *testing.T) {
	// When ShouldBindJSON fails - e.g. A required field is not present
	var mockUsecase mocks.PatientUseCase
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/patient/1", strings.NewReader(mocks.InvalidParametersPatientJson))

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestUpdatePatientByID_JSONMarshallingError_Failure(t *testing.T) {
	// When marshalling into a JSON struct fails - e.g. data type mismatch between the JSON and the expected golang struct field type
	var mockUsecase mocks.PatientUseCase
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/patient/1", strings.NewReader(mocks.JSONMarshallingErrorPatientJson))

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestUpdatePatientByID_EmptyRequestBody_Failure(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/patient/1", strings.NewReader(""))

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

// Success - 200 OK
func TestGetAllAdmin_Success(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	mockUsecase.On("GetAllAdmin", context.Background()).Return(mocks.AdminArray, nil)
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get-all-admin", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

// Success - 200 OK
// Empty Search Name - 400 Bad Request
func TestSearchPatients_Success(t *testing.T) {
	var mockUsecase mocks.PatientUseCase
	mockUsecase.On("SearchPatients", context.Background(), "១២៣៤ ៥៦៧៨៩០ឥ").Return(mocks.AdminArray, nil)
	router := gin.Default()
	newTestPatientHandler(router, &mockUsecase)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search-patients/១២៣៤ ៥៦៧៨៩០ឥ", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
