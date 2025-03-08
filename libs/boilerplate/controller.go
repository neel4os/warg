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
	// in init method we check if dependencies are live or not
	for name, dependency := range c.Dependents {
		err := dependency.Ping()
		if err != nil {
			log.Fatal().Err(err).Caller().Msg("did not able to ping " + name)
		}
	}
	// create components first. Here because AddDependency will execute after New

	// and init components
	for _, component := range c.Components {
		component.Init()
	}

}

func (c *Controller) Run() {
	for _, component := range c.Components {
		component.Run()
	}
}
