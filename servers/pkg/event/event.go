package event

import (
	"strings"
	"time"

	protoV2 "google.golang.org/protobuf/proto"
)

type Event struct {
	// Type of event stored in Data field
	Type string
	// Globally scoped version of the event. Assigned in the event store in ascending order.
	Version int
	// Aggregate scoped version of the event. Assigned in the event store in ascending order.
	AggregateVersion int
	// Id of the aggregate the event belongs to. Can be empty if the event is not related to an aggregate.
	// Aggregate id is always on in the format of <aggregate type>:<aggregate id>
	AggregateId CompositeID
	// Occurrence time of the event
	At time.Time
	// Raw data of the event
	Data []byte
}

// FromMessage creates an event from an aggregate id and a proto message. Aggregate id can be empty, if the event is not related to an aggregate.
func FromMessage(aggregateId CompositeID, m protoV2.Message) (*Event, error) {
	data, err := protoV2.Marshal(m)
	if err != nil {
		return nil, err
	}
	return &Event{
		AggregateId: aggregateId,
		Type:        string(m.ProtoReflect().Descriptor().Name()),
		Data:        data,
	}, nil
}

// Unmarshall data from event into a proto message, given as a pointer
func (e *Event) Unmarshal(m protoV2.Message) error {
	return protoV2.Unmarshal(e.Data, m)
}

type CompositeID string

func ID(aggregateType, aggregateId string) CompositeID {
	return CompositeID(aggregateType + ":" + aggregateId)
}

func (id CompositeID) Type() string {
	return strings.Split(string(id), ":")[0]
}

func (id CompositeID) ID() string {
	return strings.Split(string(id), ":")[1]
}
