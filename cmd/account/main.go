package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/neel4os/warg/libs/boilerplate"
	"github.com/neel4os/warg/libs/boilerplate/clients"
	"github.com/neel4os/warg/services/account"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Caller().Msg("starting account service")
	config := account.NewConfig()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.LoggerConfig.IsDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Debug().Msg("Debug mode for logger : Activated")
	dependencyMaps := make(map[string]boilerplate.Dependent)
	dbclient := clients.NewPgClient(config.GetDsn())
	dependencyMaps["postgres"] = dbclient
	accountService := boilerplate.NewService("account", &config, dependencyMaps, account.Routes)
	accountService.Initialize()
	// migration
	log.Debug().Caller().Msg("initiating migration")
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)
	accountService.Run()
	received := <-exit
	log.Debug().Caller().Msg("received " + received.String() + "... terminating")
	accountService.Close()
}
