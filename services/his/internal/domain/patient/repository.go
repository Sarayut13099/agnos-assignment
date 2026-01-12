package patient

import "context"

type Repository interface {
	GetByID(ctx context.Context, id string) (*Patient, error)
	SearchPatients(ctx context.Context, filters PatientSearchFilter) ([]*Patient, error)
}
