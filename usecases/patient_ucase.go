package usecases

import (
	"context"
	"github.com/jieqiboh/sothea_backend/entities"
	"time"
)

type patientUsecase struct {
	patientRepo    entities.PatientRepository
	contextTimeout time.Duration
}

// NewPatientUseCase
func NewPatientUsecase(p entities.PatientRepository, timeout time.Duration) entities.PatientUseCase {
	return &patientUsecase{
		patientRepo:    p,
		contextTimeout: timeout,
	}
}

func (p patientUsecase) GetPatientByID(ctx context.Context, id int32) (*entities.Patient, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.patientRepo.GetPatientByID(ctx, id)
}

func (p patientUsecase) InsertPatient(ctx context.Context, patient *entities.Patient) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.patientRepo.InsertPatient(ctx, patient)
}

func (p patientUsecase) DeletePatientByID(ctx context.Context, id int32) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.patientRepo.DeletePatientByID(ctx, id)
}

func (p patientUsecase) UpdatePatientByID(ctx context.Context, id int32, patient *entities.Patient) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.patientRepo.UpdatePatientByID(ctx, id, patient)
}

func (p patientUsecase) GetAllAdmin(ctx context.Context) ([]entities.PartAdmin, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.patientRepo.GetAllAdmin(ctx)
}
