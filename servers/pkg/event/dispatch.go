package event

import "context"

type Dispatcher interface {
	Dispatch(ctx context.Context, evts ...*Event) error
}
