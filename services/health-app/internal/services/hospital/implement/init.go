package implement

import (
	hospitalDb "health-app/internal/domain/hospital"
	"health-app/internal/services/hospital"
)

type hospitalService struct {
	repo hospitalDb.Repository
}

func NewHospitalService(repo hospitalDb.Repository) hospital.Services {
	return &hospitalService{
		repo: repo,
	}
}
