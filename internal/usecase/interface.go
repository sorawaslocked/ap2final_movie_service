package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_movie_service/internal/model"
)

type MovieRepository interface {
	Create(ctx context.Context, movie model.Movie) (model.Movie, error)
	FindOne(ctx context.Context, filter model.MovieFilter) (model.Movie, error)
	Find(ctx context.Context, filter model.MovieFilter) ([]model.Movie, error)
	UpdateOne(ctx context.Context, filter model.MovieFilter, update model.MovieUpdateData) (model.Movie, error)
	DeleteOne(ctx context.Context, filter model.MovieFilter) (model.Movie, error)
}
