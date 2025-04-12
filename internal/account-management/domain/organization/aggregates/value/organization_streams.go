package value

import "github.com/google/uuid"

var (
	organizationStreamID   = "faa40d29-3066-4eb3-b886-8cab2de58919"
	organizationStreamName = "organization"
)

type organizationStream struct {
	streamID   uuid.UUID
	streamName string
}

func GetOrganizationStream() *organizationStream {
	instance := &organizationStream{
		streamID:   uuid.MustParse(organizationStreamID),
		streamName: organizationStreamName,
	}
	return instance
}

func (o *organizationStream) StreamID() uuid.UUID {
	return o.streamID
}

func (o *organizationStream) StreamName() string {
	return o.streamName
}
