package main

import (
	"fmt"
	"html/template"
	"log"
	. "loganfilms/types"
	"net/http"
	"sync"
	"time"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	theaters := []Theater{
		{
			Screenings: make([]Screening, 0),
			Id:         "0010",
			Name:       "Megaplex Providence",
		},
		{
			Screenings: make([]Screening, 0),
			Id:         "0008",
			Name:       "Megaplex University",
		},
	}

	var wg sync.WaitGroup
	for i := range theaters {
		wg.Add(1)
		go func(i int) {
			scheduledMovies := getMegaplexTheaterScheduledMovies(theaters[i].Id)
			screenings := make([]Screening, 0)
			for _, scheduledMovie := range scheduledMovies {
				showTime := Screening{
					Title:       scheduledMovie.Title,
					ShowTimes:   getShowTimesFromMegaplexSession(scheduledMovie),
					TmdbId:      0,
					ReleaseDate: time.Time(scheduledMovie.NationalOpeningDate),
				}
				screenings = append(screenings, showTime)
			}
			theaters[i].Screenings = screenings
			sortScreeningsByReleaseDate(&theaters[i].Screenings)
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

func main() {
	http.HandleFunc("/", homepageHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
