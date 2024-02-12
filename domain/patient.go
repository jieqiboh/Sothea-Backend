package domain

import (
	"context"
	"github.com/jieqiboh/sothea_backend/models"
)

// Use pointers so that some structs are optional
type Patient struct {
	Admin               *models.Admin               `json:"admin"`
	PastMedicalHistory  *models.PastMedicalHistory  `json:"pastmedicalhistory"`
	SocialHistory       *models.SocialHistory       `json:"socialhistory"`
	VitalStatistics     *models.VitalStatistics     `json:"vitalstatistics"`
	HeightAndWeight     *models.HeightAndWeight     `json:"heightandweight"`
	VisualAcuity        *models.VisualAcuity        `json:"visualacuity"`
	DoctorsConsultation *models.DoctorsConsultation `json:"doctorsconsultation"`
}

type PatientUseCase interface {
	GetPatientByID(ctx context.Context, id int32) (*Patient, error) // Takes an ID and returns a Patient object
	DeletePatientByID(ctx context.Context, id int32) (int32, error) // Deletes a Patient by ID
	UpdatePatientByID(ctx context.Context, id int32, patient *Patient) (int32, error)
	InsertPatient(ctx context.Context, patient *Patient) (int32, error) // Creates a new patient and inserts in database
	GetAllFromAdmin(ctx context.Context) ([]models.Admin, error)
}

type PatientRepository interface {
	GetPatientByID(ctx context.Context, id int32) (*Patient, error) // Takes an ID and returns a Patient object
	DeletePatientByID(ctx context.Context, id int32) (int32, error) // Deletes a Patient by ID
	UpdatePatientByID(ctx context.Context, id int32, patient *Patient) (int32, error)
	InsertPatient(ctx context.Context, patient *Patient) (int32, error) // Creates a new patient and inserts in database
	GetAllFromAdmin(ctx context.Context) ([]models.Admin, error)        // returns all entries from a given category
}
