package event

type Dispatcher interface {
	Dispatch(evts ...*Event) error
}
