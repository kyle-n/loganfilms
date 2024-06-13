package main

import (
	. "loganfilms/types"
	"slices"
)

func sortScreeningsByReleaseDate(screenings *[]Screening) {
	slices.SortStableFunc(*screenings, func(a, b Screening) int {
		if a.ReleaseDate.Before(b.ReleaseDate) {
			return -1
		} else {
			return 1
		}
	});
}