package postgres

import (
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"

	// load the postgres migration driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	// load the file source
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

type migrationLogger struct{}

func (logger migrationLogger) Printf(format string, v ...interface{}) {
	log.Infof(format, v...)
}

func (logger migrationLogger) Verbose() bool {
	return false
}

// NewMigration ...
func NewMigration(datasource *Datasource) *Migration {
	return &Migration{datasource}
}

// Run ...
func (m *Migration) Run() error {
	source := m.datasource.AsFileSource()
	datasource := m.datasource.AsDatasourceString()

	log.Info("Running Database Migrations: ", map[string]interface{}{
		"source":     source,
		"datasource": datasource,
	})

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", m.datasource.Username, m.datasource.Password, m.datasource.Host, m.datasource.Port, m.datasource.Database)
	migrations, err := migrate.New("file://migrations", dbURL) // file://migrations | file:///var/www/migrations

	if err != nil {
		return err
	}

	if err = migrations.Up(); err != nil {
		log.Printf("UP_ERROR: %v", err.Error())
		if err.Error() != "no change" {
			return err
		}
	}
	return nil
}
