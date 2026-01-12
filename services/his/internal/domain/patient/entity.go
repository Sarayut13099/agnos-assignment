package patient

import "time"

type Patient struct {
	FirstNameTH  string    `gorm:"column:first_name_th;not null" json:"first_name_th"`
	MiddleNameTH *string   `gorm:"column:middle_name_th" json:"middle_name_th"`
	LastNameTH   string    `gorm:"column:last_name_th;not null" json:"last_name_th"`
	FirstNameEN  *string   `gorm:"column:first_name_en" json:"first_name_en"`
	MiddleNameEN *string   `gorm:"column:middle_name_en" json:"middle_name_en"`
	LastNameEN   *string   `gorm:"column:last_name_en" json:"last_name_en"`
	DateOfBirth  time.Time `gorm:"column:date_of_birth;not null" json:"date_of_birth"`
	PatientHN    string    `gorm:"column:patient_hn;not null" json:"patient_hn"`
	NationalID   string    `gorm:"column:national_id;not null" json:"national_id"`
	PassportID   *string   `gorm:"column:passport_id" json:"passport_id"`
	PhoneNumber  *string   `gorm:"column:phone_number" json:"phone_number"`
	Email        *string   `gorm:"column:email" json:"email"`
	Gender       string    `gorm:"column:gender;not null" json:"gender"`
}

func (Patient) TableName() string {
	return "patient"
}

type PatientSearchFilter struct {
	NationalID  string
	PassportID  string
	FirstName   string
	MiddleName  string
	LastName    string
	DateOfBirth *time.Time
	PhoneNumber string
	Email       string
}

type PatientsSearchRequest struct {
	NationalID  string `form:"national_id"`
	PassportID  string `form:"passport_id"`
	FirstName   string `form:"first_name"`
	MiddleName  string `form:"middle_name"`
	LastName    string `form:"last_name"`
	DateOfBirth string `form:"date_of_birth"`
	PhoneNumber string `form:"phone_number"`
	Email       string `form:"email"`
}
