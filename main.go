package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	universityScreenings := getMegaplexScreenings("0008")
	homePageData := HomePageData{
		Screenings: universityScreenings,
		Today:     time.Now(),
	}
	template, _ := template.ParseFiles("home.html")
	_ = template.Execute(w, homePageData)
}

func getMegaplexScreenings(theaterId string) []Screening {
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

	screenings := make([]Screening, 0)
	for _, session := range megaplexTheaterSessions {
		showTime := Screening{
			Title:     session.Title,
			ShowTimes: make([]time.Time, 0),
			TmdbId:    0,
		}
		screenings = append(screenings, showTime)
	}
	fmt.Println(screenings)

	return screenings
}

func main() {
	http.HandleFunc("/", homepageHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
