package clients

import (
	"database/sql"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgClient struct {
	Dbcon  *gorm.DB
	sqlCon *sql.DB
}

func NewPgClient(dsn string) *PgClient {
	dbcon, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError:  true,
		CreateBatchSize: 1000,
	})
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("unable to create database connection")
	}
	sqlcon, err := dbcon.DB()
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("unable to extract sql connection")
	}
	sqlcon.SetConnMaxIdleTime(3)
	return &PgClient{
		Dbcon:  dbcon,
		sqlCon: sqlcon,
	}
}

func (p *PgClient) Ping() {
	err := p.sqlCon.Ping()
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("failed to ping")
	}
}

func (p *PgClient) Close() error {
	return p.sqlCon.Close()
}

