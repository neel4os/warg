package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/neel4os/warg/libs/boilerplate"
	"github.com/neel4os/warg/libs/boilerplate/clients"
	"github.com/neel4os/warg/services/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("starting storage service")
	config := storage.NewConfig()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.LoggerConfig.IsDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Debug().Msg("Debug mode for logger : Activated")
	dependencyMaps := make(map[string]boilerplate.Dependent)
	dbclient := clients.NewPgClient(config.GetDsn())
	dependencyMaps["postgres"] = dbclient
	dependencyMaps["objectstorage"] = storage.NewS3Client(config.ObjectStorageConfig)
	storageService := boilerplate.NewService("storage", &config, dependencyMaps, storage.Routes)
	storageService.Initialize()
	// do the migration here
	log.Debug().Caller().Msg("initiating migration")
	err := dbclient.Dbcon.AutoMigrate(&storage.Storage{}, &storage.InternalEvent{})
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("did not able to do migration")
	}
	log.Debug().Caller().Msg("migration successful")
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)
	storageService.Run()
	received := <-exit
	log.Debug().Caller().Msg("received " + received.String() + "... terminating")
	storageService.Close()
}
