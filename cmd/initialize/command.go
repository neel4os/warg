package initialize

import (
	"github.com/neel4os/warg/internal/common/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize warg instance",
		Long:  `Initialize a new Warg instance`,
		RunE:  InitWarg,
	}
}

func InitWarg(cmd *cobra.Command, args []string) error {
	log.Info().Caller().Msg("Initializing Warg instance")
	log.Info().Caller().Msg("Reading Warg configuration")
	config := config.GetConfig()
	log.Debug().Interface("config", config).Caller().Msg("Configuration loaded")
	log.Debug().Caller().Msg("Initializing warg")
	initializer := NewInitilizer(config)
	if err := initializer.DoInitialize(); err != nil {
		log.Fatal().Caller().Err(err).Msg("Failed to initialize")
		return err
	}
	log.Info().Caller().Msg("Warg instance initialized")
	return nil
}
