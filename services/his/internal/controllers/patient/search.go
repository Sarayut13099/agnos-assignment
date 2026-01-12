package patient

import (
	"his/internal/domain/patient"
	"his/internal/utils/response"

	"github.com/gin-gonic/gin"
)

func (h PatientHandler) SearchPatients(c *gin.Context) {
	var req patient.PatientsSearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c.Writer, response.BadRequestError(err.Error()))
		return
	}

	patients, err := h.service.SearchPatients(c.Request.Context(), req)
	if err != nil {
		response.Error(c.Writer, response.UnprocessableEntityError(err.Error()))
		return
	}

	response.Success(c.Writer, patients)
}
