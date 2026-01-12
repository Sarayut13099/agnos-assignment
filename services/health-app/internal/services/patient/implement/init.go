package implement

import (
	"health-app/internal/external"
	"health-app/internal/services/patient"
)

type patientService struct {
	hisApi external.HISAPI
}

func NewPatientService(hisApi external.HISAPI) patient.Services {
	return &patientService{
		hisApi: hisApi,
	}
}
