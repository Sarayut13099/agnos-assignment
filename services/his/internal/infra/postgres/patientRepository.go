package postgres

import (
	"context"
	"his/internal/domain/patient"
	"log"

	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) GetByID(ctx context.Context, id string) (*patient.Patient, error) {
	var patient patient.Patient
	err := r.db.WithContext(ctx).Where("national_id = ? OR passport_id = ?", id, id).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *PatientRepository) SearchPatients(ctx context.Context, filter patient.PatientSearchFilter) ([]*patient.Patient, error) {
	var patients []*patient.Patient
	query := r.db.WithContext(ctx)

	if filter.NationalID != "" {
		query = query.Where("national_id = ?", filter.NationalID)
	}
	if filter.PassportID != "" {
		query = query.Where("passport_id = ?", filter.PassportID)
	}
	if filter.FirstName != "" {
		query = query.Where("first_name_th LIKE ? OR first_name_en LIKE ?", "%"+filter.FirstName+"%", "%"+filter.FirstName+"%")
	}
	if filter.MiddleName != "" {
		query = query.Where("middle_name_th LIKE ? OR middle_name_en LIKE ?", "%"+filter.MiddleName+"%", "%"+filter.MiddleName+"%")
	}
	if filter.LastName != "" {
		query = query.Where("last_name_th LIKE ? OR last_name_en LIKE ?", "%"+filter.LastName+"%", "%"+filter.LastName+"%")
	}
	if filter.DateOfBirth != nil {
		query = query.Where("date_of_birth = ?", *filter.DateOfBirth)
	}
	if filter.PhoneNumber != "" {
		query = query.Where("phone_number LIKE ?", "%"+filter.PhoneNumber+"%")
	}
	if filter.Email != "" {
		query = query.Where("email LIKE ?", "%"+filter.Email+"%")
	}

	log.Println("query:", query.Statement.SQL.String())

	err := query.Limit(20).Find(&patients).Error
	if err != nil {
		return nil, err
	}
	return patients, nil
}
