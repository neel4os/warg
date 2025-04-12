package value

import "github.com/google/uuid"

var userStreamID = "264aa0a0-d4b1-4cdd-ab31-dd1a1239acc3"

type userStream struct {
	streamID   uuid.UUID
	streamName string
}

func GetUserStream() *userStream {
	instance := &userStream{
		streamID:   uuid.MustParse(userStreamID),
		streamName: "user",
	}
	return instance
}

func (u *userStream) StreamID() uuid.UUID {
	return u.streamID
}

func (u *userStream) StreamName() string {
	return u.streamName
}
