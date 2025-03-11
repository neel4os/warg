package experiment

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ServerOptions       ServerOptions
	PostgresConfig      PostgresConfig
	LoggerConfig        LoggerConfig
	ObjectStorageConfig ObjectStorageConfig
}

type ServerOptions struct {
	Port         int  `env:"SERVER_PORT" envDefault:"9999"`
	TLSActivated bool `env:"SERVER_HOST" envDefault:"false"`
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     int    `env:"POSTGRES_HOST" envDefault:"5433"`
	Database string `env:"POSTGRES_DB" envDefault:"storage"`
	Username string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres123"`
	SSLMode  string `env:"POSTGRES_SSLMODE" envDefault:"disable"`
}

type ObjectStorageConfig struct {
	Url string `env:"STORAGE_URL" envDefault:"http://localhost:9999"`
}

type LoggerConfig struct {
	IsDebug bool `env:"LOGGER_DEBUG" envDefault:"true"`
}

func (c *Config) GetDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.PostgresConfig.Host,
		c.PostgresConfig.Username,
		c.PostgresConfig.Password,
		c.PostgresConfig.Database,
		c.PostgresConfig.Port,
		c.PostgresConfig.SSLMode)
}

func (c Config) GetServerPort() int {
	return c.ServerOptions.Port
}

func NewConfig() Config {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("could not parse config")
	}
	return cfg
}
