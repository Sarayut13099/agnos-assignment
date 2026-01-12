package patient

import "his/internal/services/patient"

type PatientHandler struct {
	service patient.Service
}

func NewPatientHandler(service patient.Service) *PatientHandler {
	return &PatientHandler{
		service: service,
	}
}
