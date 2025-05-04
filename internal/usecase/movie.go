package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_movie_service/internal/model"
	"time"
)

type MovieUseCase struct {
	repo MovieRepository
}

func NewMovie(repo MovieRepository) *MovieUseCase {
	return &MovieUseCase{repo: repo}
}

func (uc *MovieUseCase) Create(ctx context.Context, movie model.Movie) (model.Movie, error) {
	movie.CreatedAt = time.Now().UTC()
	movie.UpdatedAt = time.Now().UTC()

	newMovie, err := uc.repo.InsertOne(ctx, movie)
	if err != nil {
		return model.Movie{}, err
	}

	return newMovie, nil
}

func (uc *MovieUseCase) GetAll(ctx context.Context) ([]model.Movie, error) {
	movies, err := uc.repo.Find(ctx, model.MovieFilter{})
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (uc *MovieUseCase) GetAllWithFilter(ctx context.Context, filter model.MovieFilter) ([]model.Movie, error) {
	movies, err := uc.repo.Find(ctx, filter)
	if err != nil {
		return []model.Movie{}, err
	}

	return movies, nil
}

func (uc *MovieUseCase) GetByID(ctx context.Context, id string) (model.Movie, error) {
	movie, err := uc.repo.FindOne(ctx, model.MovieFilter{ID: &id})
	if err != nil {
		return model.Movie{}, err
	}

	return movie, nil
}

func (uc *MovieUseCase) UpdateByID(ctx context.Context, id string, update model.MovieUpdateData) (model.Movie, error) {
	update.UpdatedAt = time.Now().UTC()

	updatedMovie, err := uc.repo.UpdateOne(
		ctx,
		model.MovieFilter{ID: &id},
		update,
	)
	if err != nil {
		return model.Movie{}, err
	}

	return updatedMovie, nil
}

func (uc *MovieUseCase) DeleteByID(ctx context.Context, id string) (model.Movie, error) {
	deletedMovie, err := uc.repo.DeleteOne(ctx, model.MovieFilter{ID: &id})
	if err != nil {
		return model.Movie{}, err
	}

	return deletedMovie, nil
}
