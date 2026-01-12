package staff

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, staff *Staff) error
	GetByUsername(ctx context.Context, username string) (*Staff, error)
}
