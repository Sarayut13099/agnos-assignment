package postgres

import (
	"context"
	"health-app/internal/domain/staff"

	"gorm.io/gorm"
)

type StaffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
	return &StaffRepository{db: db}
}

func (r *StaffRepository) Create(ctx context.Context, staff *staff.Staff) error {
	return r.db.WithContext(ctx).Create(staff).Error
}

func (r *StaffRepository) GetByUsername(ctx context.Context, username string) (*staff.Staff, error) {
	var staff staff.Staff
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&staff).Error
	if err != nil {
		return nil, err
	}
	return &staff, nil
}
