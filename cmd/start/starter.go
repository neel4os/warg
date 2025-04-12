package start

import (
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/server"
)

type starter struct {
	cfg *config.Config
}

func NewStarter(cfg *config.Config) *starter {
	return &starter{cfg: cfg}
}

func (s *starter) DoStart() error {
	server.StartServer(s.cfg)
	return nil
}
