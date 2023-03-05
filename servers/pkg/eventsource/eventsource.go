package eventsource

import (
	"context"
	"fmt"

	"github.com/nohns/semesterprojekt2/pkg/event"
)

var (
	ErrCommandNotHandled = fmt.Errorf("command not handled")
)

type EventStore interface {
	// Puts an event into the event store
	Put(ctx context.Context, event ...*event.Event) error
	// Gets all events from a specific aggregate
	Get(ctx context.Context, aggregateId event.CompositeID) (Cursor, error)
	// Gets a range of events from a specific aggregate. The range is inclusive.
	// If fromVersion is 0, range starts at the first event.
	// If toVersion is 0, range ends at the last event.
	Range(ctx context.Context, aggregateId event.CompositeID, fromVersion, toVersion int) (Cursor, error)
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

type AggregateRoot interface {
	Apply(evts ...*event.Event) error
	Process(cmd any) ([]*event.Event, error)
}

func Apply(cursor Cursor, ar AggregateRoot) error {
	for cursor.Next() {
		evt, err := cursor.Event()
		if err != nil {
			return err
		}

		if err := ar.Apply(evt); err != nil {
			return err
		}
	}

	return nil
}
