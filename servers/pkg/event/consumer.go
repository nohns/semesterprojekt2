package event

import "context"

type Consumer interface {
	Consume(ctx context.Context, evt *Event) error
}
