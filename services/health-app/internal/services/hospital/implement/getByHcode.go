package implement

import (
	"context"
	"health-app/internal/domain/hospital"
)

func (h *hospitalService) GetByHCode(ctx context.Context, hcode string) (*hospital.Hospital, error) {
	return h.repo.GetByHCode(ctx, hcode)
}
