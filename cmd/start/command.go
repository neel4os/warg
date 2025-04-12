package start

import (
	"github.com/neel4os/warg/internal/common/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start warg instance",
		Long:  `Start warg instance without initializing`,
		RunE:  StartWarg,
	}
}

func StartWarg(cmd *cobra.Command, args []string) error {
	log.Info().Caller().Msg("Starting Warg instance")
	log.Info().Caller().Msg("Reading Warg configuration")
	config := config.GetConfig()
	//log.Debug().Interface("config", config).Caller().Msg("Configuration loaded")
	//log.Debug().Caller().Msg("Starting warg")
	starter := NewStarter(config)
	if err := starter.DoStart(); err != nil {
		log.Fatal().Caller().Err(err).Msg("Failed to start")
		return err
	}
	return nil
}
