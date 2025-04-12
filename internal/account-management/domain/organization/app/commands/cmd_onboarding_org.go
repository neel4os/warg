package commands

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	"gorm.io/datatypes"

	organization "github.com/neel4os/warg/internal/account-management/domain/organization/aggregates"
	"github.com/neel4os/warg/internal/account-management/domain/organization/aggregates/value"
	persistence_organization "github.com/neel4os/warg/internal/account-management/persistence/organization"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/common/errors"
	"github.com/neel4os/warg/internal/eventstore/domain/aggregates"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/neel4os/warg/internal/eventstore/domain/repositories"
	"github.com/neel4os/warg/internal/eventstore/persistence"
	"github.com/rs/zerolog/log"
)

type CreateOrgCommand struct {
	AccountId      uuid.UUID `json:"account_id" valid:"uuid,required~account_id required and must be a valid uuid"`
	OrgId          uuid.UUID
	OrgName        string `json:"org_name" valid:"alphanum,required~org_name required and must be alphanumeric"`
	DomainName     string
	OwnerFirstName string `json:"owner_first_name" valid:"alphanum,required~owner_first_name required and must be alphanumeric"`
	OwnerLastName  string `json:"owner_last_name" valid:"alphanum,required~owner_last_name required and must be alphanumeric"`
	OwnerEmail     string `json:"owner_email" valid:"email,required~owner_email required and must be a valid email"`
}

type CreateOrgCommandHandler struct {
	eventRepo repositories.EventRepositories
	eventBus  *cqrs.EventBus
	idpRepo   *persistence_organization.OrganizationKeycloakRepository
}

func NewCreateOrgCommandHandler() *CreateOrgCommandHandler {
	_config := config.GetConfig()
	dbcon := database.GetDataConn(*_config)
	eventRepo := persistence.NewEventDatabaseRepository(dbcon)
	eventPlatform := app.GetEventPlatform()
	idpRepo := persistence_organization.NewOrganizationKeycloakRepository()
	return &CreateOrgCommandHandler{
		eventRepo: eventRepo,
		eventBus:  eventPlatform.EventBus,
		idpRepo:   idpRepo,
	}
}

func (c *CreateOrgCommandHandler) Handle(ctx context.Context, cmd *CreateOrgCommand) error {
	log.Debug().Caller().Interface("Handling command org", &cmd).Msg("")

	// We dont know the org id yet because keycloak will generate it
	org_id, err := c.idpRepo.CreateOrganization(cmd.DomainName)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while creating organization")
		return err
	}
	cmd.OrgId, err = uuid.Parse(org_id)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while parsing organization id")
		return err
	}
	orgStream := value.GetOrganizationStream()
	_event := aggregates.NewEvent(
		orgStream.StreamID(),
		orgStream.StreamName()+"."+cmd.OrgId.String(),
	)
	_org_data := organization.Organization{
		ID:         cmd.OrgId,
		Name:       cmd.OrgName,
		AccountId:  cmd.AccountId,
		DomainName: cmd.DomainName,
		Status:     value.OrganizationStatusPending,
	}
	_req_bytes, err := json.Marshal(_org_data)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while marshalling command")
		return errors.NewJSONMarhsalError(err.Error())
	}
	_event = _event.SetInitiatorType(string(aggregates.InitiatorTypeSystem)).
		SetMetadata(datatypes.JSON{}).
		SetEventType("organization_created").
		SetEventData(datatypes.JSON(_req_bytes)).
		SetInitiatorName("CreateOrgCommandHandler")
	// create the event
	tx, err := c.eventRepo.CreateEvent(_event)
	if err != nil {
		tx.Rollback()
		// TODO: Delete the org from keycloak
		log.Error().Caller().Err(err).Msg("failed to create event")
		return errors.NewDatabaseOperationError("failed to create event")
	}
	err = c.eventBus.Publish(ctx, &organization.OrganizationCreated{
		AccountId:      cmd.AccountId,
		OrganizationId: cmd.OrgId,
		OwnerEmail:     cmd.OwnerEmail,
		OwnerFirstName: cmd.OwnerFirstName,
		OwnerLastName:  cmd.OwnerLastName,
	})
	if err != nil {
		log.Error().Caller().Err(err).Msg("failed to publish event")
		tx.Rollback()
		return errors.NewInternalServerError("failed to publish event")
	}
	// refresh the materialized view
	// tx = tx.Exec("REFRESH MATERIALIZED VIEW CONCURRENTLY organization")
	// if tx.Error != nil {
	// 	tx.Rollback()
	// 	log.Error().Caller().Err(tx.Error).Msg("failed to refresh materialized view")
	// 	return errors.NewDatabaseOperationError("failed to refresh materialized view")
	// }
	tx.Commit()
	return nil
}
