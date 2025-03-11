package boilerplate

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type HttpComponents struct {
	e         *echo.Echo
	addRoutes func(*echo.Echo, map[string]Dependent, ConfigIface)
	deps      map[string]Dependent
	c         ConfigIface
}

func NewHttpComponent(deps map[string]Dependent, c ConfigIface, addRouteFn func(*echo.Echo, map[string]Dependent, ConfigIface)) *HttpComponents {
	return &HttpComponents{
		addRoutes: addRouteFn,
		deps:      deps,
		c:         c,
	}
}

func (h *HttpComponents) Init() {
	h.e = echo.New()
	h.applyMiddleware()
	h.addRoutes(h.e, h.deps, h.c)
}

func (h *HttpComponents) applyMiddleware() {
	// hide the banner
	h.e.HideBanner = true
	// apply logger middleware
	logger := zerolog.New(os.Stdout)
	h.e.Use(middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogURI:       true,
			LogMethod:    true,
			LogRequestID: true,
			LogError:     true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				logger.Info().Str("Uri", v.URI).Str("Method", v.Method).Str("RequestID", v.RequestID).Msg("Request")
				return nil
			},
		},
	))
	// apply recover middleware
	h.e.Use(middleware.Recover())
	// apply request id middleware
	h.e.Use(middleware.RequestID())
}

func (h *HttpComponents) Run() {
	err := h.e.Start(fmt.Sprintf(":%d", h.c.GetServerPort()))
	if err != http.ErrServerClosed {
		log.Fatal().Err(err).Caller().Msg("could not able to start the server")
	}
}

func (h *HttpComponents) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5)
	defer cancel()
	err := h.e.Shutdown(ctx)
	if err != nil {
		log.Error().Err(err).Caller().Msg("not able to shutdown server within grace period")
	}
}
