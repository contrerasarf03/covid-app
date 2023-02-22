package covidapp

import (
	"net/http"
	"strconv"
	"time"

	covidapp "github.com/Test/CovidApp/internal/component"
	"github.com/sirupsen/logrus"
)

func getConfirmedHandler(covidAppService covidapp.Service) http.HandlerFunc {
	logrus.Info("handler registered: GET api/v1/top/confirmed")

	return func(w http.ResponseWriter, req *http.Request) {

		query := req.URL.Query()

		// Set default limit to 10
		maxResults := 10
		// Sets default to time.Now()
		date := time.Now().Format("2006-01-02")
		if limit, ok := query["max_results"]; ok {
			intLimit, err := strconv.ParseInt(limit[0], 10, 64)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, "Parameter `limit` is invalid")
				return
			}
			maxResults = int(intLimit)
		}

		if observationDate, ok := query["observation_date"]; ok {
			date = observationDate[0]
		}

		logrus.Infof("GET TOP CONFIRM REQUEST: DATE = %+v MAX_RESULTS: %+v", date, maxResults)

		resp, httpErr := covidAppService.GetTopCofirmed(date, maxResults)
		if httpErr != nil {
			respondWithError(w, http.StatusInternalServerError, httpErr.Error())
			return
		}

		respondWithJSON(w, 200, resp)
	}
}
