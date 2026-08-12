package migrations

import (
	"embed"
	"errors"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/gorm"
)

// This loads the .sql files in memory of the exe file
//
//go:embed *.sql
var migrationFiles embed.FS

func Run(db *gorm.DB) error {
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}

	sourceDriver, err := iofs.New(migrationFiles, ".")
	if err != nil {
		return err
	}

	dbDriver, err := postgres.WithInstance(sqlDb, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", dbDriver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			slog.Info("database is up to date")
			return nil
		}
		return err
	}

	slog.Info("migrations applied successfully!")
	return nil
}
