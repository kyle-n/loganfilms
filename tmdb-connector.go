package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	. "loganfilms/types"
)

func searchForMovieTmdbId(title string) int {
	const noId = -1
	req, buildReqErr := http.NewRequest("GET", "https://api.themoviedb.org/3/search/movie", nil)
	if buildReqErr != nil {
		return noId
	}

	tmdbApiKey := ""
	req.Header.Set("Authorization", "Bearer "+tmdbApiKey)
	req.Header.Set("accept", "application/json")

	currentYear := time.Now().Year()

	q := req.URL.Query()
	q.Add("query", title)
	q.Add("include_adult", "false")
	q.Add("language", "en-US")
	q.Add("primary_release_year", strconv.Itoa(currentYear))
	req.URL.RawQuery = q.Encode()

	res, httpErr := http.DefaultClient.Do(req)
	if httpErr != nil {
		return noId
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode, "tmdb status")

	var tmdbSearchResults TmdbSearchResponse
	decodeErr := json.NewDecoder(res.Body).Decode(&tmdbSearchResults)
	if decodeErr != nil {
		return noId
	}

	return noId
}
