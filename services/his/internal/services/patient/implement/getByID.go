package patient

import (
	"context"
	"his/internal/domain/patient"
)

func (s *patientService) GetByID(ctx context.Context, id string) (*patient.Patient, error) {
	return s.repo.GetByID(ctx, id)
}
