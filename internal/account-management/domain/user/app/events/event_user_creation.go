package events

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	organization "github.com/neel4os/warg/internal/account-management/domain/organization/aggregates"
	"github.com/neel4os/warg/internal/account-management/domain/user/app/commands"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/rs/zerolog/log"
)

type CreateUserOnOrgCreatedEventHandler struct {
	commandBus *cqrs.CommandBus
}

func NewCreateUserOnOrgCreatedEventHandler() *CreateUserOnOrgCreatedEventHandler {
	ep := app.GetEventPlatform()
	return &CreateUserOnOrgCreatedEventHandler{
		commandBus: ep.CommandBus,
	}
}

func (c *CreateUserOnOrgCreatedEventHandler) Handle(ctx context.Context, event *organization.OrganizationCreated) error {
	log.Info().Caller().Interface("Handling event CreateUserOnOrgCreated ", &event).Msg("")
	createUserCommand := &commands.CreateUserCommand{
		AccountId:      event.AccountId,
		OrgId:          event.OrganizationId,
		OwnerFirstName: event.OwnerFirstName,
		OwnerLastName:  event.OwnerLastName,
		OwnerEmail:     event.OwnerEmail,
	}
	return c.commandBus.Send(ctx, createUserCommand)
}
