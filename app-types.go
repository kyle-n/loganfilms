package main

import (
	"time"
)

type Screening struct {
	Title     string
	ShowTimes []time.Time
	TmdbId    int
}

type HomePageData struct {
	Screenings []Screening
	Today     time.Time
}
