package implement_test

import (
	"context"
	"errors"
	"health-app/internal/domain/staff"
	"health-app/internal/domain/staff/mocks"
	staffSvc "health-app/internal/services/staff/implement"
	mockEncrypt "health-app/internal/utils/mocks"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStaffService_Login_Success(t *testing.T) {
	ctx := context.Background()

	repo := new(mocks.Repository)
	encryptMock := new(mockEncrypt.Encrypt)

	staffEntity := &staff.Staff{
		Username:     "admin",
		PasswordHash: "hashed-password",
		HCode:        "H001",
	}

	repo.
		On("GetByUsername", mock.Anything, "admin").
		Return(staffEntity, nil).
		Once()

	encryptMock.
		On("CheckPasswordHash", "plain123", "hashed-password").
		Return(true).
		Once()

	jwtSecret := []byte("secret")

	svc := staffSvc.NewStaffService(
		repo,
		encryptMock,
		"secret",
		8,
	)

	token, err := svc.Login(ctx, "admin", "plain123", "H001")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// ตรวจ token จริง (optional แต่แนะนำ)
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	assert.NoError(t, err)
	assert.True(t, parsed.Valid)

	claims := parsed.Claims.(jwt.MapClaims)
	assert.Equal(t, "H001", claims["hospital_code"])

	repo.AssertExpectations(t)
	encryptMock.AssertExpectations(t)
}

func TestStaffService_Login_UserNotFound(t *testing.T) {
	ctx := context.Background()

	repo := new(mocks.Repository)
	encryptMock := new(mockEncrypt.Encrypt)

	repo.
		On("GetByUsername", mock.Anything, "admin").
		Return(nil, errors.New("not found")).
		Once()

	svc := staffSvc.NewStaffService(
		repo,
		encryptMock,
		"secret",
		1,
	)

	token, err := svc.Login(ctx, "admin", "123", "H001")

	assert.Error(t, err)
	assert.EqualError(t, err, "user not found")
	assert.Empty(t, token)

	repo.AssertExpectations(t)
}

func TestStaffService_Login_WrongPassword(t *testing.T) {
	ctx := context.Background()

	repo := new(mocks.Repository)
	encryptMock := new(mockEncrypt.Encrypt)

	staffEntity := &staff.Staff{
		Username:     "admin",
		PasswordHash: "hashed-password",
		HCode:        "H001",
	}

	repo.
		On("GetByUsername", mock.Anything, "admin").
		Return(staffEntity, nil).
		Once()

	encryptMock.
		On("CheckPasswordHash", "wrongpass", "hashed-password").
		Return(false).
		Once()

	svc := staffSvc.NewStaffService(
		repo,
		encryptMock,
		"secret",
		1,
	)

	token, err := svc.Login(ctx, "admin", "wrongpass", "H001")

	assert.Error(t, err)
	assert.EqualError(t, err, "wrong password")
	assert.Empty(t, token)

	repo.AssertExpectations(t)
	encryptMock.AssertExpectations(t)
}
