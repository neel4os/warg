package app

import (
	"context"
	"sync"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	//"github.com/neel4os/warg/internal/eventstore/logs"
	"github.com/rs/zerolog/log"
	//"github.com/neel4os/warg/internal/account-management/domain/account/app/commands"
)

type EventPlatform struct {
	CommandBus       *cqrs.CommandBus
	CommandProcessor *cqrs.CommandProcessor
	EventBus         *cqrs.EventBus
	EventProcessor   *cqrs.EventProcessor
	Router           *message.Router
	Subscriber       message.Subscriber
	Publisher        message.Publisher
}

func newEventPlatform() *EventPlatform {
	//logger := logs.NewZerologLoggerAdapter(log.Logger.With().Str("component", "event-platform").Logger())
	logger := watermill.NewStdLogger(true, true)
	cqrsMarshaler := cqrs.JSONMarshaler{
		GenerateName: cqrs.StructName,
	}
	pubsub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		logger.Error("failed to create router", err, nil)
		return nil
	}
	router.AddMiddleware(middleware.Recoverer)
	commandBus, err := cqrs.NewCommandBusWithConfig(pubsub, cqrs.CommandBusConfig{
		GeneratePublishTopic: func(cbgptp cqrs.CommandBusGeneratePublishTopicParams) (string, error) {
			return "commands." + cbgptp.CommandName, nil
		},
		OnSend: func(params cqrs.CommandBusOnSendParams) error {
			logger.Info("Sending command in router code", watermill.LogFields{
				"command": params.CommandName,
			})
			params.Message.Metadata.Set("sent_at", time.Now().Format(time.RFC3339))
			log.Info().Interface("params", params).Msg("Sending command in router code")
			return nil
		},
		Marshaler: cqrsMarshaler,
		Logger:    logger,
	})
	if err != nil {
		logger.Error("failed to create command bus", err, nil)
		return nil
	}
	commandProcessor, err := cqrs.NewCommandProcessorWithConfig(router, cqrs.CommandProcessorConfig{
		GenerateSubscribeTopic: func(cpgstp cqrs.CommandProcessorGenerateSubscribeTopicParams) (string, error) {
			return "commands." + cpgstp.CommandName, nil
		},
		SubscriberConstructor: func(cpscp cqrs.CommandProcessorSubscriberConstructorParams) (message.Subscriber, error) {
			return pubsub, nil
		},
		Marshaler: cqrsMarshaler,
		Logger:    logger,
		OnHandle: func(params cqrs.CommandProcessorOnHandleParams) error {
			start := time.Now()
			err := params.Handler.Handle(params.Message.Context(), params.Command)
			logger.Info("Handled command", watermill.LogFields{
				"command_name": params.CommandName,
				"took":         time.Since(start),
				"error":        err,
			})
			return err
		},
	})
	if err != nil {
		logger.Error("failed to create command processor", err, nil)
		return nil
	}
	eventBus, err := cqrs.NewEventBusWithConfig(pubsub, cqrs.EventBusConfig{
		GeneratePublishTopic: func(ebgptp cqrs.GenerateEventPublishTopicParams) (string, error) {
			return "events." + ebgptp.EventName, nil
		},
		OnPublish: func(params cqrs.OnEventSendParams) error {
			logger.Info("Publishing event", watermill.LogFields{
				"event": params.EventName,
			})
			params.Message.Metadata.Set("published_at", time.Now().Format(time.RFC3339))
			return nil
		},
		Marshaler: cqrsMarshaler,
		Logger:    logger,
	})
	if err != nil {
		logger.Error("failed to create event bus", err, nil)
		return nil
	}
	eventProcessor, err := cqrs.NewEventProcessorWithConfig(router, cqrs.EventProcessorConfig{
		GenerateSubscribeTopic: func(epgstp cqrs.EventProcessorGenerateSubscribeTopicParams) (string, error) {
			return "events." + epgstp.EventName, nil
		},
		SubscriberConstructor: func(epscp cqrs.EventProcessorSubscriberConstructorParams) (message.Subscriber, error) {
			return pubsub, nil
		},
		Marshaler: cqrsMarshaler,
		Logger:    logger,
	})
	if err != nil {
		logger.Error("failed to create event processor", err, nil)
		return nil
	}
	return &EventPlatform{
		CommandBus:       commandBus,
		CommandProcessor: commandProcessor,
		EventBus:         eventBus,
		EventProcessor:   eventProcessor,
		Router:           router,
		Subscriber:       pubsub,
		Publisher:        pubsub,
	}
}

var (
	instance *EventPlatform
	once     sync.Once
)

func GetEventPlatform() *EventPlatform {
	once.Do(func() {
		instance = newEventPlatform()
	})
	return instance
}

func (ep *EventPlatform) Name() string {
	return "EventPlatform"
}

func (ep *EventPlatform) Init() {
}

func (ep *EventPlatform) Run() {
	if err := ep.Router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func (ep *EventPlatform) Stop() {
	err := ep.Router.Close()
	if err != nil {
		panic(err)
	}
}

func (ep *EventPlatform) AddCommandProcessorHandler(chandler cqrs.CommandHandler) error {
	return ep.CommandProcessor.AddHandlers(chandler)
}

func (ep *EventPlatform) AddEventProcessorHandler(ehandler cqrs.EventHandler) error {
	return ep.EventProcessor.AddHandlers(ehandler)
}
