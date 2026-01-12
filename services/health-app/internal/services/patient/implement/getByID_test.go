package implement_test

import (
	"context"
	"health-app/internal/domain/patient"
	"health-app/internal/external/mocks"
	"testing"

	patientSvc "health-app/internal/services/patient/implement"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPatientService_GetByID_Success(t *testing.T) {
	ctx := context.Background()

	hisMock := mocks.NewHISAPI(t)

	expected := &patient.Patient{
		NationalID:  "1101700000011",
		FirstNameTH: "จอห์น",
		LastNameTH:  "โด",
	}

	hisMock.On("GetPatientByID", mock.Anything, "http://his.local", "1101700000011").Return(expected, nil)

	svc := patientSvc.NewPatientService(hisMock)

	result, err := svc.GetByID(
		ctx,
		"http://his.local",
		"1101700000011",
	)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	hisMock.AssertExpectations(t)
}
