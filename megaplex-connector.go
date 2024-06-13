package main

import (
	"encoding/json"
	"log"
	. "loganfilms/types"
	"net/http"
)

func getMegaplexTheaterScheduledMovies(theaterId string) MegaplexScheduledMovies {
	url := "https://apiv2.megaplextheatres.com/api/film/cinemaFilms/" + theaterId
	res, httpErr := http.Get(url)
	if httpErr != nil {
		log.Fatal(httpErr, "Failed to get showtimes from Megaplex")
	}
	defer res.Body.Close()
	var megaplexTheaterSessions MegaplexScheduledMovies
	decodeErr := json.NewDecoder(res.Body).Decode(&megaplexTheaterSessions)
	if decodeErr != nil {
		log.Fatal(decodeErr, "Failed to decode showtimes from Megaplex")
	}
	return megaplexTheaterSessions
}
