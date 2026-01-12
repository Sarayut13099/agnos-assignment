package controllers

import (
	"his/internal/controllers/patient"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	PatientHandler *patient.PatientHandler
}

func NewRouter(h Handlers) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	patientGroup := r.Group("patient")
	patientGroup.GET("/search/:id", h.PatientHandler.GetByID)
	patientGroup.GET("/search", h.PatientHandler.SearchPatients)

	return r
}
