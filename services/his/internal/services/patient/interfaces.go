package patient

import (
	"context"
	"his/internal/domain/patient"
)

//go:generate mockery --name=Service
type Service interface {
	GetByID(ctx context.Context, id string) (*patient.Patient, error)
	SearchPatients(ctx context.Context, filters patient.PatientsSearchRequest) ([]*patient.Patient, error)
}
