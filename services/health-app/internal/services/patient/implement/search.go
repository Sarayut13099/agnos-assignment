package implement

import (
	"context"
	"errors"
	"health-app/internal/domain/patient"
	"strings"
)

func (p *patientService) SearchPatients(ctx context.Context, hisUrl string, filters patient.PatientsSearchRequest) ([]*patient.Patient, error) {
	if strings.TrimSpace(hisUrl) == "" {
		return nil, errors.New("hisUrl cannot be empty")
	}

	return p.hisApi.SearchPatients(ctx, hisUrl, filters)
}
