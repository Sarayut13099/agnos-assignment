package patient

import (
	"health-app/internal/domain/patient"
	"health-app/internal/utils/response"

	"github.com/gin-gonic/gin"
)

func (h PatientHandler) SearchPatients(c *gin.Context) {
	var req patient.PatientsSearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c.Writer, response.BadRequestError(err.Error()))
		return
	}

	hcode, ok := c.Get("hospital_code")
	if !ok {
		response.Error(c.Writer, response.UnprocessableEntityError("hospital code not found"))
		return
	}

	hospitalData, err := h.hospitalService.GetByHCode(c.Request.Context(), hcode.(string))
	if err != nil {
		if err.Error() == response.RecordNotFound {
			response.Error(c.Writer, response.NotFoundError(err.Error()))
			return
		}
		response.Error(c.Writer, response.UnprocessableEntityError(err.Error()))
		return
	}

	patients, err := h.service.SearchPatients(c.Request.Context(), hospitalData.HISBaseURL, req)
	if err != nil {
		response.Error(c.Writer, response.UnprocessableEntityError(err.Error()))
		return
	}

	response.Success(c.Writer, patients)
}
