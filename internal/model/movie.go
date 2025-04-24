package model

import "time"

type Movie struct {
	ID               string
	Rating           string // Rating as in PG-13, R, etc.
	PrimaryTitle     string
	OriginalTitle    string
	ReleaseYear      uint16
	RuntimeInMinutes uint16
	Genres           []string
	InTheatricalRun  bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
