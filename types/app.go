package types

import (
	"time"
)

type Screening struct {
	Title       string
	ShowTimes   []time.Time
	TmdbId      int
	ReleaseDate time.Time
}

type Theater struct {
	Screenings []Screening
	Id         string
	Name       string
	Url        string
}

type HomePageData struct {
	Theaters []Theater
	Today    time.Time
}
