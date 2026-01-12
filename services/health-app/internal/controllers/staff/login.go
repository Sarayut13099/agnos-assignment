package staff

import (
	"health-app/internal/utils/response"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	HospitalCode string `json:"hospital_code" binding:"required"`
}

func (t *StaffHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c.Writer, response.BadRequestError(err.Error()))
		return
	}

	token, err := t.service.Login(c.Request.Context(), req.Username, req.Password, req.HospitalCode)
	if err != nil {
		response.Error(c.Writer, response.UnauthorizedError(err.Error()))
		return
	}

	response.Authorized(c.Writer, gin.H{"token": token})
}
