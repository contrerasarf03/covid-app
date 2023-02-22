package covidapp

import "time"

type GetTopConfirmedResponse struct {
	ObservationDate string           `json:"observation_date"`
	Countries       []CountriesModel `json:"countries"`
}

type CountriesModel struct {
	Country   string `json:"country"`
	Confirmed int    `json:"confirmed"`
	Deaths    int    `json:"deaths"`
	Recovered int    `json:"recovered"`
}

type CovidObservations struct {
	ID              string     `json:"primary_key"`
	ObservationDate *time.Time `json:"observation_date"`
	State           string     `json:"state"`
	Country         string     `json:"country"`
	Confirmed       int        `json:"confirmed"`
	Deaths          int        `json:"deaths"`
	Recovered       int        `json:"recovered"`
}
