package controller

import (
	account_service "github.com/neel4os/warg/internal/account-management/domain/account/service"
	organization_service "github.com/neel4os/warg/internal/account-management/domain/organization/service"
	user_service "github.com/neel4os/warg/internal/account-management/domain/user/service"
	"github.com/neel4os/warg/internal/common/cache"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/rs/zerolog/log"
)

type controller struct {
	components []componentable
	cfg        *config.Config
}

func NewController(cfg *config.Config, dbcon *database.DataConn) *controller {
	_components := make([]componentable, 0)
	_components = append(_components, app.GetEventPlatform())
	_components = append(_components, NewHTTPComponent(cfg, dbcon))
	_components = append(_components, cache.NewIMCache(cfg))
	return &controller{components: _components, cfg: cfg}
}

func (c *controller) Init() {
	for _, comp := range c.components {
		log.Debug().Str("component", comp.Name()).Caller().Msg("Initializing " + comp.Name())
		comp.Init()
	}
	// Here we add all the command handlers to the event platform
	account_service.RegisterCommandHandlers(app.GetEventPlatform())
	organization_service.RegisterEventHandlers(app.GetEventPlatform())
	organization_service.RegisterCommandHandlers(app.GetEventPlatform())
	user_service.RegisterEventHandlers(app.GetEventPlatform())
	user_service.RegisterCommandHandlers(app.GetEventPlatform())
}

func (c *controller) Run() {
	for _, comp := range c.components {
		log.Debug().Str("component", comp.Name()).Caller().Msg("Running " + comp.Name())
		go comp.Run()
	}
}

func (c *controller) Stop() {
	for _, comp := range c.components {
		log.Debug().Str("component", comp.Name()).Caller().Msg("Stopping " + comp.Name())
		comp.Stop()
	}
}
