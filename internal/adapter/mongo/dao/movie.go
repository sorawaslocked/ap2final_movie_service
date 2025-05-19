package dao

import (
	"github.com/sorawaslocked/ap2final_movie_service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Movie struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	AgeRating        string             `bson:"ageRating"`
	PrimaryTitle     string             `bson:"primaryTitle"`
	OriginalTitle    string             `bson:"originalTitle"`
	ReleaseYear      uint16             `bson:"releaseYear"`
	RuntimeInMinutes uint16             `bson:"runtimeInMinutes"`
	Genres           []string           `bson:"genres"`
	CreatedAt        time.Time          `bson:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt"`
	IsDeleted        bool               `bson:"isDeleted"`
}

func FromMovie(movie model.Movie) (Movie, error) {
	objID, err := primitive.ObjectIDFromHex(movie.ID)
	if err != nil && movie.ID != "" {
		return Movie{}, err
	}

	return Movie{
		ID:               objID,
		AgeRating:        movie.AgeRating,
		PrimaryTitle:     movie.PrimaryTitle,
		OriginalTitle:    movie.OriginalTitle,
		ReleaseYear:      movie.ReleaseYear,
		RuntimeInMinutes: movie.RuntimeInMinutes,
		Genres:           movie.Genres,
		IsDeleted:        movie.IsDeleted,
		CreatedAt:        movie.CreatedAt,
		UpdatedAt:        movie.UpdatedAt,
	}, nil
}

func ToMovie(movie Movie) model.Movie {
	return model.Movie{
		ID:               movie.ID.Hex(),
		AgeRating:        movie.AgeRating,
		PrimaryTitle:     movie.PrimaryTitle,
		OriginalTitle:    movie.OriginalTitle,
		ReleaseYear:      movie.ReleaseYear,
		RuntimeInMinutes: movie.RuntimeInMinutes,
		Genres:           movie.Genres,
		IsDeleted:        movie.IsDeleted,
		CreatedAt:        movie.CreatedAt,
		UpdatedAt:        movie.UpdatedAt,
	}
}

func FromMovieFilter(filter model.MovieFilter) (bson.M, error) {
	query := bson.M{}

	if filter.ID != nil {
		objID, err := primitive.ObjectIDFromHex(*filter.ID)
		if err != nil {
			return query, err
		}

		query["_id"] = objID
	}

	if filter.AgeRating != nil {
		query["ageRating"] = *filter.AgeRating
	}

	if filter.PrimaryTitle != nil {
		query["primaryTitle"] = *filter.PrimaryTitle
	}

	if filter.OriginalTitle != nil {
		query["originalTitle"] = *filter.OriginalTitle
	}

	if filter.ReleaseYearRange != nil {
		query["releaseYear"] = bson.M{
			"$gte": filter.ReleaseYearRange.YearFrom,
			"$lte": filter.ReleaseYearRange.YearTo,
		}
	}

	if filter.RuntimeInMinutesRange != nil {
		query["runtimeInMinutes"] = bson.M{
			"$gte": filter.RuntimeInMinutesRange.RuntimeFrom,
			"$lte": filter.RuntimeInMinutesRange.RuntimeTo,
		}
	}

	if filter.Genres != nil {
		query["genres"] = bson.M{
			"$all": filter.Genres,
		}
	}

	if filter.IsDeleted != nil {
		query["isDeleted"] = *filter.IsDeleted
	}

	return query, nil
}

func FromMovieUpdateData(update model.MovieUpdateData) bson.M {
	query := bson.M{}

	if update.AgeRating != nil {
		query["ageRating"] = *update.AgeRating
	}

	if update.PrimaryTitle != nil {
		query["primaryTitle"] = *update.PrimaryTitle
	}

	if update.OriginalTitle != nil {
		query["originalTitle"] = *update.OriginalTitle
	}

	if update.ReleaseYear != nil {
		query["releaseYear"] = *update.ReleaseYear
	}

	if update.RuntimeInMinutes != nil {
		query["runtimeInMinutes"] = *update.RuntimeInMinutes
	}

	if update.Genres != nil {
		query["genres"] = update.Genres
	}

	if update.IsDeleted != nil {
		query["isDeleted"] = *update.IsDeleted
	}

	query["updatedAt"] = update.UpdatedAt

	return bson.M{"$set": query}
}
