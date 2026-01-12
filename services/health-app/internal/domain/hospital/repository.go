package hospital

import "context"

type Repository interface {
	GetByHCode(ctx context.Context, hcode string) (*Hospital, error)
}
