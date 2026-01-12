package implement

import (
	"context"
	"errors"
	"health-app/internal/domain/patient"
	"strings"
)

func (p *patientService) GetByID(ctx context.Context, hisUrl, id string) (*patient.Patient, error) {
	if strings.TrimSpace(hisUrl) == "" {
		return nil, errors.New("hisUrl cannot be empty")
	}

	return p.hisApi.GetPatientByID(ctx, hisUrl, id)
}
