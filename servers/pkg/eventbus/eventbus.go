package eventbus

import (
	"context"
	"sort"
	"sync"

	"github.com/nohns/servers/pkg/event"
)

var consumeContext = context.Background()

type subscriber struct {
	priority int
	c        event.Consumer
}

type bus struct {
	// subscriptions contains all subscribers with their priority
	subscriptions []subscriber
	mu            sync.RWMutex
}

// Creates new event bus, which implements the Bus interface, exposing Subscribe() and Publish() methods
func New() *bus {
	return &bus{
		subscriptions: make([]subscriber, 0),
	}
}

// Subscribe adds a new consumer to the bus
func (b *bus) Listen(priority int, c event.Consumer) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Make sure consumer slice for topic is initialized
	b.subscriptions = append(b.subscriptions, subscriber{
		priority: priority,
		c:        c,
	})

	// Sort subscribers by priority
	sort.Slice(b.subscriptions, func(i, j int) bool {
		return b.subscriptions[i].priority < b.subscriptions[j].priority
	})
}

// Publish publishes an event to all subscribers
func (b *bus) Dispatch(evts ...*event.Event) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, e := range evts {
		for _, s := range b.subscriptions {
			if err := s.c.Consume(consumeContext, e); err != nil {
				return err
			}
		}
	}

	return nil
}

type ConsumeFunc func(ctx context.Context, evt *event.Event) error

func (f ConsumeFunc) Consume(ctx context.Context, evt *event.Event) error {
	return f(ctx, evt)
}
