package hospital

import (
	"context"
	"health-app/internal/domain/hospital"
)

type Services interface {
	GetByHCode(ctx context.Context, hcode string) (*hospital.Hospital, error)
}
