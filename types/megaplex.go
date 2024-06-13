package types

import "time"

type MegaplexScheduledMovies []MegaplexScheduledMovie

type MegaplexTime time.Time

func (t *MegaplexTime) UnmarshalJSON(b []byte) error {
	tt, err := time.Parse("\"2006-01-02T15:04:05\"", string(b))
	if err != nil {
		return err
	}
	*t = MegaplexTime(tt)
	return nil
}

type MegaplexScheduledMovie struct {
	ID                                   string             `json:"id"`
	ScheduledFilmID                      string             `json:"scheduledFilmId"`
	CinemaID                             string             `json:"cinemaId"`
	Sessions                             []MegaplexSession  `json:"sessions"`
	SessionCount                         int64              `json:"sessionCount"`
	FirstDaysSessions                    interface{}        `json:"firstDaysSessions"`
	FutureSessions                       interface{}        `json:"futureSessions"`
	HasFutureSessions                    bool               `json:"hasFutureSessions"`
	Title                                string             `json:"title"`
	TitleAlt                             string             `json:"titleAlt"`
	Distributor                          string             `json:"distributor"`
	Rating                               string             `json:"rating"`
	RatingAlt                            string             `json:"ratingAlt"`
	RatingDescription                    string             `json:"ratingDescription"`
	RatingDescriptionAlt                 string             `json:"ratingDescriptionAlt"`
	Synopsis                             string             `json:"synopsis"`
	SynopsisAlt                          string             `json:"synopsisAlt"`
	OpeningDate                          MegaplexTime       `json:"openingDate"`
	FilmHOPK                             string             `json:"filmHOPK"`
	FilmHOCode                           string             `json:"filmHOCode"`
	ShortCode                            string             `json:"shortCode"`
	RunTime                              string             `json:"runTime"`
	TrailerURL                           string             `json:"trailerUrl"`
	Cast                                 interface{}        `json:"cast"`
	DisplaySequence                      int64              `json:"displaySequence"`
	TwitterTag                           string             `json:"twitterTag"`
	HasSessionsAvailable                 bool               `json:"hasSessionsAvailable"`
	GraphicURL                           string             `json:"graphicUrl"`
	CinemaName                           MegaplexCinemaName `json:"cinemaName"`
	CinemaNameAlt                        string             `json:"cinemaNameAlt"`
	AllowTicketSales                     bool               `json:"allowTicketSales"`
	AdvertiseAdvanceBookingDate          bool               `json:"advertiseAdvanceBookingDate"`
	AdvanceBookingDate                   MegaplexTime       `json:"advanceBookingDate"`
	LoyaltyAdvanceBookingDate            interface{}        `json:"loyaltyAdvanceBookingDate"`
	HasDynamicallyPricedTicketsAvailable bool               `json:"hasDynamicallyPricedTicketsAvailable"`
	IsPlayThroughMarketingFilm           bool               `json:"isPlayThroughMarketingFilm"`
	PlayThroughFilms                     interface{}        `json:"playThroughFilms"`
	NationalOpeningDate                  MegaplexTime       `json:"nationalOpeningDate"`
	CorporateFilmID                      string             `json:"corporateFilmId"`
	EDICode                              interface{}        `json:"ediCode"`
	Genres                               []interface{}      `json:"genres"`
	GenreID                              *string            `json:"genreId"`
	GenreId2                             *string            `json:"genreId2"`
	GenreId3                             *string            `json:"genreId3"`
	GenreIDs                             []*string          `json:"genreIDs"`
	AdditionalUrls                       interface{}        `json:"additionalUrls"`
}

type MegaplexSession struct {
	ID                                   string                   `json:"id"`
	CinemaID                             string                   `json:"cinemaId"`
	ScheduledFilmID                      string                   `json:"scheduledFilmId"`
	SessionID                            string                   `json:"sessionId"`
	AreaCategoryCodes                    []string                 `json:"areaCategoryCodes"`
	Showtime                             MegaplexTime             `json:"showtime"`
	IsAllocatedSeating                   bool                     `json:"isAllocatedSeating"`
	AllowChildAdmits                     bool                     `json:"allowChildAdmits"`
	SeatsAvailable                       int64                    `json:"seatsAvailable"`
	AllowComplimentaryTickets            bool                     `json:"allowComplimentaryTickets"`
	EventID                              string                   `json:"eventId"`
	PriceGroupCode                       string                   `json:"priceGroupCode"`
	ScreenName                           string                   `json:"screenName"`
	ScreenNameAlt                        string                   `json:"screenNameAlt"`
	ScreenNumber                         int64                    `json:"screenNumber"`
	CinemaOperatorCode                   string                   `json:"cinemaOperatorCode"`
	FormatCode                           MegaplexFormat           `json:"formatCode"`
	FormatHOPK                           MegaplexFormat           `json:"formatHOPK"`
	SalesChannels                        string                   `json:"salesChannels"`
	SessionAttributesNames               []MegaplexAttributesName `json:"sessionAttributesNames"`
	ConceptAttributesNames               []MegaplexAttributesName `json:"conceptAttributesNames"`
	AllowTicketSales                     bool                     `json:"allowTicketSales"`
	HasDynamicallyPricedTicketsAvailable bool                     `json:"hasDynamicallyPricedTicketsAvailable"`
	PlayThroughID                        interface{}              `json:"playThroughId"`
	SessionBusinessDate                  MegaplexTime             `json:"sessionBusinessDate"`
	SessionDisplayPriority               int64                    `json:"sessionDisplayPriority"`
	GroupSessionsByAttribute             bool                     `json:"groupSessionsByAttribute"`
	SoldoutStatus                        int64                    `json:"soldoutStatus"`
	TypeCode                             MegaplexTypeCode         `json:"typeCode"`
	SponsoredAuditoriumImageURL          interface{}              `json:"sponsoredAuditoriumImageUrl"`
}

type MegaplexCinemaName string

const (
	University MegaplexCinemaName = "University"
)

type MegaplexAttributesName string

const (
	Cc       MegaplexAttributesName = "CC"
	Dbox     MegaplexAttributesName = "DBOX"
	Dvs      MegaplexAttributesName = "DVS"
	Luxury   MegaplexAttributesName = "Luxury"
	OpenCapt MegaplexAttributesName = "Open Capt"
	Sensory  MegaplexAttributesName = "Sensory"
	The2D    MegaplexAttributesName = "2D"
	The3D    MegaplexAttributesName = "3D"
)

type MegaplexFormat string

const (
	The0000000001 MegaplexFormat = "0000000001"
	The0000000017 MegaplexFormat = "0000000017"
	Vs00000001    MegaplexFormat = "VS00000001"
)

const (
	Screen01 string = "Screen 01"
	Screen02 string = "Screen 02"
	Screen03 string = "Screen 03"
	Screen04 string = "Screen 04"
	Screen05 string = "Screen 05"
	Screen06 string = "Screen 06"
)

type MegaplexTypeCode string

const (
	N MegaplexTypeCode = "N"
)
