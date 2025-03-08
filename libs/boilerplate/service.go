package boilerplate

import (
	"github.com/labstack/echo/v4"
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
	s.Controller.Init()
}

func (s *Service) Run() {
	s.Controller.Run()
}
