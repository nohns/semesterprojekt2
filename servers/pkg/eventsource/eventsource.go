package eventsource

import (
	"context"

	"github.com/nohns/semesterprojekt2/pkg/event"
)

type EventStore interface {
	// Puts an event into the event store
	Put(ctx context.Context, event *event.Event) error
	// Gets all events from a specific aggregate
	Get(ctx context.Context, aggregateId string) (Cursor, error)
	// Gets a range of events from a specific aggregate. The range is inclusive.
	// If fromVersion is 0, range starts at the first event.
	// If toVersion is 0, range ends at the last event.
	Range(ctx context.Context, aggregateId string, fromVersion, toVersion int) (Cursor, error)
	// Plays all events regardsless of aggregate
	Play(ctx context.Context) (Cursor, error)
}

type Cursor interface {
	// Checks if there is another event to read
	Next() bool
	// Reads the next event
	Event() (*event.Event, error)
	// Closes the cursor. Must be called when done reading events.
	Close() error
}
