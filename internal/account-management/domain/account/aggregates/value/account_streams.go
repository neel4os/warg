package value

import "github.com/google/uuid"

var accountStreamID = "328ab4ec-b5d6-4650-ada3-18f26e0f5752"

type accountStream struct {
	streamID   uuid.UUID
	streamName string
}

func GetAccountStream() *accountStream {
	instance := &accountStream{
		streamID:   uuid.MustParse(accountStreamID),
		streamName: "account",
	}
	return instance
}

func (a *accountStream) StreamID() uuid.UUID {
	return a.streamID
}

func (a *accountStream) StreamName() string {
	return a.streamName
}
