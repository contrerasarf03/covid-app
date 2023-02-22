package covidapp

type Service interface {
	GetTopCofirmed(observationDate string, limit int) (GetTopConfirmedResponse, error)
	InsertCovidObservations(list []*CovidObservations) (err error)
}
