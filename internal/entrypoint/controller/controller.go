package covidapp

import (
	"net/http"
	"os"
	"time"

	covidapp "github.com/Test/CovidApp/internal/component"
	"github.com/gocarina/gocsv"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	CovidAppService covidapp.Service
}

// NewController creates a new CovidAPP Controller.
func NewController(covidappService covidapp.Service) *Controller {
	return &Controller{covidappService}
}

// RegisterHandlers sets up the routing for the CovidAPP Controller
func (c *Controller) RegisterHandlers(router *mux.Router) {

	covidAppService := c.CovidAppService
	// Parses csv and inserts into covid_obervations trable
	ParseCSV(covidAppService)

	get := router.
		PathPrefix("/").
		Methods(http.MethodGet, http.MethodOptions).
		Subrouter()

	get.
		Path("/top/confirmed").
		HandlerFunc(getConfirmedHandler(covidAppService))

}

func ParseCSV(service covidapp.Service) {

	type CovidObervations struct {
		SerialNo        string `csv:"SNo"`
		ObservationDate string `csv:"ObservationDate"` // .csv column headers
		State           string `csv:"Province/State"`
		Country         string `csv:"Country/Region"`
		LastUpdate      string `csv:"Last Update"`
		Confirmed       int    `csv:"Confirmed"`
		Deaths          int    `csv:"Deaths"`
		Recovered       int    `csv:"Recovered"`
	}

	in, err := os.Open("/tmp/config/covid_19_data.csv")
	if err != nil {
		logrus.Error(err)
	}

	defer in.Close()

	cObs := []CovidObervations{}
	if err := gocsv.UnmarshalFile(in, &cObs); err != nil {
		logrus.Error(err)
	}

	observationDates := []*covidapp.CovidObservations{}
	for _, obs := range cObs {
		layOut := "01/02/2006"

		parsedObservationDate, err := time.Parse(layOut, obs.ObservationDate)
		if err != nil {
			logrus.Error("Failed to parse observation date")
		}

		observationDates = append(observationDates, &covidapp.CovidObservations{
			ID:              obs.SerialNo,
			ObservationDate: &parsedObservationDate,
			State:           obs.State,
			Country:         obs.Country,
			Confirmed:       obs.Confirmed,
			Deaths:          obs.Deaths,
			Recovered:       obs.Recovered,
		})
	}

	logrus.Infof("Observation Dates: %+v", observationDates)

	errInsert := service.InsertCovidObservations(observationDates)
	if errInsert != nil {
		logrus.WithError(errInsert).Error("Failed to insert covid observations")
	}
}
