package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/common/server/controller"
	"github.com/neel4os/warg/migration"
	"github.com/rs/zerolog/log"
)

func StartServer(cfg *config.Config) {
	// first we should do migration
	// we check db first
	dbcon := database.GetDataConn(*cfg)
	dbcon.Ping()
	migration.DoMigration(cfg)

	ctrlr := controller.NewController(cfg, dbcon)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT)
	ctrlr.Init()
	ctrlr.Run()
	v := <-exit
	log.Info().Str("signal", v.String()).Caller().Msg("Received signal")
	ctrlr.Stop()
}
