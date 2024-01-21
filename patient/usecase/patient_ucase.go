package usecase

import (
	"context"
	"github.com/jieqiboh/sothea_backend/domain"
	"github.com/jieqiboh/sothea_backend/models"
	"time"
)

type patientUsecase struct {
	patientRepo    domain.PatientRepository
	contextTimeout time.Duration
}

// NewPatientUseCase
func NewPatientUsecase(p domain.PatientRepository, timeout time.Duration) domain.PatientUseCase {
	return &patientUsecase{
		patientRepo:    p,
		contextTimeout: timeout,
	}
}

func (p patientUsecase) GetPatientByID(ctx context.Context, id int64) (*domain.Patient, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.patientRepo.GetPatientByID(ctx, id)
}

func (p patientUsecase) InsertPatient(ctx context.Context, patient *domain.Patient) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.patientRepo.InsertPatient(ctx, patient)
}

func (p patientUsecase) DeletePatientByID(ctx context.Context, id int64) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.patientRepo.DeletePatientByID(ctx, id)
}

func (p patientUsecase) UpdatePatientByID(ctx context.Context, id int64, patient *domain.Patient) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.patientRepo.UpdatePatientByID(ctx, id, patient)
}

func (p patientUsecase) GetAllFromAdmin(ctx context.Context) ([]models.Admin, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.patientRepo.GetAllFromAdmin(ctx)
}
