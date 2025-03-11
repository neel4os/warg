package boilerplate

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type Service struct {
	Name       string
	Controller *Controller
}

func NewService(serviceName string,
	config ConfigIface,
	dependencies map[string]Dependent,
	routes func(*echo.Echo, map[string]Dependent, ConfigIface)) *Service {
	ctrlr := NewController(config, dependencies, routes)
	return &Service{
		Name:       serviceName,
		Controller: ctrlr,
	}
}

func (s *Service) Initialize() {
	log.Debug().Caller().Msg("initializing service")
	s.Controller.Init()
}

func (s *Service) Run() {
	s.Controller.Run()
}

func (s *Service) Close() {
	s.Controller.Close()
}
