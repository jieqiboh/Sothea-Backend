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
	GetPatientVisit(ctx context.Context, id int32, vid int32) (*Patient, error)          // Takes an ID and returns a Patient object
	CreatePatient(ctx context.Context, admin *Admin) (int32, error)                      // Creates a new patient and returns their id
	CreatePatientVisit(ctx context.Context, id int32, admin *Admin) (int32, error)       // Creates a new Patient Visit for an existing patient
	DeletePatientVisit(ctx context.Context, id int32, vid int32) error                   // Deletes a Patient Visit
	UpdatePatientVisit(ctx context.Context, id int32, vid int32, patient *Patient) error // Updates a Patient Visit
	GetAllAdmin(ctx context.Context) ([]PartAdmin, error)
	SearchPatients(ctx context.Context, search string) ([]PartAdmin, error)
	ExportDatabaseToCSV(ctx context.Context) error
}

type PatientRepository interface {
	GetPatientVisit(ctx context.Context, id int32, vid int32) (*Patient, error)          // Takes ID and VID returns a Patient Entry
	CreatePatient(ctx context.Context, admin *Admin) (int32, error)                      // Creates a new patient
	CreatePatientVisit(ctx context.Context, id int32, admin *Admin) (int32, error)       // Takes ID, and creates a new Patient Entry
	DeletePatientVisit(ctx context.Context, id int32, vid int32) error                   // Takes ID and VID, and deletes a Patient's Entry
	UpdatePatientVisit(ctx context.Context, id int32, vid int32, patient *Patient) error // Takes ID and VID, and updates a Patient's Entry
	//GetPatientMeta(ctx context.Context, id int32) (*Patient, error)                      // Returns a Patient's Meta Data
	//GetPatientVisitMeta(ctx context.Context, id int32, vid int32) (*Patient, error)      // Returns a Patient's Visit's Meta Data
	//GetAllPatientVisitMeta(ctx context.Context) ([]PartAdmin, error)                     // Returns all Patient's Visit's Meta Data
	GetAllAdmin(ctx context.Context) ([]PartAdmin, error)
	SearchPatients(ctx context.Context, search string) ([]PartAdmin, error)
	ExportDatabaseToCSV(ctx context.Context) error
}
