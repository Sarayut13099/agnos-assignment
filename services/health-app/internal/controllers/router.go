package controllers

import (
	"health-app/config"
	"health-app/internal/controllers/patient"
	"health-app/internal/controllers/staff"
	"health-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	AppConfig      *config.Config
	StaffHandler   *staff.StaffHandler
	PatientHandler *patient.PatientHandler
}

func NewRouter(h Handlers) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	println("AccessTokenSigningKey:", h.AppConfig.AccessTokenSigningKey)
	midAuth := middleware.Auth(h.AppConfig.AccessTokenSigningKey)

	staff := r.Group("/staff")
	staff.POST("", h.StaffHandler.CreateStaff)
	staff.POST("/login", h.StaffHandler.Login)

	patient := r.Group("/patient", midAuth)
	patient.GET("/search", h.PatientHandler.SearchPatients)
	patient.GET("/search/:id", h.PatientHandler.GetByID)

	return r
}
