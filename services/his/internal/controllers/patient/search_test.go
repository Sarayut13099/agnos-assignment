package patient_test

import (
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

func TestSearchPatients_Success(t *testing.T) {
	mockSvc := new(mocks.Service)

	mockSvc.
		On(
			"SearchPatients",
			mock.Anything,
			mock.MatchedBy(func(req patient.PatientsSearchRequest) bool {
				return req.FirstName == "จอห์น"
			}),
		).
		Return([]*patient.Patient{
			{
				NationalID:  "1101700000011",
				FirstNameTH: "จอห์น โด",
			},
			{
				NationalID:  "1101700000012",
				FirstNameTH: "จอห์น สมิธ",
			},
		}, nil)

	h := patientHl.NewPatientHandler(mockSvc)

	r := gin.New()
	r.GET("/patient/search", h.SearchPatients)

	req := httptest.NewRequest(
		http.MethodGet,
		"/patient/search?first_name=จอห์น",
		nil,
	)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestSearchPatients_UnprocessableEntity(t *testing.T) {
	mockSvc := new(mocks.Service)
	mockSvc.
		On(
			"SearchPatients",
			mock.Anything,
			mock.Anything,
		).
		Return(nil, assert.AnError)
	h := patientHl.NewPatientHandler(mockSvc)

	r := gin.New()
	r.GET("/patient/search", h.SearchPatients)
	req := httptest.NewRequest(
		http.MethodGet,
		"/patient/search?first_name=จอห์น",
		nil,
	)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	mockSvc.AssertExpectations(t)
}
