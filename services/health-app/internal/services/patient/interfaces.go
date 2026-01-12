package patient

import (
	"context"
	"health-app/internal/domain/patient"
)

type Services interface {
	GetByID(ctx context.Context, hisUrl, id string) (*patient.Patient, error)
	SearchPatients(ctx context.Context, hisUrl string, filters patient.PatientsSearchRequest) ([]*patient.Patient, error)
}
