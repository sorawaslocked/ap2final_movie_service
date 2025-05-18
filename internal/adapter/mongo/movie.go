package mongo

import (
	"context"
	"errors"
	"github.com/sorawaslocked/ap2final_movie_service/internal/adapter/mongo/dao"
	"github.com/sorawaslocked/ap2final_movie_service/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionMovies = "movies"

type Movie struct {
	col *mongo.Collection
}

func NewMovie(conn *mongo.Database) *Movie {
	collection := conn.Collection(collectionMovies)

	return &Movie{col: collection}
}

func (db *Movie) InsertOne(ctx context.Context, movie model.Movie) (model.Movie, error) {
	movieDao := dao.FromMovie(movie)

	res, err := db.col.InsertOne(ctx, movieDao)

	if err != nil {
		return model.Movie{}, mongoError("insertOne", err)
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return db.FindOne(
		ctx,
		model.MovieFilter{ID: &id},
	)
}

func (db *Movie) FindOne(ctx context.Context, filter model.MovieFilter) (model.Movie, error) {
	var movieDao dao.Movie

	err := db.col.FindOne(
		ctx,
		dao.FromMovieFilter(filter),
	).Decode(&movieDao)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Movie{}, model.ErrNotFound
		}

		return model.Movie{}, mongoError("FindOne", err)
	}

	return dao.ToMovie(movieDao), nil
}

func (db *Movie) Find(ctx context.Context, filter model.MovieFilter) ([]model.Movie, error) {
	var movieDaos []dao.Movie

	cur, err := db.col.Find(
		ctx,
		dao.FromMovieFilter(filter),
	)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return []model.Movie{}, model.ErrNotFound
		}

		return []model.Movie{}, mongoError("Find", err)
	}

	if err = cur.All(ctx, &movieDaos); err != nil {
		return []model.Movie{}, mongoError("Cursor.All", err)
	}

	movies := make([]model.Movie, len(movieDaos))

	for i := range movieDaos {
		movies[i] = dao.ToMovie(movieDaos[i])
	}

	return movies, nil
}

func (db *Movie) UpdateOne(ctx context.Context, filter model.MovieFilter, update model.MovieUpdateData) (model.Movie, error) {
	res, err := db.col.UpdateOne(
		ctx,
		dao.FromMovieFilter(filter),
		dao.FromMovieUpdateData(update),
	)

	if err != nil {
		return model.Movie{}, mongoError("UpdateOne", err)
	}

	if res.MatchedCount == 0 {
		return model.Movie{}, model.ErrNotFound
	}

	return db.FindOne(ctx, filter)
}

func (db *Movie) DeleteOne(ctx context.Context, filter model.MovieFilter) (model.Movie, error) {
	res, err := db.col.DeleteOne(
		ctx,
		dao.FromMovieFilter(filter),
	)

	if err != nil {
		return model.Movie{}, mongoError("DeleteOne", err)
	}

	if res.DeletedCount == 0 {
		return model.Movie{}, model.ErrNotFound
	}

	return db.FindOne(ctx, filter)
}
