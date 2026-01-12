package patient

import (
	hospitalsv "health-app/internal/services/hospital"
	patientsv "health-app/internal/services/patient"
)

type PatientHandler struct {
	service         patientsv.Services
	hospitalService hospitalsv.Services
}

func NewPatientHandler(service patientsv.Services, hospitalService hospitalsv.Services) *PatientHandler {
	return &PatientHandler{
		service:         service,
		hospitalService: hospitalService,
	}
}
