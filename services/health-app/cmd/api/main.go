package main

import (
	"health-app/config"
	"health-app/internal/controllers"
	"health-app/internal/controllers/patient"
	"health-app/internal/controllers/staff"
	"health-app/internal/external"
	"health-app/internal/infra/postgres"
	hospitalsv "health-app/internal/services/hospital/implement"
	patientsv "health-app/internal/services/patient/implement"
	staffsv "health-app/internal/services/staff/implement"
	"health-app/internal/utils"

	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	encrypt := utils.NewEncrypt()

	db, err := postgres.NewPostgres(postgres.PostgresConfig{
		Host:     cfg.PostgresHost,
		Port:     cfg.PostgresPort,
		User:     cfg.PostgresUser,
		Password: cfg.PostgresPassword,
		DBName:   cfg.PostgresDBName,
		SSLMode:  cfg.PostgresSSLMode,
		TimeZone: cfg.PostgresTimeZone,
	})
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	staffRepo := postgres.NewStaffRepository(db)
	hospitalRepo := postgres.NewHospitalRepository(db)

	hisApiClient := external.NewHISAPIClient()

	staffService := staffsv.NewStaffService(staffRepo, encrypt, cfg.AccessTokenSigningKey, cfg.AccessTokenTTL)
	hospitalService := hospitalsv.NewHospitalService(hospitalRepo)
	patientService := patientsv.NewPatientService(hisApiClient)

	staffController := staff.NewStaffHandler(staffService)
	patientController := patient.NewPatientHandler(patientService, hospitalService)

	router := controllers.NewRouter(controllers.Handlers{
		AppConfig:      cfg,
		StaffHandler:   staffController,
		PatientHandler: patientController,
	})

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
