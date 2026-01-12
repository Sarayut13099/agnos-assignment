package implement_test

import (
	"context"
	"health-app/internal/domain/hospital"
	"health-app/internal/domain/hospital/mocks"
	"testing"

	hospitalSvc "health-app/internal/services/hospital/implement"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHospitalService_GetByHcode_Success(t *testing.T) {
	ctx := context.Background()
	repo := new(mocks.Repository)

	expected := &hospital.Hospital{
		HCode: "11001",
		Name:  "โรงพยาบาลตัวอย่าง",
	}

	repo.
		On("GetByHCode", mock.Anything, "11001").
		Return(expected, nil)
	svc := hospitalSvc.NewHospitalService(repo)

	result, err := svc.GetByHCode(ctx, "11001")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	repo.AssertExpectations(t)
}
