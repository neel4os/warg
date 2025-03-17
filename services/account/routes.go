package account

import (
	"github.com/labstack/echo/v4"
	"github.com/neel4os/warg/libs/boilerplate"
	"github.com/rs/zerolog/log"
)

func Routes(e *echo.Echo, deps map[string]boilerplate.Dependent, cfg boilerplate.ConfigIface) {
	_, ok := cfg.(*Config)
	if !ok {
		log.Fatal().Caller().Msg("could not able to convert into storage config")
	}
	h := NewAccountHandler(deps, cfg)
	v0 := e.Group("api/v0")
	v0.GET("/health", h.GetHealth)
	v0.POST("/user", h.CreateUser)
}
