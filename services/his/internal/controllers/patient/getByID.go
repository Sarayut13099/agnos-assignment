package patient

import (
	"his/internal/utils/response"

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

	patient, err := h.service.GetByID(c.Request.Context(), req.ID)
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
