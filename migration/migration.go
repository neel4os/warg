package migration

import (
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/rs/zerolog/log"
)

func DoMigration(config *config.Config) {
	dbconfig := config.DbConfig
	dbdsn := dbconfig.GetDbDsn()
	log.Info().Str("dbdsn", dbdsn).Msg("dbdsn")
	migrationDir := filepath.Join(".", "migration")
	m, err := migrate.New(
		"file://"+migrationDir,
		dbdsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create migration instance")
	}
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Info().Msg("no migration needed")
			return
		}
		log.Fatal().Err(err).Msg("failed to apply migration")
	}
}
