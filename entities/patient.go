package entities

import (
	"context"
)

// Use pointers so that some structs are optional
type Patient struct {
	Admin               *Admin               `json:"admin"`
	PastMedicalHistory  *PastMedicalHistory  `json:"pastmedicalhistory"`
	SocialHistory       *SocialHistory       `json:"socialhistory"`
	VitalStatistics     *VitalStatistics     `json:"vitalstatistics"`
	HeightAndWeight     *HeightAndWeight     `json:"heightandweight"`
	VisualAcuity        *VisualAcuity        `json:"visualacuity"`
	DoctorsConsultation *DoctorsConsultation `json:"doctorsconsultation"`
}

type PatientUseCase interface {
	GetPatientByID(ctx context.Context, id int32) (*Patient, error) // Takes an ID and returns a Patient object
	DeletePatientByID(ctx context.Context, id int32) (int32, error) // Deletes a Patient by ID
	UpdatePatientByID(ctx context.Context, id int32, patient *Patient) (int32, error)
	InsertPatient(ctx context.Context, patient *Patient) (int32, error) // Creates a new patient and inserts in database
	GetAllAdmin(ctx context.Context) ([]PartAdmin, error)
}

type PatientRepository interface {
	GetPatientByID(ctx context.Context, id int32) (*Patient, error) // Takes an ID and returns a Patient object
	DeletePatientByID(ctx context.Context, id int32) (int32, error) // Deletes a Patient by ID
	UpdatePatientByID(ctx context.Context, id int32, patient *Patient) (int32, error)
	InsertPatient(ctx context.Context, patient *Patient) (int32, error) // Creates a new patient and inserts in database
	GetAllAdmin(ctx context.Context) ([]PartAdmin, error)
}
