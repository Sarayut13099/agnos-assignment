package implement_test

import (
	"context"
	"health-app/internal/domain/staff"
	"health-app/internal/domain/staff/mocks"
	"testing"

	staffSvc "health-app/internal/services/staff/implement"
	mockEncrypt "health-app/internal/utils/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStaffService_Create_Success(t *testing.T) {
	ctx := context.Background()

	repo := new(mocks.Repository)
	encryptMock := new(mockEncrypt.Encrypt)

	req := &staff.CreateStaffRequest{
		Username:     "admin",
		Password:     "plain123",
		HospitalCode: "H001",
	}

	hashedPassword := "hashed-password-123"

	// mock encrypt
	encryptMock.
		On("HashPassword", "plain123").
		Return(hashedPassword).
		Once()

	// mock repo
	repo.
		On(
			"Create",
			mock.Anything,
			mock.MatchedBy(func(s *staff.Staff) bool {
				return s.Username == "admin" &&
					s.PasswordHash == hashedPassword &&
					s.HCode == "H001"
			}),
		).
		Return(nil).
		Once()

	svc := staffSvc.NewStaffService(repo, encryptMock, mock.Anything, 8)

	err := svc.Create(ctx, req)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
	encryptMock.AssertExpectations(t)
}
