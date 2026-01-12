package staff

import "time"

type Staff struct {
	HCode        string    `gorm:"column:hcode;type:varchar(5);not null"`
	Username     string    `gorm:"column:username;type:varchar(100);not null"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (Staff) TableName() string {
	return "staff"
}

type CreateStaffRequest struct {
	Username     string
	Password     string
	HospitalCode string
}
