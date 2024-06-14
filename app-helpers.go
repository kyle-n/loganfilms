package main

import (
	. "loganfilms/types"
	"slices"
	"time"
)

func sortScreeningsByNextShowTime(screenings *[]Screening) {
	slices.SortStableFunc(*screenings, func(a, b Screening) int {
		if a.ShowTimes == nil || len(a.ShowTimes) == 0 {
			return 1
		}
		if b.ShowTimes == nil || len(b.ShowTimes) == 0 {
			return -1
		}
		if a.ShowTimes[0].Before(b.ShowTimes[0]) {
			return -1
		} else {
			return 1
		}
	})
}

func getShowTimesFromMegaplexSession(movie MegaplexScheduledMovie) []time.Time {
	showTimes := make([]time.Time, 0)
	for _, session := range movie.Sessions {
		showTimes = append(showTimes, time.Time(session.Showtime))
	}
	return showTimes
}
