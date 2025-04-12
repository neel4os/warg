package aggregates

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type InitiatorType string

const (
	InitiatorTypeSystem  InitiatorType = "system"
	InitiatorTypeUser    InitiatorType = "user"
	InitiatorTypeService InitiatorType = "service"
)

type Event struct {
	ID            uuid.UUID      `gorm:"primaryKey"`
	StreamID      uuid.UUID      `gorm:"not null"`
	StreamName    string         `gorm:"not null"`
	EventType     string         `gorm:"not null"`
	EventData     datatypes.JSON `gorm:"type:jsonb;not null"`
	Metadata      datatypes.JSON `gorm:"type:jsonb"`
	Version       int            `gorm:"not null"`
	InitiatorType InitiatorType  `gorm:"not null"`
	InitiatorName string         `gorm:"not null"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
}

func NewEvent(streamID uuid.UUID, streamName string) *Event {
	return &Event{
		ID:            uuid.New(),
		StreamID:      streamID,
		StreamName:    streamName,
		EventType:     "",
		EventData:     nil,
		Metadata:      nil,
		Version:       0,
		InitiatorType: "",
		InitiatorName: "",
	}
}

func (e *Event) SetEventType(eventType string) *Event {
	e.EventType = eventType
	return e
}

func (e *Event) SetEventData(eventData datatypes.JSON) *Event {
	e.EventData = eventData
	return e
}

func (e *Event) SetMetadata(metadata datatypes.JSON) *Event {
	e.Metadata = metadata
	return e
}

func (e *Event) SetInitiatorType(initiatorType string) *Event {
	e.InitiatorType = InitiatorType(initiatorType)
	return e
}

func (e *Event) SetInitiatorName(initiatorName string) *Event {
	e.InitiatorName = initiatorName
	return e
}
