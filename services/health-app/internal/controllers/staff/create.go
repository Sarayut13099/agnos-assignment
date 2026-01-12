package staff

import (
	"health-app/internal/domain/staff"
	"health-app/internal/utils/response"

	"github.com/gin-gonic/gin"
)

type CreateStaffRequest struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	HospitalCode string `json:"hospital_code" binding:"required"`
}

func (t *StaffHandler) CreateStaff(c *gin.Context) {
	var req CreateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c.Writer, response.BadRequestError(err.Error()))
		return
	}

	err := t.service.Create(c.Request.Context(), &staff.CreateStaffRequest{
		Username:     req.Username,
		Password:     req.Password,
		HospitalCode: req.HospitalCode,
	})
	if err != nil {
		response.Error(c.Writer, response.UnprocessableEntityError(err.Error()))
		return
	}

	response.Success(c.Writer, nil)
}
