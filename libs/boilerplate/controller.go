package boilerplate

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type Controller struct {
	Configuration any
	Dependents    map[string]Dependent
	Components    map[string]ComponentIface
}

func NewController(c ConfigIface, deps map[string]Dependent, routeDefnFn func(*echo.Echo, map[string]Dependent, ConfigIface)) *Controller {
	_components := make(map[string]ComponentIface)
	// we will add http component default
	_components["http"] = NewHttpComponent(deps, c, routeDefnFn)
	return &Controller{
		Configuration: c,
		Dependents:    deps,
		Components:    _components,
	}
}

func (c *Controller) Init() {
	log.Debug().Caller().Msg("initializing controller")
	// in init method we check if dependencies are live or not
	for name, dependency := range c.Dependents {
		log.Debug().Caller().Msg("pinging " + name)
		dependency.Ping()
	}
	// create components first. Here because AddDependency will execute after New

	// and init components
	for name, component := range c.Components {
		log.Debug().Caller().Msg("initializing " + name + " component")
		component.Init()
	}

}

func (c *Controller) Run() {
	for name, component := range c.Components {
		log.Debug().Caller().Msg("running " + name + " component")
		go component.Run()
	}
}

func (c *Controller) Close() {
	for name, dependency := range c.Dependents {
		log.Debug().Caller().Msg("closing dependency " + name)
		err := dependency.Close()
		if err != nil {
			log.Error().Err(err).Caller().Msg("did not able to ping " + name)
		}
	}
	for name, component := range c.Components {
		log.Debug().Caller().Msg("stopping " + name + " component")
		component.Stop()
	}
}
