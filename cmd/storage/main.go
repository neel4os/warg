package main

import (
	"github.com/neel4os/warg/libs/boilerplate"
	"github.com/neel4os/warg/libs/boilerplate/clients"
	"github.com/neel4os/warg/services/storage"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("starting storage service")
	config := storage.NewConfig()
	dependencyMaps := make(map[string]boilerplate.Dependent)
	dependencyMaps["postgres"] = clients.NewPgClient(config.GetDsn())
	storageService := boilerplate.NewService("storage", &config, dependencyMaps, storage.Routes)
	storageService.Initialize()
	storageService.Run()

}
