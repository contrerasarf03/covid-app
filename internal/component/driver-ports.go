package covidapp

type Repository interface {
	GetTopCofirmed(observationDate string, limit int) (GetTopConfirmedResponse, error)
	InsertCovidObservations(list []*CovidObservations) (err error)
}
