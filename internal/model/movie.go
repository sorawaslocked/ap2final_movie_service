package model

import "time"

type Movie struct {
	ID               string
	Rating           string // Возрастной рейтинг (К,БА, Б14 и т.д.)
	PrimaryTitle     string
	OriginalTitle    string
	ReleaseYear      uint16
	RuntimeInMinutes uint16
	Genres           []string
	InTheatricalRun  bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type MovieFilter struct {
	ID                    *string
	Ratings               []string
	PrimaryTitle          *string
	OriginalTitle         *string
	ReleaseYearRange      *ReleaseYearRange
	RuntimeInMinutesRange *RuntimeInMinutesRange
	Genres                []string
	InTheatricalRun       *bool
}

// ReleaseYearRange is inclusive
type ReleaseYearRange struct {
	YearFrom uint16
	YearTo   uint16
}

// RuntimeInMinutesRange is inclusive
type RuntimeInMinutesRange struct {
	RuntimeFrom uint16
	RuntimeTo   uint16
}

type MovieUpdateData struct {
	Rating           *string // Rating as in PG-13, R, etc.
	PrimaryTitle     *string
	OriginalTitle    *string
	ReleaseYear      *uint16
	RuntimeInMinutes *uint16
	Genres           []string
	InTheatricalRun  *bool
	UpdatedAt        time.Time
}
