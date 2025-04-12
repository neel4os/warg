package service

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/neel4os/warg/internal/account-management/domain/organization/app/commands"
	"github.com/neel4os/warg/internal/account-management/domain/organization/app/events"
	eventstore "github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/rs/zerolog/log"
)

func RegisterEventHandlers(ep *eventstore.EventPlatform) {
	log.Info().Str("component", "service").Msg("Registering event handlers")
	err := ep.AddEventProcessorHandler(cqrs.NewEventHandler("CreateOrgOnAccountCreated", events.NewCreateOrgOnAccountCreatedEventHandler().Handle))
	if err != nil {
		panic(err)
	}
}

func RegisterCommandHandlers(ep *eventstore.EventPlatform) {
	log.Info().Str("component", "service").Msg("Registering command handlers")
	err := ep.AddCommandProcessorHandler(cqrs.NewCommandHandler("OrganizationCreated", commands.NewCreateOrgCommandHandler().Handle))
	if err != nil {
		panic(err)
	}
}
