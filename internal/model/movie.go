package model

import "time"

type Movie struct {
	ID               string
	AgeRating        string
	PrimaryTitle     string
	OriginalTitle    string
	ReleaseYear      uint16
	RuntimeInMinutes uint16
	Genres           []string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	IsDeleted        bool
}

type MovieFilter struct {
	ID                    *string
	AgeRating             *string
	PrimaryTitle          *string
	OriginalTitle         *string
	ReleaseYearRange      *ReleaseYearRange
	RuntimeInMinutesRange *RuntimeInMinutesRange
	Genres                []string
	IsDeleted             *bool
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
	AgeRating        *string
	PrimaryTitle     *string
	OriginalTitle    *string
	ReleaseYear      *uint16
	RuntimeInMinutes *uint16
	Genres           []string
	UpdatedAt        time.Time
	IsDeleted        *bool
}
