package patient_test

import (
	"context"
	"his/internal/domain/patient"
	"his/internal/domain/patient/mocks"
	patientSvc "his/internal/services/patient/implement"
	"his/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPatientService_SearchPatients_Success(t *testing.T) {
	// arrange
	ctx := context.Background()
	repo := new(mocks.Repository)

	dob := "1990-01-01"

	req := patient.PatientsSearchRequest{
		NationalID:  "1101700000011",
		FirstName:   "จอห์น",
		LastName:    "โด",
		DateOfBirth: dob,
	}

	parsedDOB, _ := utils.ParseDatePtr(dob)

	expectedFilter := patient.PatientSearchFilter{
		NationalID:  "1101700000011",
		FirstName:   "จอห์น",
		LastName:    "โด",
		DateOfBirth: parsedDOB,
	}

	expectedResult := []*patient.Patient{
		{
			NationalID:  "1101700000011",
			FirstNameTH: "จอห์น",
			LastNameTH:  "โด",
		},
	}

	repo.
		On(
			"SearchPatients",
			mock.Anything,
			mock.MatchedBy(func(f patient.PatientSearchFilter) bool {
				assert.Equal(t, expectedFilter.NationalID, f.NationalID)
				assert.Equal(t, expectedFilter.FirstName, f.FirstName)
				assert.Equal(t, expectedFilter.LastName, f.LastName)
				assert.NotNil(t, f.DateOfBirth)
				assert.True(t, f.DateOfBirth.Equal(*parsedDOB))
				return true
			}),
		).
		Return(expectedResult, nil)

	svc := patientSvc.NewPatientService(repo)

	result, err := svc.SearchPatients(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	repo.AssertExpectations(t)
}
