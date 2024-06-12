package main

import (
	"log"
	"net/http"
	"time"
)

type Showtime struct {
	Title     string
	StartTime time.Time
	TMDBID    int
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", homepageHandler)
	log.Default().Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
