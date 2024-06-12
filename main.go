package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Showtime struct {
	Title     string
	ShowTimes []time.Time
	TMDBID    int
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	// getMegaplexShowtimes("https://www.megaplextheatres.com/university")
	getMegaplexShowtimes("0008")
	w.Write([]byte("Hello, World!"))
}

func getMegaplexShowtimes(theaterId string) []Showtime {
	const url = "https://apiv2.megaplextheatres.com/api/film/cinemaFilms" + theaterId
	showtimes, err := http.Get(url)

	return showTimes
}

func main() {
	http.HandleFunc("/", homepageHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
