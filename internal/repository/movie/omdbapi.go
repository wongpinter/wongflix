package movie

import (
	"context"
	"net/http"
	"time"

	"github.com/wongpinter/wongflix/internal/entity/movie"
	"github.com/wongpinter/wongflix/internal/infrastructure/datasource"
)

var (
	client = &http.Client{}
	api    = datasource.MovieApi(client, 30*time.Second)
)

type repository struct{}

func NewOmdbAPIRepository() movie.Repository {
	return &repository{}
}

func (*repository) Search(ctx context.Context, query string, page int) (*movie.SearchResult, error) {
	return api.Search(ctx, query, page)
}

func (*repository) GetByID(ctx context.Context, id string) (*movie.Movie, error) {
	return api.GetByID(ctx, id)
}

func (*repository) GetByTitle(ctx context.Context, title string) (*movie.Movie, error) {
	return api.GetByTitle(ctx, title)
}
