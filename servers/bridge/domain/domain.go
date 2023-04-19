package domain

import (
	"context"
)

type domain struct {
	uart uart
}

type uart interface {
	AwaitResponse(context.Context, string) (byte, error)
	Write([]byte) error
	Read() ([]byte, error)
}

func New(uart uart) *domain {
	return &domain{uart: uart}
}
