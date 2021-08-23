package membership

import "context"

type Read interface {
	GetByID(ctx context.Context, id string)
	All(ctx context.Context)
}

type Write interface {
	Store(ctx context.Context, membership *Membership) error
}

type Repository interface {
	Read
	Write
}
