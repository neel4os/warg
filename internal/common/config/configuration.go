package config

import (
	"os"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/neel4os/warg/internal/common/logging"
	"github.com/rs/zerolog/log"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	IdpConfig    IdpConfig
	LoggerConfig LoggerConfig
	ServerConfig ServerConfig
	DbConfig     DbConfig
}

type DbConfig struct {
	Port     string `env:"WARG_DBCONFIG_PORT" envDefault:"5433"`
	Host     string `env:"WARG_DBCONFIG_HOST" envDefault:"localhost"`
	User     string `env:"WARG_DBCONFIG_USER" envDefault:"postgres"`
	Password string `env:"WARG_DBCONFIG_PASSWORD" envDefault:"postgres123"`
	DbName   string `env:"WARG_DBCONFIG_DBNAME" envDefault:"warg"`
	SslMode  string `env:"WARG_DBCONFIG_SSLMODE" envDefault:"disable"`
}

func (dbConfig *DbConfig) GetDbDsn() string {
	return "postgres://" +
		dbConfig.User + ":" +
		dbConfig.Password + "@" +
		dbConfig.Host + ":" +
		dbConfig.Port + "/" +
		dbConfig.DbName + "?sslmode=" +
		dbConfig.SslMode
}

type ServerConfig struct {
	Port                  string `env:"WARG_SERVERCONFIG_PORT" envDefault:"9999"`
	HidePortInStdOut      bool   `env:"WARG_SERVERCONFIG_HIDE_PORT_IN_STDOUT" envDefault:"true"`
	ReadTimeout           int    `envDefault:"10"`
	WriteTimeout          int    `envDefault:"10"`
	GraceFullShutdownTime int    `envDefault:"10"`
}

type LoggerConfig struct {
	IsDebugLog bool `env:"WARG_LOGGERCONFIG_IS_DEBUG_LOG" envDefault:"true"`
}

type IdpConfig struct {
	IdpName     string `env:"WARG_IDPCONFIG_IDP_NAME" envDefault:"keycloak"`
	Url         string `env:"WARG_IDPCONFIG_IDP_URL" envDefault:"http://localhost:8080"`
	RealmName   string `env:"WARG_IDPCONFIG_IDP_REALM_NAME" envDefault:"warg"`
	Username    string `env:"WARG_IDPCONFIG_IDP_USERNAME" envDefault:"admin"`
	Password    string `env:"WARG_IDPCONFIG_IDP_PASSWORD" envDefault:"admin123"`
	TokenExpiry int    `env:"WARG_IDPCONFIG_TOKEN_EXPIRY" envDefault:"55"`
}

func new() *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Error().Err(err).Caller().Msg("failed to parse environment variables")
		os.Exit(1)
	}
	if cfg.LoggerConfig.IsDebugLog {
		logging.SetLogConfig(true)
	} else {
		logging.SetLogConfig(false)
	}
	return &cfg
}

func GetConfig() *Config {
	once.Do(func() {
		instance = new()
	})

	return instance
}
