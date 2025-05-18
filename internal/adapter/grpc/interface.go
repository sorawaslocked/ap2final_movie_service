package grpc

import (
	"context"
	"github.com/sorawaslocked/ap2final_movie_service/internal/model"
)

type MovieUseCase interface {
	Create(ctx context.Context, movie model.Movie) (model.Movie, error)
	GetAll(ctx context.Context) ([]model.Movie, error)
	GetAllWithFilter(ctx context.Context, filter model.MovieFilter) ([]model.Movie, error)
	GetByID(ctx context.Context, id string) (model.Movie, error)
	UpdateByID(ctx context.Context, id string, update model.MovieUpdateData) (model.Movie, error)
	DeleteByID(ctx context.Context, id string) (model.Movie, error)
}
