package main

import (
	"his/config"
	"his/internal/controllers"
	"his/internal/controllers/patient"
	"his/internal/infra/postgres"
	patientsv "his/internal/services/patient/implement"

	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

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

	patientRepo := postgres.NewPatientRepository(db)

	patientService := patientsv.NewPatientService(patientRepo)

	patientController := patient.NewPatientHandler(patientService)

	router := controllers.NewRouter(controllers.Handlers{
		PatientHandler: patientController,
	})

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
