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

// Mutates screenings to include just the next 10 shows, or just shows in the next 2 days if there are more than 10
func limitScreeningsToMaxOrTimeLimit(screenings *[]Screening) {
	maxShowTimes := 10
	maxTimeInFuture := time.Now().AddDate(0, 0, 2)

	for i := 0; i < len(*screenings); i++ {
		numberOfShowTimesForScreening := len((*screenings)[i].ShowTimes)
		if numberOfShowTimesForScreening < maxShowTimes {
			continue
		}

		for s := 0; s < numberOfShowTimesForScreening; s++ {
			if (*screenings)[i].ShowTimes[s].After(maxTimeInFuture) {
				(*screenings)[i].ShowTimes = (*screenings)[i].ShowTimes[:s]
				break
			}
		}
	}
}
