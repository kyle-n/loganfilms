package main

import (
	. "loganfilms/types"
	"slices"
	"time"
)

func sortScreeningsByReleaseDate(screenings *[]Screening) {
	slices.SortStableFunc(*screenings, func(a, b Screening) int {
		if a.ReleaseDate.Before(b.ReleaseDate) {
			return -1
		} else {
			return 1
		}
	})
}

func getShowTimesFromMegaplexSession(movie TheaterSession) []time.Time {
	showTimes := make([]time.Time, 0)
	for _, session := range movie.Sessions {
		showTimes = append(showTimes, time.Time(session.Showtime))
	}
	return showTimes
}
