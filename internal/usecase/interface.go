package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_movie_service/internal/model"
)

type MovieRepository interface {
	Create(ctx context.Context, movie model.Movie) (model.Movie, error)
	Get(ctx context.Context, filter model.MovieFilter) (model.Movie, error)
	GetAll(ctx context.Context, filter model.MovieFilter) ([]model.Movie, error)
	Update(ctx context.Context, filter model.MovieFilter, update model.MovieUpdateData) (model.Movie, error)
	Delete(ctx context.Context, filter model.MovieFilter) (model.Movie, error)
}
