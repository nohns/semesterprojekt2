package bridge

import (
	"context"
	"fmt"

	"github.com/nohns/proto/bridge/events/v1"
	"github.com/nohns/servers/pkg/event"
	"github.com/nohns/servers/pkg/eventsource"
)

// Command for engaging a lock
type engageLock struct{}

// Command for disengaging a lock
type disengageLock struct{}

const (
	aggregateNameLock = "lock"
)

type Lock struct {
	ID      string
	Engaged bool
}

func (l *Lock) Apply(evts ...*event.Event) error {
	for _, e := range evts {
		switch e.Type {
		case "LockEngaged":
			l.Engaged = true
		case "LockDisengaged":
			l.Engaged = false
		}
	}

	return nil
}

func (l *Lock) Process(cmd any) ([]*event.Event, error) {
	evts := []*event.Event{}
	switch cmd.(type) {
	case *disengageLock:
		if l.Engaged {
			break
		}

		evt, err := event.FromMessage(event.ID(aggregateNameLock, l.ID), &events.LockDisengaged{
			LockId: l.ID,
		})
		if err != nil {
			return nil, err
		}
		evts = append(evts, evt)

	case *engageLock:
		if !l.Engaged {
			break
		}

		evt, err := event.FromMessage(event.ID(aggregateNameLock, l.ID), &events.LockEngaged{
			LockId: l.ID,
		})
		if err != nil {
			return nil, err
		}
		evts = append(evts, evt)
	default:
		return nil, eventsource.ErrCommandNotHandled
	}

	return evts, nil
}

// LockService is a service that handles lock commands, typically sent by the stream commands
type LockService struct {
	store eventsource.EventStore
	d     event.Dispatcher
}

func NewLockService(store eventsource.EventStore, d event.Dispatcher) *LockService {
	return &LockService{
		store: store,
		d:     d,
	}
}

// Disengages a lock by its id
func (s *LockService) Disengage(ctx context.Context, lockId string) error {
	err := s.processCmdForLock(ctx, lockId, &disengageLock{})
	if err != nil {
		return err
	}

	return nil
}

// Engages a lock by its id
func (s *LockService) Engage(ctx context.Context, lockId string) error {
	err := s.processCmdForLock(ctx, lockId, &engageLock{})
	if err != nil {
		return err
	}

	return nil
}

func (s *LockService) processCmdForLock(ctx context.Context, lockId string, cmd any) error {
	evts, err := s.store.Get(ctx, event.ID(aggregateNameLock, lockId))
	if err != nil {
		return fmt.Errorf("error getting lock events from store: %v", err)
	}

	// Create lock aggregate from events
	var l Lock
	if err := eventsource.Apply(evts, &l); err != nil {
		return fmt.Errorf("error applying lock events: %v", err)
	}

	// Handle command
	newEvts, err := l.Process(cmd)
	if err != nil {
		return fmt.Errorf("error handling command: %w", err)
	}

	// Put new events into store
	if err := s.store.Put(ctx, newEvts...); err != nil {
		return fmt.Errorf("error putting event into store: %v", err)
	}

	// Dispatch events
	if err := s.d.Dispatch(newEvts...); err != nil {
		return fmt.Errorf("error dispatching events: %v", err)
	}

	return nil
}
