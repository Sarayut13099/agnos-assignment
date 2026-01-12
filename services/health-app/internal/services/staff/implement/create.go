package implement

import (
	"context"
	"health-app/internal/domain/staff"
	"log"
)

func (s *StaffService) Create(ctx context.Context, req *staff.CreateStaffRequest) error {
	log.Println("Password:" + req.Password + ";")
	req.Password = s.encrypt.HashPassword(req.Password)

	staff := &staff.Staff{
		Username:     req.Username,
		PasswordHash: req.Password,
		HCode:        req.HospitalCode,
	}

	return s.repo.Create(ctx, staff)
}
