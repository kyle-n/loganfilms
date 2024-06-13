package main

import (
	"time"
)

type Showtime struct {
	Title     string
	ShowTimes []time.Time
	TMDBID    int
}

type HomePageData struct {
	Showtimes []Showtime
	Today		 time.Time
}