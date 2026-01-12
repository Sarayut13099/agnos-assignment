package patient_test

import (
	"errors"
	patientHl "his/internal/controllers/patient"
	"his/internal/domain/patient"
	"his/internal/services/patient/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPatientByID_Success(t *testing.T) {
	mockSvc := new(mocks.Service)

	mockSvc.
		On("GetByID", mock.Anything, "1101700000011").
		Return(&patient.Patient{
			NationalID: "1101700000011",
		}, nil)

	h := patientHl.NewPatientHandler(mockSvc)

	r := gin.New()
	r.GET("/patient/search/:id", h.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/patient/search/1101700000011", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetPatientByID_NotFound(t *testing.T) {
	mockSvc := new(mocks.Service)

	mockSvc.
		On("GetByID", mock.Anything, "11017000000112").
		Return(nil, errors.New("record not found"))

	h := patientHl.NewPatientHandler(mockSvc)

	r := gin.New()
	r.GET("/patient/search/:id", h.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/patient/search/11017000000112", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetPatientByID_UnprocessableEntity(t *testing.T) {
	mockSvc := new(mocks.Service)
	mockSvc.
		On("GetByID", mock.Anything, "1111").
		Return(nil, errors.New("invalid id"))

	h := patientHl.NewPatientHandler(mockSvc)

	r := gin.New()
	r.GET("/patient/search/:id", h.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/patient/search/1111", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetPatientByID_BadRequest(t *testing.T) {
	mockSvc := new(mocks.Service)

	h := patientHl.NewPatientHandler(mockSvc)

	r := gin.New()
	r.GET("/patient/search/:id", h.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/patient/search/AAA", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
