package staff

import (
	"context"
	"health-app/internal/domain/staff"
)

type Services interface {
	Create(ctx context.Context, req *staff.CreateStaffRequest) error
	Login(ctx context.Context, username, password, hospital_code string) (string, error)
}
