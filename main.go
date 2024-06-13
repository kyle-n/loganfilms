package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"html/template"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	// getMegaplexShowtimes("https://www.megaplextheatres.com/university")
	universityShowTimes := getMegaplexShowtimes("0008")
	homePageData := HomePageData{
		Showtimes: universityShowTimes,
		Today: time.Now(),
	}
	template, _ := template.ParseFiles("home.html")
	_ = template.Execute(w, homePageData)
}

func getMegaplexShowtimes(theaterId string) []Showtime {
	url := "https://apiv2.megaplextheatres.com/api/film/cinemaFilms/" + theaterId
	res, httpErr := http.Get(url)
	if httpErr != nil {
		log.Fatal(httpErr, "Failed to get showtimes from Megaplex")
	}
	defer res.Body.Close()
	var megaplexTheaterSessions TheaterSessions
	decodeErr := json.NewDecoder(res.Body).Decode(&megaplexTheaterSessions)
	fmt.Println(res.Status, "Status")
	if decodeErr != nil {
		log.Fatal(decodeErr, "Failed to decode showtimes from Megaplex")
	}

	showTimes := make([]Showtime, 0)
	for _, session := range megaplexTheaterSessions {
		showTime := Showtime{
			Title:     session.Title,
			ShowTimes: make([]time.Time, 0),
			TMDBID:    0,
		}
		showTimes = append(showTimes, showTime)
	}
	fmt.Println(showTimes)

	return showTimes
}

func main() {
	http.HandleFunc("/", homepageHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
