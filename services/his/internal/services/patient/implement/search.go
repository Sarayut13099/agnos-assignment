package patient

import (
	"context"
	"his/internal/domain/patient"
	"his/internal/utils"
)

func (s *patientService) SearchPatients(ctx context.Context, filters patient.PatientsSearchRequest) ([]*patient.Patient, error) {
	newDateOfBirth, err := utils.ParseDatePtr(filters.DateOfBirth)
	if err != nil {
		return nil, err
	}

	mappedFilters := patient.PatientSearchFilter{
		NationalID:  filters.NationalID,
		PassportID:  filters.PassportID,
		FirstName:   filters.FirstName,
		MiddleName:  filters.MiddleName,
		LastName:    filters.LastName,
		DateOfBirth: newDateOfBirth,
		PhoneNumber: filters.PhoneNumber,
		Email:       filters.Email,
	}
	return s.repo.SearchPatients(ctx, mappedFilters)
}
