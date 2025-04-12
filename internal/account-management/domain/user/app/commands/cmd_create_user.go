package commands

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	user "github.com/neel4os/warg/internal/account-management/domain/user/aggregates"
	"github.com/neel4os/warg/internal/account-management/domain/user/aggregates/value"
	persistence_user "github.com/neel4os/warg/internal/account-management/persistence/users"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/common/errors"
	"github.com/neel4os/warg/internal/eventstore/domain/aggregates"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	"github.com/neel4os/warg/internal/eventstore/domain/repositories"
	"github.com/neel4os/warg/internal/eventstore/persistence"
	"github.com/rs/zerolog/log"
	"gorm.io/datatypes"
)

type CreateUserCommand struct {
	AccountId      uuid.UUID `json:"account_id" valid:"uuid,required~account_id required and must be a valid uuid"`
	OrgId          uuid.UUID `json:"org_id" valid:"uuid,required~org_id required and must be a valid uuid"`
	OwnerFirstName string    `json:"owner_first_name" valid:"alphanum,required~owner_first_name required and must be alphanumeric"`
	OwnerLastName  string    `json:"owner_last_name" valid:"alphanum,required~owner_last_name required and must be alphanumeric"`
	OwnerEmail     string    `json:"owner_email" valid:"email,required~owner_email required and must be a valid email"`
	UserId         uuid.UUID
}

type CreateUserCommandHandler struct {
	eventRepo repositories.EventRepositories
	eventBus  *cqrs.EventBus
	idpRepo   *persistence_user.UserKeycloakRepository
}

func NewCreateUserCommandHandler() *CreateUserCommandHandler {
	_config := config.GetConfig()
	dbcon := database.GetDataConn(*_config)
	eventRepo := persistence.NewEventDatabaseRepository(dbcon)
	eventPlatform := app.GetEventPlatform()
	idpRepo := persistence_user.NewUserKeycloakRepository()
	return &CreateUserCommandHandler{
		eventRepo: eventRepo,
		eventBus:  eventPlatform.EventBus,
		idpRepo:   idpRepo,
	}
}

func (c *CreateUserCommandHandler) Handle(ctx context.Context, cmd *CreateUserCommand) error {
	log.Debug().Caller().Interface("Handling command create user", &cmd).Msg("")
	user_id, err := c.idpRepo.CreateUser(cmd.OwnerEmail, cmd.OwnerFirstName, cmd.OwnerLastName)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while creating user")
		return err
	}
	cmd.UserId, err = uuid.Parse(user_id)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while parsing user id")
		return err
	}
	userStream := value.GetUserStream()
	_event := aggregates.NewEvent(
		userStream.StreamID(),
		userStream.StreamName()+"."+user_id,
	)
	_user_data := user.User{
		ID:        cmd.UserId,
		FirstName: cmd.OwnerFirstName,
		LastName:  cmd.OwnerLastName,
		Email:     cmd.OwnerEmail,
		Status:    value.UserStatusPending,
	}
	_req_bytes, err := json.Marshal(_user_data)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while marshalling user data")
		return errors.NewJSONMarhsalError(err.Error())
	}
	_event = _event.SetInitiatorType(string(aggregates.InitiatorTypeSystem)).
		SetInitiatorName("CreateUserCommandHandler").
		SetEventType("user_created").
		SetEventData(datatypes.JSON(_req_bytes)).
		SetMetadata(datatypes.JSON{})
	// create the event
	tx, err := c.eventRepo.CreateEvent(_event)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while creating event")
		tx.Rollback()
		return errors.NewDatabaseOperationError(err.Error())
	}
	// publish the event
	tx.Commit()
	return nil
}
