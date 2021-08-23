package movie

import "context"

type Repository interface {
	Search(ctx context.Context, query string, page int) (*SearchResult, error)
	GetByID(ctx context.Context, id string) (*Movie, error)
	GetByTitle(ctx context.Context, title string) (*Movie, error)
}

type UseCase interface {
	Search(ctx context.Context, query string, page int) (*SearchResult, error)
	GetByID(ctx context.Context, id string) (*Movie, error)
	GetByTitle(ctx context.Context, title string) (*Movie, error)
}
