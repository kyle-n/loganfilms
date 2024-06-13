package main

import (
	"time"
)

type Screening struct {
	Title     string
	ShowTimes []time.Time
	TmdbId    int
}

type Theater struct {
	Screenings []Screening
	Id         string
	Name       string
}

type HomePageData struct {
	Theaters []Theater
	Today             time.Time
}
