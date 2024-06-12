package main

import "time"

type TheaterSessions []TheaterSession

type TheaterSession struct {
	ID                                   string        `json:"id"`
	ScheduledFilmID                      string        `json:"scheduledFilmId"`
	CinemaID                             string        `json:"cinemaId"`
	Sessions                             []Session     `json:"sessions"`
	SessionCount                         int64         `json:"sessionCount"`
	FirstDaysSessions                    interface{}   `json:"firstDaysSessions"`
	FutureSessions                       interface{}   `json:"futureSessions"`
	HasFutureSessions                    bool          `json:"hasFutureSessions"`
	Title                                string        `json:"title"`
	TitleAlt                             string        `json:"titleAlt"`
	Distributor                          string        `json:"distributor"`
	Rating                               string        `json:"rating"`
	RatingAlt                            string        `json:"ratingAlt"`
	RatingDescription                    string        `json:"ratingDescription"`
	RatingDescriptionAlt                 string        `json:"ratingDescriptionAlt"`
	Synopsis                             string        `json:"synopsis"`
	SynopsisAlt                          string        `json:"synopsisAlt"`
	OpeningDate                          time.Time     `json:"openingDate"`
	FilmHOPK                             string        `json:"filmHOPK"`
	FilmHOCode                           string        `json:"filmHOCode"`
	ShortCode                            string        `json:"shortCode"`
	RunTime                              string        `json:"runTime"`
	TrailerURL                           string        `json:"trailerUrl"`
	Cast                                 interface{}   `json:"cast"`
	DisplaySequence                      int64         `json:"displaySequence"`
	TwitterTag                           string        `json:"twitterTag"`
	HasSessionsAvailable                 bool          `json:"hasSessionsAvailable"`
	GraphicURL                           string        `json:"graphicUrl"`
	CinemaName                           CinemaName    `json:"cinemaName"`
	CinemaNameAlt                        string        `json:"cinemaNameAlt"`
	AllowTicketSales                     bool          `json:"allowTicketSales"`
	AdvertiseAdvanceBookingDate          bool          `json:"advertiseAdvanceBookingDate"`
	AdvanceBookingDate                   time.Time     `json:"advanceBookingDate"`
	LoyaltyAdvanceBookingDate            interface{}   `json:"loyaltyAdvanceBookingDate"`
	HasDynamicallyPricedTicketsAvailable bool          `json:"hasDynamicallyPricedTicketsAvailable"`
	IsPlayThroughMarketingFilm           bool          `json:"isPlayThroughMarketingFilm"`
	PlayThroughFilms                     interface{}   `json:"playThroughFilms"`
	NationalOpeningDate                  time.Time     `json:"nationalOpeningDate"`
	CorporateFilmID                      string        `json:"corporateFilmId"`
	EDICode                              interface{}   `json:"ediCode"`
	Genres                               []interface{} `json:"genres"`
	GenreID                              *string       `json:"genreId"`
	GenreId2                             *string       `json:"genreId2"`
	GenreId3                             *string       `json:"genreId3"`
	GenreIDs                             []*string     `json:"genreIDs"`
	AdditionalUrls                       interface{}   `json:"additionalUrls"`
}

type Session struct {
	ID                                   string           `json:"id"`
	CinemaID                             string           `json:"cinemaId"`
	ScheduledFilmID                      string           `json:"scheduledFilmId"`
	SessionID                            string           `json:"sessionId"`
	AreaCategoryCodes                    []string         `json:"areaCategoryCodes"`
	Showtime                             time.Time        `json:"showtime"`
	IsAllocatedSeating                   bool             `json:"isAllocatedSeating"`
	AllowChildAdmits                     bool             `json:"allowChildAdmits"`
	SeatsAvailable                       int64            `json:"seatsAvailable"`
	AllowComplimentaryTickets            bool             `json:"allowComplimentaryTickets"`
	EventID                              string           `json:"eventId"`
	PriceGroupCode                       string           `json:"priceGroupCode"`
	ScreenName                           string           `json:"screenName"`
	ScreenNameAlt                        string           `json:"screenNameAlt"`
	ScreenNumber                         int64            `json:"screenNumber"`
	CinemaOperatorCode                   string           `json:"cinemaOperatorCode"`
	FormatCode                           Format           `json:"formatCode"`
	FormatHOPK                           Format           `json:"formatHOPK"`
	SalesChannels                        string           `json:"salesChannels"`
	SessionAttributesNames               []AttributesName `json:"sessionAttributesNames"`
	ConceptAttributesNames               []AttributesName `json:"conceptAttributesNames"`
	AllowTicketSales                     bool             `json:"allowTicketSales"`
	HasDynamicallyPricedTicketsAvailable bool             `json:"hasDynamicallyPricedTicketsAvailable"`
	PlayThroughID                        interface{}      `json:"playThroughId"`
	SessionBusinessDate                  time.Time        `json:"sessionBusinessDate"`
	SessionDisplayPriority               int64            `json:"sessionDisplayPriority"`
	GroupSessionsByAttribute             bool             `json:"groupSessionsByAttribute"`
	SoldoutStatus                        int64            `json:"soldoutStatus"`
	TypeCode                             TypeCode         `json:"typeCode"`
	SponsoredAuditoriumImageURL          interface{}      `json:"sponsoredAuditoriumImageUrl"`
}

type CinemaName string

const (
	University CinemaName = "University"
)

type AttributesName string

const (
	Cc       AttributesName = "CC"
	Dbox     AttributesName = "DBOX"
	Dvs      AttributesName = "DVS"
	Luxury   AttributesName = "Luxury"
	OpenCapt AttributesName = "Open Capt"
	Sensory  AttributesName = "Sensory"
	The2D    AttributesName = "2D"
	The3D    AttributesName = "3D"
)

type Format string

const (
	The0000000001 Format = "0000000001"
	The0000000017 Format = "0000000017"
	Vs00000001    Format = "VS00000001"
)

const (
	Screen01 string = "Screen 01"
	Screen02 string = "Screen 02"
	Screen03 string = "Screen 03"
	Screen04 string = "Screen 04"
	Screen05 string = "Screen 05"
	Screen06 string = "Screen 06"
)

type TypeCode string

const (
	N TypeCode = "N"
)
