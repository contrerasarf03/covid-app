package repository

import (
	"time"

	covidapp "github.com/Test/CovidApp/internal/component"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type GormRepository struct {
	*gorm.DB
}

type CovidObservations struct {
	ID              string     `gorm:"primary_key"`
	ObservationDate *time.Time `gorm:"observation_date"`
	State           string     `gorm:"state"`
	Country         string     `gorm:"country"`
	Confirmed       int        `gorm:"confirmed"`
	Deaths          int        `gorm:"deaths"`
	Recovered       int        `gorm:"recovered"`
	CreatedAt       *time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"updated_at" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"deleted_at" json:"deleted_at"`
}

// TableName ...
func (CovidObservations) TableName() string {
	return "covid_observations"
}

func fromCovidObservationsToGORM(cvd *covidapp.CovidObservations) CovidObservations {
	return CovidObservations{
		ID:              cvd.ID,
		ObservationDate: cvd.ObservationDate,
		State:           cvd.State,
		Country:         cvd.Country,
		Confirmed:       cvd.Confirmed,
		Deaths:          cvd.Deaths,
		Recovered:       cvd.Recovered,
	}
}

// NewGormRepository creates the service that connects to the database.
func NewGormRepository(db *gorm.DB) covidapp.Repository {
	return &GormRepository{
		db.Debug(),
	}
}

func (repo *GormRepository) GetTopCofirmed(observationDate string, limit int) (covidapp.GetTopConfirmedResponse, error) {
	var cos CovidObservations
	var topConfirmed []covidapp.CovidObservations
	err := repo.Table(cos.TableName()).Where("observation_date = ?", observationDate).Order("confirmed DESC").Limit(limit).Find(&topConfirmed).Error
	if err != nil {
		return covidapp.GetTopConfirmedResponse{}, err
	}

	var countries []covidapp.CountriesModel
	if len(topConfirmed) > 0 {
		for _, covidObservation := range topConfirmed {
			countries = append(countries, covidapp.CountriesModel{
				Country:   covidObservation.Country,
				Confirmed: covidObservation.Confirmed,
				Deaths:    covidObservation.Deaths,
				Recovered: covidObservation.Recovered,
			})
		}
	}

	return covidapp.GetTopConfirmedResponse{
		ObservationDate: observationDate,
		Countries:       countries,
	}, nil
}

func (repo *GormRepository) InsertCovidObservations(list []*covidapp.CovidObservations) (err error) {
	for _, data := range list {
		if err = repo.insertCovidObservation(data); err != nil {
			return
		}
	}
	return
}

func (repo *GormRepository) insertCovidObservation(data *covidapp.CovidObservations) (err error) {
	a := fromCovidObservationsToGORM(data)

	if err := repo.Create(&a).Error; err != nil {
		logrus.Error("Failed to insert record in `covid_observations` table: ", err)
		return err
	}

	return
}
