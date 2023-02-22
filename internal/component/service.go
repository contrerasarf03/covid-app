package covidapp

// Config ...
type Config struct {
	Host string
	Port int
	Spec string
}

type CovidApp struct {
	repo   Repository
	config Config
}

func NewCovidApp(repo Repository, config Config) Service {
	service := &CovidApp{
		repo:   repo,
		config: config,
	}

	return service
}

func (s *CovidApp) GetTopCofirmed(observationDate string, limit int) (GetTopConfirmedResponse, error) {
	return s.repo.GetTopCofirmed(observationDate, limit)
}

func (s *CovidApp) InsertCovidObservations(list []*CovidObservations) (err error) {
	return s.repo.InsertCovidObservations(list)
}
