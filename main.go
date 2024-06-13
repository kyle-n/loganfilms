package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)
import . "loganfilms/types"

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	theaters := []Theater{
		{
			Screenings: make([]Screening, 0),
			Id:         "0008",
			Name:       "Megaplex University",
		},
		{
			Screenings: make([]Screening, 0),
			Id:         "0010",
			Name:       "Megaplex Providence",
		},
	}

	ch := make(chan []Screening, len(theaters))
	var wg sync.WaitGroup
	for i := range theaters {
		wg.Add(1)
		go func(i int) {
			getMegaplexScreenings(theaters[i].Id, ch)
			theaters[i].Screenings = <-ch
			wg.Done()
		}(i)
	}
	wg.Wait()

	homePageData := HomePageData{
		Theaters: theaters,
		Today:    time.Now(),
	}
	template, _ := template.ParseFiles("home.html")
	_ = template.Execute(w, homePageData)
}

func getMegaplexScreenings(theaterId string, ch chan []Screening) {
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

	screenings := make([]Screening, 0)
	for _, session := range megaplexTheaterSessions {
		showTime := Screening{
			Title:     session.Title,
			ShowTimes: make([]time.Time, 0),
			TmdbId:    0,
		}
		screenings = append(screenings, showTime)
	}

	ch <- screenings
}

func main() {
	http.HandleFunc("/", homepageHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
