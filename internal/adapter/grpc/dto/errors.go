package dto

import (
	"github.com/sorawaslocked/ap2final_movie_service/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FromError(err error) error {
	switch err {
	case model.ErrNotFound:
		return status.Error(codes.NotFound, "movie not found")
	default:
		return status.Error(codes.Internal, "something went wrong")
	}
}
