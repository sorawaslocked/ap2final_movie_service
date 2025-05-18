package dto

import (
	"github.com/sorawaslocked/ap2final_movie_service/internal/model"
	"github.com/sorawaslocked/ap2final_protos_gen/base"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/movie"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToMovieFromCreateRequest(req *svc.CreateRequest) model.Movie {
	return model.Movie{
		AgeRating:        req.AgeRating,
		PrimaryTitle:     req.PrimaryTitle,
		OriginalTitle:    req.OriginalTitle,
		ReleaseYear:      uint16(req.ReleaseYear),
		RuntimeInMinutes: uint16(req.RuntimeInMinutes),
		Genres:           req.Genres,
	}
}

func ToMovieUpdateFromUpdateRequest(req *svc.UpdateRequest) (string, model.MovieUpdateData) {
	var releaseYear, runtimeInMinutes *uint16

	if req.ReleaseYear != nil {
		val := uint16(*req.ReleaseYear)
		releaseYear = &val
	}

	if req.RuntimeInMinutes != nil {
		val := uint16(*req.RuntimeInMinutes)
		runtimeInMinutes = &val
	}

	return req.ID, model.MovieUpdateData{
		AgeRating:        req.AgeRating,
		PrimaryTitle:     req.PrimaryTitle,
		OriginalTitle:    req.OriginalTitle,
		ReleaseYear:      releaseYear,
		RuntimeInMinutes: runtimeInMinutes,
		Genres:           req.Genres,
		IsDeleted:        req.IsDeleted,
	}
}

func FromMovieToPb(movie model.Movie) *base.Movie {
	return &base.Movie{
		ID:               movie.ID,
		AgeRating:        movie.AgeRating,
		PrimaryTitle:     movie.PrimaryTitle,
		OriginalTitle:    movie.OriginalTitle,
		ReleaseYear:      uint32(movie.ReleaseYear),
		RuntimeInMinutes: uint32(movie.RuntimeInMinutes),
		Genres:           movie.Genres,
		CreatedAt:        timestamppb.New(movie.CreatedAt),
		UpdatedAt:        timestamppb.New(movie.UpdatedAt),
		IsDeleted:        movie.IsDeleted,
	}
}
