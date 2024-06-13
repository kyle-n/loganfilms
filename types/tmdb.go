package types

type TmdbSearchResponse struct {
	Page         int64    `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int64    `json:"total_pages"`
	TotalResults int64    `json:"total_results"`
}

type Result struct {
	Adult            bool             `json:"adult"`
	BackdropPath     *string          `json:"backdrop_path"`
	GenreIDs         []int64          `json:"genre_ids"`
	ID               int64            `json:"id"`
	OriginalLanguage OriginalLanguage `json:"original_language"`
	OriginalTitle    string           `json:"original_title"`
	Overview         string           `json:"overview"`
	Popularity       float64          `json:"popularity"`
	PosterPath       *string          `json:"poster_path"`
	ReleaseDate      string           `json:"release_date"`
	Title            string           `json:"title"`
	Video            bool             `json:"video"`
	VoteAverage      float64          `json:"vote_average"`
	VoteCount        int64            `json:"vote_count"`
}

type OriginalLanguage string

const (
	De OriginalLanguage = "de"
	En OriginalLanguage = "en"
	Ja OriginalLanguage = "ja"
)
