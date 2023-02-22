package postgres

import (
	"fmt"
)

type (
	Datasource struct {
		Type       string
		Host       string
		Port       int
		Database   string
		Username   string
		Password   string
		SSLMode    string
		Migrations string
	}

	Migration struct {
		datasource *Datasource
	}
)

func (d *Datasource) AsDatasourceString() string {
	return fmt.Sprintf("postgres://postgres:%s@%s:%d/%s?&sslmode=%s", d.Password, d.Host, d.Port, d.Database, d.SSLMode)

}

func (d *Datasource) AsFileSource() string {
	return fmt.Sprintf("file://%s", d.Migrations)
}

func (d *Datasource) AsPQString() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=%s",
		d.Database,
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.SSLMode)
}
