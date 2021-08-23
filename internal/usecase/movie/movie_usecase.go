package usecase

import (
	"context"

	"github.com/wongpinter/wongflix/internal/entity/movie"
)

type movieUseCase struct {
	mr movie.Repository
}

func NewMovieUseCase(movieRepo movie.Repository) movie.UseCase {
	return &movieUseCase{
		mr: movieRepo,
	}
}

func (movieUC *movieUseCase) Search(ctx context.Context, query string, page int) (*movie.SearchResult, error) {
	return movieUC.mr.Search(ctx, query, page)
}

func (movieUC *movieUseCase) GetByID(ctx context.Context, id string) (*movie.Movie, error) {
	return movieUC.mr.GetByID(ctx, id)
}
func (movieUC *movieUseCase) GetByTitle(ctx context.Context, title string) (*movie.Movie, error) {
	return movieUC.mr.GetByTitle(ctx, title)
}
