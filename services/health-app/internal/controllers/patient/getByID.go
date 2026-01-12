package patient

import (
	"health-app/internal/utils/response"

	"github.com/gin-gonic/gin"
)

type GetPatientByIDRequest struct {
	ID string `uri:"id" binding:"required,numeric"`
}

func (h PatientHandler) GetByID(c *gin.Context) {
	var req GetPatientByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
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

	patient, err := h.service.GetByID(c.Request.Context(), hospitalData.HISBaseURL, req.ID)
	if err != nil {
		if err.Error() == response.RecordNotFound {
			response.Error(c.Writer, response.NotFoundError(err.Error()))
			return
		}
		response.Error(c.Writer, response.UnprocessableEntityError(err.Error()))
		return
	}

	response.Success(c.Writer, patient)
}
