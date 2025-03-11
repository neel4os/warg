package main

import (
	"github.com/neel4os/warg/libs/boilerplate"
	"github.com/neel4os/warg/libs/boilerplate/clients"
	"github.com/neel4os/warg/services/experiment"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"

)


func main(){
	log.Info().Caller().Msg("starting experiment service")
	config := experiment.NewConfig()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.LoggerConfig.IsDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Debug().Msg("Debug mode for logger : Activated")
	dependencyMaps := make(map[string]boilerplate.Dependent)
	dbclient := clients.NewPgClient(config.GetDsn())
	dependencyMaps["postgres"] = dbclient
	experimentService := boilerplate.NewService("experiment", &config, dependencyMaps, experiment.Routes)
	experimentService.Initialize()
	// migration
	log.Debug().Caller().Msg("initiating migration")
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)
	experimentService.Run()
	received := <-exit
	log.Debug().Caller().Msg("received " + received.String() + "... terminating")
	experimentService.Close()
}

