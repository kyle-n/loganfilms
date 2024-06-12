package main

import (
	"encoding/json"
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
	url := "https://apiv2.megaplextheatres.com/api/film/cinemaFilms/" + theaterId
	res, httpErr := http.Get(url)
	if httpErr != nil {
		log.Fatal(httpErr, "Failed to get showtimes from Megaplex")
	}
	defer res.Body.Close()
	var showTimes TheaterSessions
	decodeErr := json.NewDecoder(res.Body).Decode(&showTimes)
	fmt.Println(res.Status, "Status")
	if decodeErr != nil {
		log.Fatal(decodeErr, "Failed to decode showtimes from Megaplex")
	}

	for _, showTime := range showTimes {
		fmt.Println(showTime.Title)
	}

	return make([]Showtime, 0)
}

func main() {
	http.HandleFunc("/", homepageHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
