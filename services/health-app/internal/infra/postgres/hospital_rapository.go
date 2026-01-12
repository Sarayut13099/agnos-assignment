package postgres

import (
	"context"
	"health-app/internal/domain/hospital"

	"gorm.io/gorm"
)

type HospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) *HospitalRepository {
	return &HospitalRepository{db: db}
}

func (r *HospitalRepository) GetByHCode(ctx context.Context, hcode string) (*hospital.Hospital, error) {
	var hospital hospital.Hospital
	err := r.db.WithContext(ctx).Where("hcode = ?", hcode).First(&hospital).Error
	if err != nil {
		return nil, err
	}
	return &hospital, nil
}
