package implement_test

import (
	"context"
	"health-app/internal/domain/patient"
	"health-app/internal/external/mocks"
	patientSvc "health-app/internal/services/patient/implement"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPatientService_Search_Success(t *testing.T) {
	ctx := context.Background()
	hisMock := mocks.NewHISAPI(t)

	dob := "1990-01-01"
	req := patient.PatientsSearchRequest{
		NationalID:  "1101700000011",
		FirstName:   "จอห์น",
		LastName:    "โด",
		DateOfBirth: dob,
	}

	expectedResult := []*patient.Patient{
		{
			NationalID:  "1101700000011",
			FirstNameTH: "จอห์น",
			LastNameTH:  "โด",
		},
	}

	hisMock.On("SearchPatients", mock.Anything, "http://his.local", req).Return(expectedResult, nil)

	svc := patientSvc.NewPatientService(hisMock)

	result, err := svc.SearchPatients(ctx, "http://his.local", req)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

	hisMock.AssertExpectations(t)
}
