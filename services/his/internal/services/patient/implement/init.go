package patient

import (
	dbPatient "his/internal/domain/patient"
	"his/internal/services/patient"
)

type patientService struct {
	repo dbPatient.Repository
}

func NewPatientService(repo dbPatient.Repository) patient.Service {
	return &patientService{
		repo: repo,
	}
}
