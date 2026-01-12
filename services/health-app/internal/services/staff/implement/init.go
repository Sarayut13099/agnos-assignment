package implement

import (
	"health-app/internal/domain/staff"
	"health-app/internal/utils"
)

type StaffService struct {
	repo      staff.Repository
	encrypt   utils.Encrypt
	jwtSecret []byte
	jwtTTL    int64
}

func NewStaffService(repo staff.Repository, encrypt utils.Encrypt, secret string, jwtTTL int64) *StaffService {
	return &StaffService{
		repo:      repo,
		encrypt:   encrypt,
		jwtSecret: []byte(secret),
		jwtTTL:    jwtTTL,
	}
}
