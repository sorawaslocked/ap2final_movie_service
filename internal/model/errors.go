package model

import "errors"

var (
	ErrNotFound      = errors.New("movie not found")
	ErrInvalidRating = errors.New("invalid rating")
)
