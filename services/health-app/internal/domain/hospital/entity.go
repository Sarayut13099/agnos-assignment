package hospital

import "time"

type Hospital struct {
	HCode      string    `gorm:"column:hcode;not null"`
	Name       string    `gorm:"column:name;not null"`
	HISBaseURL string    `gorm:"column:his_base_url;not null"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (Hospital) TableName() string {
	return "hospital"
}
