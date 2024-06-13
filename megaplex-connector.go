package main

import (
	"encoding/json"
	"log"
	. "loganfilms/types"
	"net/http"
)

func getMegaplexTheaterSessions(theaterId string) TheaterSessions {
	url := "https://apiv2.megaplextheatres.com/api/film/cinemaFilms/" + theaterId
	res, httpErr := http.Get(url)
	if httpErr != nil {
		log.Fatal(httpErr, "Failed to get showtimes from Megaplex")
	}
	defer res.Body.Close()
	var megaplexTheaterSessions TheaterSessions
	decodeErr := json.NewDecoder(res.Body).Decode(&megaplexTheaterSessions)
	if decodeErr != nil {
		log.Fatal(decodeErr, "Failed to decode showtimes from Megaplex")
	}
	return megaplexTheaterSessions
}
