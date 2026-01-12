package patient_test

import (
	"context"
	"his/internal/domain/patient"
	"his/internal/domain/patient/mocks"
	"testing"

	patientSvc "his/internal/services/patient/implement"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPatientService_GetByID_Success(t *testing.T) {
	ctx := context.Background()
	repo := new(mocks.Repository)

	expected := &patient.Patient{
		NationalID:  "1101700000011",
		FirstNameTH: "จอห์น",
		LastNameTH:  "โด",
	}

	repo.
		On("GetByID", mock.Anything, "1101700000011").
		Return(expected, nil)
	svc := patientSvc.NewPatientService(repo)

	result, err := svc.GetByID(ctx, "1101700000011")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	repo.AssertExpectations(t)
}
