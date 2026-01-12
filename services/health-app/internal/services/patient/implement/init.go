package implement

import (
	"health-app/internal/external"
	"health-app/internal/services/patient"
)

type patientService struct {
	hisApi *external.HISAPIClient
}

func NewPatientService(hisApi *external.HISAPIClient) patient.Services {
	return &patientService{
		hisApi: hisApi,
	}
}
