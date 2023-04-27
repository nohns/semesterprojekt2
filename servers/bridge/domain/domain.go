package domain

import "context"

type domain struct {
	uart uart
}

type uart interface {
	AwaitResponse(ctx context.Context, cmd int) ([]byte, error)
}

func New(uart uart) *domain {
	return &domain{uart: uart}
}
