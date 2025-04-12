package commands

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/neel4os/warg/internal/account-management/domain/account/aggregates/value"
	"github.com/neel4os/warg/internal/account-management/domain/account/app/events"

	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/database"
	"github.com/neel4os/warg/internal/common/errors"
	"github.com/neel4os/warg/internal/eventstore/domain/aggregates"
	"github.com/neel4os/warg/internal/eventstore/domain/app"
	domain_repository "github.com/neel4os/warg/internal/eventstore/domain/repositories"
	event_persistence "github.com/neel4os/warg/internal/eventstore/persistence"
	"gorm.io/datatypes"
)

// type AccountOnboardHandler decorators.CommandHandler[value.AccountCreationRequest]

type OnBoardAccount struct {
	AccountName string `json:"account_name" valid:"alphanum,required~account_name required and must be alphanumeric"`
	FirstName   string `json:"first_name" valid:"alpha,required~first_name required and must be alphabetic"`
	LastName    string `json:"last_name" valid:"alpha,required~last_name required and must be alphabetic"`
	Email       string `json:"email" valid:"email,required~email required and must be a valid email address"`
	AccountId   uuid.UUID
}

type AccountOnboardingCommandHandler struct {
	// accountRepo account_repository.AccountRepositoryInterface
	eventRepo domain_repository.EventRepositories
	eventBus  *cqrs.EventBus
}

func NewAccountOnboardCommandHandler() *AccountOnboardingCommandHandler {
	_config := config.GetConfig()
	dbcon := database.GetDataConn(*_config)
	// accountRepo := account_persistence.NewAccountDatabaseRepository()
	eventRepo := event_persistence.NewEventDatabaseRepository(dbcon)
	eventPlatform := app.GetEventPlatform()
	return &AccountOnboardingCommandHandler{
		// accountRepo: accountRepo,
		eventRepo: eventRepo,
		eventBus:  eventPlatform.EventBus,
	}
}

func (h *AccountOnboardingCommandHandler) Handle(ctx context.Context, cmd *OnBoardAccount) error {
	log.Debug().Caller().Interface("Handling command", &cmd).Msg("")
	// return h.eventBus.Publish(ctx, &events.AccountOnboarded{})
	// create account ID
	_account_ID := cmd.AccountId
	// get stream info of account stream
	accountStream := value.GetAccountStream()
	// create event
	_event := aggregates.NewEvent(
		accountStream.StreamID(),
		accountStream.StreamName()+"."+_account_ID.String(),
	)
	_req_bytes, err := json.Marshal(cmd)
	if err != nil {
		return errors.NewJSONMarhsalError(err.Error())
	}
	// make event more informative
	_event = _event.SetInitiatorType("user").
		SetInitiatorName(cmd.Email).
		SetMetadata(datatypes.JSON{}).
		SetEventData(datatypes.JSON(_req_bytes)).
		SetEventType("account_onboarded")
	// create the event
	tx, err := h.eventRepo.CreateEvent(_event)
	if err != nil {
		log.Error().Caller().Err(err).Msg("failed to create event")
		return errors.NewDatabaseOperationError("failed to create event")
	}
	// refresh the materialized view
	tx = tx.Exec("REFRESH MATERIALIZED VIEW CONCURRENTLY account")
	if tx.Error != nil {
		tx.Rollback()
		log.Error().Caller().Err(tx.Error).Msg("failed to refresh materialized view")
		return errors.NewDatabaseOperationError("failed to refresh materialized view")
	}
	// Now publish the event
	err = h.eventBus.Publish(ctx, &events.AccountOnboarded{
		AccountId:   _account_ID,
		AccountName: cmd.AccountName,
		FirstName:   cmd.FirstName,
		LastName:    cmd.LastName,
		Email:       cmd.Email,
		Status:      value.AccountStatusPending,
	})
	if err != nil {
		log.Error().Caller().Err(err).Msg("failed to publish event")
		tx.Rollback()
		return errors.NewInternalServerError("failed to publish event")
	}
	// if everything is ok, commit the transaction
	tx.Commit()
	return nil
}
